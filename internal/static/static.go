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

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    6772,
		modtime: 1507622765,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
