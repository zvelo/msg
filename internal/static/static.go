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

	"/schema.graphql": {
		local:   "schema.graphql",
		size:    7156,
		modtime: 1531965751,
		compressed: `
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
bzGgX+ZFwG+7/wkAAP//8kFRh/QbAAA=
`,
	},

	"/": {
		isDir: true,
		local: "",
	},
}
