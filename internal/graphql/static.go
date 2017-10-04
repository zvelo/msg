package graphql

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    6748,
		modtime: 1507103933,
		compressed: `
H4sIAAAAAAAA/3RYT4/kto6/16fw3LLAu7xkkQX6Rku0zS5Z0lCSq6uDIKjtqc0Mtqd70l29wOwi333h
EqmqmbycTNkSSfHPj6RfHz4ePx+6/9t03R9vx5evN9379bH5c7M5ff1yrKvz57eXxx/eXh5vunR6+fT0
+7t/dA+Hx8f/PDz8t776R/fhcDq8Hk833S/2cDqk4yl//XJ89+u7fxO2fPzy+PXdpusenp9Ox6fTD5uu
a4ubrrAzlX5XP3wn4Pzy72Rsuu4vYl6Or2+Ppx9ejn+8HV9PZG86slfKrB/Xmx6f3j53V9zO9zWQcQxM
95Ap+E3XzeDIUChp03VoprCe/PT05e3UTcfDh+PL+dTT4fOxmWjTdf9zeHy7etGOXK6qxr265Mczv5vu
l8r43a+bKyNdsbp46HzjM6dv7rrpuuPLy/PLeupwentth+ryfODh+cPxpqNq88/H19fD78e/lbJa7F+K
eTm+fnl+ej2KFW/UnBpXVaBYvi4uzGVvVedwOv7+/PLpfw+nT89PN90vpr74erbC58Pjp4dPzysnOTTr
m/WuDx+fmxm/82vbtxxfPnx6qMKK3/qwW51rHMK3Tv5OuXb+WsuvN11Tb3V2Zf1X3UTmd+b8xgmfvzwe
T8ebrn9+fjwent79xXVd91/H08NHo+7adN3j84NY6bs7q1LXl/xNAnq/6TroA69R/du/b7oucjBTIIO6
cjQIPf7zpzMxwtw78qMs9ufn+gKZ6lZfLOX6PgZujAfw+poho+UzmfDu/ESrVA49GBPO9Ezu/FwoOPSi
1g4hBp/ONIwjY0q01E8DeQTO9ZM+Sw4ROKe2YoxAVXxf9gmdM1CXBtgUlwtXdiaG9jr4BYWjCSXKhrKc
n5YwYVUVF5zIuPp9gmymHsy2rvY9k1VZrtwV3utqJk8LVGPNIQc2e2UShoED2GpH5EGPRDLbEqtBA1hI
iVIWo1phlbmI8IV8hhH18A5G8flqkTlkNSFUz/YUMppJ6PsUhiwGI1dVGQJjyrzXLWOotvDQ7DcU9tRW
k9x0qBLmbRYCM1TbxWmf0CjDyKFKYszEOIW58kkTxagR2JdEHlN17m3owS4avgYYkaMTU5jgbkOvYQJk
9cigwXAb+oTAcmtfOKkUxlRUupmCEyNmdCbMc8mo0bqyuBJeyYmchb4kUYtpTltyTiLEbFXKBMLIchmT
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
