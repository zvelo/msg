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
		size:    25633,
		modtime: 1515179146,
		compressed: `
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
ZAAA
`,
	},

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    6963,
		modtime: 1515180578,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
