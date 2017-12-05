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
		size:    25851,
		modtime: 1512498290,
		compressed: `
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
`,
	},

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    7049,
		modtime: 1510787196,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
