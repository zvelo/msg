package mock

import (
	"bytes"
	"compress/gzip"
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
	"strings"
	"time"

	"github.com/NYTimes/gziphandler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"zvelo.io/msg"
)

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

type ServeOption func(*serveOpts)

type serveOpts struct {
	ready            chan<- struct{}
	compressionLevel int
}

func defaultServeOpts() *serveOpts {
	return &serveOpts{
		compressionLevel: gzip.DefaultCompression,
	}
}

func WithOnReady(val chan<- struct{}) ServeOption {
	return func(o *serveOpts) {
		o.ready = val
	}
}

func WithCompressionLevel(val int) ServeOption {
	if val == 0 {
		val = gzip.DefaultCompression
	}

	return func(o *serveOpts) {
		o.compressionLevel = val
	}
}

// ListenAndServeTLS listens for tls connections using a self-signed certificate
// on addr and serves the mock APIServer
func ListenAndServeTLS(ctx context.Context, addr string, opts ...ServeOption) error {
	if addr == "" {
		addr = ":https"
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return ServeTLS(ctx, l, opts...)
}

func ServeTLS(ctx context.Context, l net.Listener, opts ...ServeOption) error {
	o := defaultServeOpts()
	for _, opt := range opts {
		opt(o)
	}

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

	graphQLHandler, err := msg.GraphQLHandler(msg.NewAPIClient(conn))
	if err != nil {
		return err
	}

	rest := msg.NewServeMux()
	if err = msg.RegisterAPIHandler(ctx, rest, conn); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphQLHandler)
	mux.Handle("/", rest)

	gz, err := gziphandler.NewGzipLevelHandler(o.compressionLevel)
	if err != nil {
		return err
	}

	h := handler{
		grpc: grpc.NewServer(),
		rest: gz(mux),
	}

	msg.RegisterAPIServer(h.grpc, &apiServer{})

	s := http.Server{
		Handler: h,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			NextProtos:   []string{"h2"},
		},
	}

	errCh := make(chan error)
	go func() { errCh <- s.ServeTLS(l, "", "") }()

	if o.ready != nil {
		close(o.ready)
	}

	return <-errCh
}
