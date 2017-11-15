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
		size:    25629,
		modtime: 1510778467,
		compressed: `
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
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
