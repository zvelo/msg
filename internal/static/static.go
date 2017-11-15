package static

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

	"/apiv1.swagger.json": {
		local:   "static/apiv1.swagger.json",
		size:    25153,
		modtime: 1510726016,
		compressed: `
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
AA==
`,
	},

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    7049,
		modtime: 1510725744,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
