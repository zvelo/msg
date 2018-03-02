// Code generated by "esc "; DO NOT EDIT.

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
		_ = f.Close()
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
		local:   "apiv1.swagger.json",
		size:    19737,
		modtime: 1519984029,
		compressed: `
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
m+v/CwAA///fHrF8GU0AAA==
`,
	},

	"/schema.graphql": {
		local:   "schema.graphql",
		size:    7156,
		modtime: 1519984016,
		compressed: `
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
E/8WA/p5nBn8sftPAAAA//97CNhU9BsAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},
}
