package msg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

// errInvalidDataSetType indicates an invalid int was used as a DataSetType enum
// value
type errInvalidDataSetType DataSetType

func (e errInvalidDataSetType) Error() string {
	return fmt.Sprintf("invalid dataset type: %d", int32(e))
}

// Various errors
var (
	ErrNilDataSet         = fmt.Errorf("dataset was nil")
	ErrInvalidField       = fmt.Errorf("dataset type does not exist in dataset definition")
	ErrInvalidDataSetType = fmt.Errorf("invalid dataset type")
)

// FieldByType returns one of the field values of a DataSet based on a
// given dsType. It determines which value to return by doing a case insensitive
// comparison of DataSetType.String() and the field name of DataSet. It
// returns an interface{} that can be type asserted into the appropriate message
// type.
func (m *DataSet) FieldByType(dsType DataSetType) (interface{}, error) {
	name, ok := DataSetType_name[int32(dsType)]
	if !ok {
		return nil, errInvalidDataSetType(dsType)
	}

	if m == nil {
		return nil, ErrNilDataSet
	}

	switch dsType {
	case DataSetType_KEYWORD, DataSetType_SENTIMENT:
		return m.Keyword, nil
	}

	name = strings.ToLower(name)
	v := reflect.ValueOf(*m).FieldByNameFunc(func(val string) bool {
		return strings.ToLower(val) == name
	})

	if v.IsValid() {
		if v.IsNil() {
			return nil, nil
		}

		return v.Interface(), nil
	}

	// NOTE: if this is reached, it indicates a problem where a valid
	// DataSetType was provided, but DataSet has no cooresponding
	// (case-insensitive) field name
	return nil, ErrInvalidField
}

// NewDataSetType returns the cooresponding DataSetType given a string. It is
// case insensitive.
func NewDataSetType(name string) (DataSetType, error) {
	name = strings.ToLower(name)

	for dst, dstName := range DataSetType_name {
		if name == strings.ToLower(dstName) {
			return DataSetType(dst), nil
		}
	}

	return DataSetType(-1), ErrInvalidDataSetType
}

// MergeDatasets returns the DataSet resulting from merging `ds1` and `ds2`
func MergeDatasets(d1, d2 *DataSet) *DataSet {
	if d1 == nil {
		d1 = &DataSet{}
	}

	if d2 == nil {
		d2 = &DataSet{}
	}

	// Start with a clone
	ret := proto.Clone(d1).(*DataSet)

	// overwrite old values with non-nil new values
	val1 := reflect.ValueOf(ret)
	val2 := reflect.ValueOf(d2)
	numFields := val1.Elem().NumField()
	for i := 0; i < numFields; i++ {
		fieldVal := val2.Elem().Field(i)
		if !fieldVal.IsNil() {
			val1.Elem().Field(i).Set(fieldVal)
		}
	}

	return ret
}
