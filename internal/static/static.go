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
		size:    18860,
		modtime: 1516138906,
		compressed: `
H4sIAAAAAAAA/9Q825LbNpbv+QqUdh9mqjLu2NmarfK+LAiCJCwSYABQsrxJObQa3c2JJCok5UyPy/++
BVKtFs9hu33LbvJk9yFwcO4XXPTuG0Jm63rXHraunT0n//MNIYTMyv1+U63Lrqp3F/9o693sG0J++taP
3Tf15WH9cWPb9Y0bob3pun179v238vraNbPnZPbsyXezHlbtrurZc/JumNBV3cb57/966zY1zUU/aCC5
K9fdaSQhs125vR96HEfI7NBsPLRf+vnFRf/1ybre3o9w27Lqx7SH/b5uuv++H9MPeX9cc1Ot3a51D65J
8qbeN5XryuZ2PPOta9qq3vmRT/2XHj67qVvPwKzcV0/GZM3elK3Ly+7Gf74YQPuyu2nvJXPx9unFrwfX
3J6Tsx8w3v3tBVhe38v/CKO5ePt0doL89O398Paw3ZY9zhlrXNk5snO/kWGds2H13jW9wsWlH/oD/N66
9aGpuluw8ruz/3vDOXQ3dVP9q8cEhvYDBqFcll3Zum42+vrT2V/vJ1nZl025dZ1rIP+AijsVvqkvz3no
v1W7h7407tdD1TjPftccHPjaG3450sTxy7837srj/LeLS3dV7SrPenuxba97GWr368G1XTvm9f2jvDau
3de71rVgydmz775DVMwuXbtuqv1R6jPI2ucTv99U7qNov/vf+5GXnEz64l0zCOJ1dfn+3L6v3Vczb+3a
w6ZrSX1FynVXvXWkbshh5/6592p91OSH6X9Wm+9u973Nt11T7a6hAdx5xL0Spv3CR6RH/OLPYrm9Mr/I
cNuuceX2dzJW0yMnjesOza4l5WZDzggnW9e25bVryb6p165t3SV5c0tAyiQTVjzgnX0tffxlEEK1uyYn
PH/9QynpcH3tRinyq6ZMM6Dvc+bRg1tyVTekJIVOP6SHI11/3vR15KAayr/HNPSHiQD7pu7qN4crvt13
t59hW6da7gzpfY0Wll1pXMfKzl2fB/2T8d3F4frNP9z6Xv33ZS+Yehqwb7wFdRWQ2sw1Td1AQT6stK7s
DmcJ+/2ZWt6Wm4ODmO7oLZumHNvUrOrcFqrwQ2sfORsJ/TEXPsqTr2/qT5JiP+H/UHZDyzEpuWPG/ThG
s3JTrav60H4St/ezHmV5faeE/zc9f/vV5D8pyHMCJmQI6h/vxaVPMs/JrJBzqZbyNaOWx0qvztrF3WE7
iskPj/WZI1DaCiVf/8c5NNeKJUowjsCpiCAwfvr9GBLTLEiFjCF0NQb4IVwLgE4WobBgZK40JjCiEg3U
1PJQj2GGvxwDeIhAVgWUMTUGZiIdAxZCpVxCoSw5zZU0YyCNY82NEQswOBKSU23BYAQorMqpthBpYZXm
ORWAw6BYGZ6mjAI4o5oVqS00IILlCg9UcsEhHUwVOZxbLMaAUHDDgaD4gieCpWBqQi1LAsrmALwKtAgR
8WnxstArBM6EFAsKrCFTVmm2QiuqKNKKhsB0uI4Q2lyweZEDY1I0pMYIY6FBhZAAqwvI1kJIS2OOVlrS
GHqb12umLDIWCjwoEMpylkDgK6MiCy1CpIDtSGlurF6h2bECGpUUm0xUaCkwOIHaiQDJ2dxCCLcUmEue
rAxniLBcK8CD5lZonqgMkGESkeco4ASFEZIb4EQvVEDDBQptjGrOdZ5CzTKVvlABcmIqQoQ2Qo75QgWG
Uw1VJgttELWamwLxxRKVQgOyPGUqywrLUbjy602xBWCJSEMaFAaKQIvMzEWaQrdlc0RtQuHyoS5iA8O2
hbav+SsYjAuWcIpcLI0SqjMYpbXmEsYoLcAwKqGDsQQuEGtV5IBciCdVsQCIglTF0HaVVTg3SW6hekNo
HUYxNMxkBkYWKk1KrQKTeUZhemJUhwZJ9z+fPgPiCYvU8rBAOSgRxioYIJhKaZhBOTCV4kIgTDnVEpmK
N2IFE0RobOEzBwDLOBUmgaqi4RQOHwYegi+VBuF4/nfEckplXNAYcBEoNYf6FIuphTSPAKc5D1EUQ8va
AilzKeYCKQ6Kxkq+BPJacFnAiFxkEHumFrDIygojGExCoVAwhYUcgOwCkQWtS+UWum0YpjAI5XpSeTwN
ufZBC0Rcmol0hbxl/vec4ogwwDhHJSOPJZUMGLiR8ykkw4qAN8a4MchFAk4LWIgGKlyhso6lyibIPyJq
EpkJA/Txgi95CpcyS5GZQiBaTYJKeMoYTL0BlcC0g1gKueCwwkkVTIORkNj8IiFxwkx4GPOokMAThDSF
pqh4nlo9K2xBU4zih0JZGDF8KpaMI5sbKgVMnbGKzWGIpS8nkY6BhYkUrEeC4AcYgl8UEo97lWAYU2xu
cQxXUcRRyS18IQOCLTcG1euh8OFXFdA+ipCmKdcxsKZIY7JStYxgshQWj3tBMSx7iWELHkMdLHjMLdUC
wpeISYyOhgmECJj0Jnmlqc1gdPY5Twsr4Hxjk4zC6GiFyWBwzyeqMxx/mTcmAJuLMOE0tQksVFTKjYWj
I2pFDCN9TmFWjlKQNHJT6JjDIBJyGuG6OOQSFeYhz/suFsaWkGsooFDQgCP3DIVROoTs8FykPDeALP6S
ayZgURpzDa1MxJoiU/HFD8xfCdcBrA1SYSxMfQJWzRPFsZDMZ+CJABZxbQXsPDMup/QbCpjClLYJIJvK
l4KjLRAeCiByDcJWnqxswjXN4VSzgh0IDUPB8LaP0VNEG/6yoCnaazEp5zmM0MKaTM05rCMszQTcIjEr
GWrUytlkpRVsrpZcxIlNFbTZiGdTBE/BqLa4d6baMk1R88xpiAvIQOhwSS1EkRevXkHlMypDDNNhTCG3
jGqrcLeCnJOJGG3uqDTlDG3YZNC4Q02XQJpawOYt5pLTVMGgmRfBRD1eCAuJ8TW35kxBX82pRZVVRue8
L24gOIakS87DdKKW95EP1ny+DUtgo64pigk6R+2JiAQEaZrjPsAwzblcwqRrLM1yaMQcarqvpbEBLJWa
sLVEBQHMX3meClyVcGnxdghnytAIhZCY6hAWxYPi8JaihzMhOQzzQlquBcxwKZWhYRTuFWqeqRDqGZPr
pWJQejQ5as2sWk2MA+ZOLaqwQ7FQGknOJhIaXEa1FmjZnGpYMhoukRB8x2Fg3ApDRI1vnI1dwQBBw0jT
AhaYCu4B+9jCkNPnsOWjKZZzBtP7xIImpxmErJaoH2Mh2hhBvYUXG6y4fCstgLMEWs2hWeJ2SGTx1Faa
kDbFjUmqGJ0AZzSmr4SExibpFI5caVQPTW7m4Zl0ZZUpdAQVjOVIfyioFgUqLGEcZbAkD+E+VCBiKkVG
UwN9MMfVyYLjYjjn6LghXChGYb/MVJb5UhFWlbFamFxJA91CZH3BBgN2ymMoXZZQPXH0kworGPSCtIA1
W2E0N6pAfo62t2lOtc3wxoVmCcgDgTTYi3KtgAHT1GqexrD/tgnHLUMRhgmCMmp9YQrdOtG+VkX9tZBw
b0iYFHrtiyKkaJ0UBktWpGjvBLVrmqcixtsLklqVwSRlrFa4jAjEFFBZKqFtJTwT+JAi5irWNE8weKJk
oZpKpVGsy5OVQWbU0ysRG0vulQd3a5nA1X/+LDcJtAdGLUX7xKbQC45cqchRCZ6onMsY9TdKpj5yocU8
ZCos4ZH96SGFHATU8IDCPf9AsBXDR7iBClcTR0uBeomGzumKzpF2uU45PI9jqchA3ca0YHPYKUUiNnN0
7BClq0iYiV01pSzmK9LcJBEuaWnGJ6AqjWDtog3HMuzB8DBBThCbUW0FTdHGYCDwCYukBp0dqnSV5ciK
+3IY89ofiiJac60CapBwBeOJYnNooVqhHWBdxLA21QU0YrSvZWhqsYQNKwJQYPZSCxSFGdDAHXIj1XJi
XFLAwI4zsVnCo55cyDhXUAWWSwk3hxa+7VphaS9pOocQyzUmWzG0GWR8rQHW6QOeYOZ7uGflkzzu3Km0
YiF0gfIkhUwxmnFNcQXJJk7/sxxtuedcizzhGiVvX4TyheCg2QyppT7GACg387wIMHABW3/ftKNcrXmu
HpAXQLlAu65Zjs6aFwrWpH2bg3h/QRcUQwzTAm4/p0IWL6Hvo2aDmSL3qod+G+DtpPx7WHVaBm1Achtw
GEW4ReyqKMI3e0LAWT5NXP4M5h8fZFDH5GvnAFWdOdyv1JzjmriQAoiusMIXgTDELgKKTrFYKnIUXpc8
CLkRsEITBvCyFDJUsJqf2DUKF1ziuxA00oLBbWOhraYL2APTwlhNU7ihF0jgEUEY26n5gXg1Be6xIihD
2xOMSgr1zagWQcDR9QddoC1ZTo3lhVaw3Z+CRRNnPrHmHMG4pUt4QSxRFjIjLIWnci9oju4D8ZcC3uby
Xd5EJy/pQsQTAdX4MIn0aXIA6Hd6PFpQ/YhwUkHzKegJdgQN7yjOrgger1t+0i3LuzmnL4/cscS3f/uv
0/cbpy8On+aNbk2OL8M+jrS/DDuJajtx3fRxfPfXTe+RfnP+Lxa0HYT76HXMk7D/RjKaCiZUYZ6Tp6Rq
SeNa17x1lz/uno3//P78z3tUZzc7j7c0xaveKu+H4HudD4wctnoGckb+yRL1oI2Nnkh9iqGNJj5qbY3b
b37P67wnaj7p4vZ44mdxf/s471/vAvfZy6cP3+MeyfJE9G/VZkOqS7frqqtb0t04ckRIyqvONaQ9vNlW
bTv26UcEd3wY+BmyO878iEC12bwp1798Fsv2xpE7BOTQbEgvg8atXfXW9SK4qnalh4xf8ZxLfV3vOrfr
PsF6x8sP00lXk1PMdV9m7YVO2ZGoaWs/o/7uyd5nUn96L9TVw+PD/uGQK9c34O3QZ7BxHnYf5eMDjxce
46HQafsZ4n/obcSnxJV2CO6f7ByjV5wPukavkNftEDE+NsD0K/yOUebo2USEval0Tbn+ZXh6Bx8Rjxcd
nlu9fsBeHzOi3+nFC5DXJ2vyOO/xIFdv9xvXPfiy6U1db1y5G4v8qm62ZXf+eboU+2r558p165vX6/ry
QTqrXeeuXfMQndWu+/7ZA3bjg/FN1+3JYM/EL3MXqS/JVVNv+3jdV1AN+e3G7UhPj7et0sf2aZo39Xqy
wv0oWz4t4Ff26WPIE+6SVDtSksZdVo1b97G9u6laclrscbP6eIt62Ga+UA/T1fbwaPeL3ebs5eUX8PiH
iwZnqfeLVPfBkuKDNulNsVx3h3JzXlocn1g/IacntJMiuHHlpUPBALxeHQaRX9ztRf/ckuzLqmkBOdMc
j0kdEIGZ5eVlr69yk09L54OK+Yzy4IPS9E7dnVVp63Lr+mDzLdnV3ZQ0Jy1j/GD3zDKAaH30qtaORHVN
3v24GwY1+zUJyuYv13V9vXFP7nA96ZH99fS8/oHv/zWgef/j7sedr3ZeGCVJ4/a+2dx1fTjqE/HP/fCf
fR/q/P+GgYP+yM/v3v/85CyJPZbcKLl2O9dU6yOuY9Qg3U3Zkdv6QNbljjTub4fWeess39bVJRl8dXdN
Lg/DT/P4xng0v/Vx9bY+NITmon1CKOlu99W63BD3z9JnSE99VxOPtupI2Z53MT/ufLnR/z3UEn6y639J
Y+fxka3rburLJySqG1Lt2q7crd3z4y/inF5M3/1eRjj1chr+TAYOAH7EfYabXW3q3/oSde29k/mIffo2
QlaMfw7If7v/8Z2LAW0PPVNG/Yv7uHn9yPuJ7brew4g0/n0Pr+Ge4uHXGkh47AJmY+P3Qvvm/f8GAAD/
/7AprrKsSQAA
`,
	},

	"/schema.graphql": {
		local:   "static/schema.graphql",
		size:    6963,
		modtime: 1515182423,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
