package mock

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"zvelo.io/msg/internal/static"
	msg "zvelo.io/msg/msgpb"
)

var staticFS = static.FS(false)

type handler struct {
	grpc *grpc.Server
	rest http.Handler
}

func grpcContentType(t string) bool {
	const e = "application/grpc"

	if !strings.HasPrefix(t, e) {
		return false
	}

	if len(t) > len(e) && t[len(e)] != '+' && t[len(e)] != ';' {
		return false
	}

	return true
}

func isGRPC(r *http.Request) bool {
	if r.ProtoMajor != 2 {
		return false
	}

	if r.Method != "POST" {
		return false
	}

	return grpcContentType(r.Header.Get("Content-Type"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isGRPC(r) {
		h.grpc.ServeHTTP(w, r)
		return
	}

	h.rest.ServeHTTP(w, r)
}

func selfSignedCert() (*x509.Certificate, tls.Certificate, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	notAfter := time.Now().Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	template := x509.Certificate{
		IsCA:                  true,
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{Organization: []string{"Acme Co"}},
		NotBefore:             time.Now(),
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"mock.api.zvelo.com"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	x509Cert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	var certPEM bytes.Buffer
	if err = pem.Encode(&certPEM, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, tls.Certificate{}, err
	}

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	var keyPEM bytes.Buffer
	if err = pem.Encode(&keyPEM, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes}); err != nil {
		return nil, tls.Certificate{}, err
	}

	tlsCert, err := tls.X509KeyPair(certPEM.Bytes(), keyPEM.Bytes())
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	return x509Cert, tlsCert, nil
}

type APIv1Option func(*apiv1)

type apiv1 struct {
	ready chan<- struct{}
}

func defaultsAPIv1() *apiv1 {
	return &apiv1{}
}

func WhenReady(val chan<- struct{}) APIv1Option {
	return func(o *apiv1) {
		o.ready = val
	}
}

type Server interface {
	ServeTLS(ctx context.Context, l net.Listener) error
	ListenAndServeTLS(ctx context.Context, addr string) error
}

func APIv1(opts ...APIv1Option) Server {
	o := defaultsAPIv1()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// ListenAndServeTLS listens for tls connections using a self-signed certificate
// on addr and serves the mock APIServer
func (srv apiv1) ListenAndServeTLS(ctx context.Context, addr string) error {
	if addr == "" {
		addr = ":https"
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return srv.ServeTLS(ctx, l)
}

func (srv apiv1) ServeTLS(ctx context.Context, l net.Listener) error {
	x509Cert, tlsCert, err := selfSignedCert()
	if err != nil {
		return err
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AddCert(x509Cert)

	conn, err := grpc.DialContext(ctx, l.Addr().String(),
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				ServerName: "mock.api.zvelo.com",
				RootCAs:    rootCAs,
			}),
		),
	)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := conn.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error closing connection: %s\n", cerr)
		}
	}()

	graphQLHandler, err := msg.GraphQLHandler(msg.NewAPIv1Client(conn))
	if err != nil {
		return err
	}

	rest := msg.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			if k, ok := runtime.DefaultHeaderMatcher(key); ok {
				return k, ok
			}

			if strings.HasPrefix(key, "Zvelo-Mock-") {
				return key, true
			}

			return "", false
		}),
	)
	if err = msg.RegisterAPIv1Handler(ctx, rest, conn); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphQLHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if shouldServeFile(r.URL.Path) {
			http.FileServer(staticFS).ServeHTTP(w, r)
			return
		}

		rest.ServeHTTP(w, r)
	})

	h := handler{
		grpc: grpc.NewServer(),
		rest: mux,
	}

	msg.RegisterAPIv1Server(h.grpc, &apiServer{})

	s := http.Server{
		Handler: h,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			NextProtos:   []string{"h2"},
		},
	}

	errCh := make(chan error)
	go func() { errCh <- s.ServeTLS(l, "", "") }()

	if srv.ready != nil {
		close(srv.ready)
	}

	return <-errCh
}

func shouldServeFile(name string) bool {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}

	if name == "/" {
		return false
	}

	f, err := staticFS.Open(path.Clean(name))
	if err != nil {
		return false
	}

	defer func() { _ = f.Close() }()

	if _, err := f.Stat(); err != nil {
		return false
	}

	return true
}
