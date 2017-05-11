// Code generated by protoc-gen-gogo.
// source: zvelo/msg/seed_result.proto
// DO NOT EDIT!

package msg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SeedResult struct {
	Url     string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Dataset *DataSet `protobuf:"bytes,2,opt,name=dataset" json:"dataset,omitempty"`
}

func (m *SeedResult) Reset()                    { *m = SeedResult{} }
func (*SeedResult) ProtoMessage()               {}
func (*SeedResult) Descriptor() ([]byte, []int) { return fileDescriptorSeedResult, []int{0} }

func (m *SeedResult) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SeedResult) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type SeedResults struct {
	Values []*SeedResult `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *SeedResults) Reset()                    { *m = SeedResults{} }
func (*SeedResults) ProtoMessage()               {}
func (*SeedResults) Descriptor() ([]byte, []int) { return fileDescriptorSeedResult, []int{1} }

func (m *SeedResults) GetValues() []*SeedResult {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*SeedResult)(nil), "zvelo.msg.SeedResult")
	proto.RegisterType((*SeedResults)(nil), "zvelo.msg.SeedResults")
}
func (this *SeedResult) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SeedResult)
	if !ok {
		that2, ok := that.(SeedResult)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SeedResult")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SeedResult but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SeedResult but is not nil && this == nil")
	}
	if this.Url != that1.Url {
		return fmt.Errorf("Url this(%v) Not Equal that(%v)", this.Url, that1.Url)
	}
	if !this.Dataset.Equal(that1.Dataset) {
		return fmt.Errorf("Dataset this(%v) Not Equal that(%v)", this.Dataset, that1.Dataset)
	}
	return nil
}
func (this *SeedResult) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SeedResult)
	if !ok {
		that2, ok := that.(SeedResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if !this.Dataset.Equal(that1.Dataset) {
		return false
	}
	return true
}
func (this *SeedResults) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SeedResults)
	if !ok {
		that2, ok := that.(SeedResults)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SeedResults")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SeedResults but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SeedResults but is not nil && this == nil")
	}
	if len(this.Values) != len(that1.Values) {
		return fmt.Errorf("Values this(%v) Not Equal that(%v)", len(this.Values), len(that1.Values))
	}
	for i := range this.Values {
		if !this.Values[i].Equal(that1.Values[i]) {
			return fmt.Errorf("Values this[%v](%v) Not Equal that[%v](%v)", i, this.Values[i], i, that1.Values[i])
		}
	}
	return nil
}
func (this *SeedResults) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SeedResults)
	if !ok {
		that2, ok := that.(SeedResults)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if !this.Values[i].Equal(that1.Values[i]) {
			return false
		}
	}
	return true
}
func (this *SeedResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&msg.SeedResult{")
	s = append(s, "Url: "+fmt.Sprintf("%#v", this.Url)+",\n")
	if this.Dataset != nil {
		s = append(s, "Dataset: "+fmt.Sprintf("%#v", this.Dataset)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SeedResults) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&msg.SeedResults{")
	if this.Values != nil {
		s = append(s, "Values: "+fmt.Sprintf("%#v", this.Values)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSeedResult(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SeedResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSeedResult(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.Dataset != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSeedResult(dAtA, i, uint64(m.Dataset.Size()))
		n1, err := m.Dataset.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SeedResults) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedResults) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSeedResult(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64SeedResult(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32SeedResult(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSeedResult(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SeedResult) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSeedResult(uint64(l))
	}
	if m.Dataset != nil {
		l = m.Dataset.Size()
		n += 1 + l + sovSeedResult(uint64(l))
	}
	return n
}

func (m *SeedResults) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovSeedResult(uint64(l))
		}
	}
	return n
}

func sovSeedResult(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSeedResult(x uint64) (n int) {
	return sovSeedResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SeedResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SeedResult{`,
		`Url:` + fmt.Sprintf("%v", this.Url) + `,`,
		`Dataset:` + strings.Replace(fmt.Sprintf("%v", this.Dataset), "DataSet", "DataSet", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SeedResults) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SeedResults{`,
		`Values:` + strings.Replace(fmt.Sprintf("%v", this.Values), "SeedResult", "SeedResult", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSeedResult(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SeedResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dataset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSeedResult
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dataset == nil {
				m.Dataset = &DataSet{}
			}
			if err := m.Dataset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SeedResults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedResults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedResults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSeedResult
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &SeedResult{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSeedResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSeedResult
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSeedResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSeedResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSeedResult
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSeedResult
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSeedResult(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSeedResult = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSeedResult   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("zvelo/msg/seed_result.proto", fileDescriptorSeedResult) }

var fileDescriptorSeedResult = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xae, 0x2a, 0x4b, 0xcd,
	0xc9, 0xd7, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x4e, 0x4d, 0x4d, 0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x4b, 0xea, 0xe5, 0x16, 0xa7, 0x4b,
	0x89, 0x23, 0xd4, 0xa5, 0x24, 0x96, 0x24, 0x16, 0xa7, 0x42, 0xd5, 0x28, 0xf9, 0x70, 0x71, 0x05,
	0xa7, 0xa6, 0xa6, 0x04, 0x81, 0xf5, 0x09, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x3a, 0x5c, 0xec, 0x50, 0x0d, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x70, 0x53, 0xf5, 0x5c, 0x12, 0x4b, 0x12, 0x83, 0x53, 0x4b, 0x82,
	0x60, 0x4a, 0x94, 0x6c, 0xb8, 0xb8, 0x11, 0xa6, 0x15, 0x0b, 0xe9, 0x72, 0xb1, 0x95, 0x25, 0xe6,
	0x94, 0xa6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0x22, 0xe9, 0x45, 0xa8, 0x0b,
	0x82, 0x2a, 0x72, 0xb2, 0xbb, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x07, 0x0f, 0xe5,
	0x18, 0x3f, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x44, 0xf1, 0x40, 0xcc, 0xcb, 0x04, 0xfb, 0x2c, 0x89,
	0x0d, 0xec, 0x25, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x8e, 0x35, 0x44, 0x15, 0x01,
	0x00, 0x00,
}
