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
		size:    23898,
		modtime: 1507655046,
		compressed: `
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
`,
	},

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    6772,
		modtime: 1507623094,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
