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
		modtime: 1507613277,
		compressed: `
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
`,
	},

	"/swagger.json": {
		local:   "static/swagger.json",
		size:    23639,
		modtime: 1507622543,
		compressed: `
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
AAD//4AhAidXXAAA
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
