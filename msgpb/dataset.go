package msgpb

import (
	"fmt"
	"reflect"
	"strings"

	proto "github.com/gogo/protobuf/proto"
)

// errInvalidDatasetType indicates an invalid int was used as a DatasetType enum
// value
type errInvalidDatasetType DatasetType

func (e errInvalidDatasetType) Error() string {
	return fmt.Sprintf("invalid dataset type: %d", int32(e))
}

// Various errors
var (
	ErrNilDataset         = fmt.Errorf("dataset was nil")
	ErrInvalidField       = fmt.Errorf("dataset type does not exist in dataset definition")
	ErrInvalidDatasetType = fmt.Errorf("invalid dataset type")
)

// FieldByType returns one of the field values of a Dataset based on a
// given dsType. It determines which value to return by doing a case insensitive
// comparison of DatasetType.String() and the field name of Dataset. It
// returns an interface{} that can be type asserted into the appropriate message
// type.
func (m *Dataset) FieldByType(dsType DatasetType) (interface{}, error) {
	name, ok := DatasetType_name[int32(dsType)]
	if !ok {
		return nil, errInvalidDatasetType(dsType)
	}

	if m == nil {
		return nil, ErrNilDataset
	}

	v := reflect.ValueOf(*m).FieldByNameFunc(func(val string) bool {
		return strings.EqualFold(name, val)
	})

	if v.IsValid() {
		if v.IsNil() {
			return nil, nil
		}

		return v.Interface(), nil
	}

	// NOTE: if this is reached, it indicates a problem where a valid
	// DatasetType was provided, but Dataset has no corresponding
	// (case-insensitive) field name
	return nil, ErrInvalidField
}

// NewDatasetType returns the corresponding DatasetType given a string. It is
// case insensitive.
func NewDatasetType(name string) (DatasetType, error) {
	for dst, dstName := range DatasetType_name {
		if strings.EqualFold(name, dstName) {
			return DatasetType(dst), nil
		}
	}

	return DatasetType(-1), ErrInvalidDatasetType
}

func (m *Dataset) Clone() (*Dataset, error) {
	data, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	var ret Dataset
	if err = proto.Unmarshal(data, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// MergeDatasets returns the Dataset resulting from merging `ds1` and `ds2`
func MergeDatasets(d1, d2 *Dataset) (*Dataset, error) {
	if d1 == nil {
		d1 = &Dataset{}
	}

	if d2 == nil {
		d2 = &Dataset{}
	}

	// Start with a clone
	ret, err := d1.Clone()
	if err != nil {
		return nil, err
	}

	// overwrite old values with non-nil new values
	val1 := reflect.ValueOf(ret)
	val2 := reflect.ValueOf(d2)
	numFields := val1.Elem().NumField()
	for i := 0; i < numFields; i++ {
		fieldVal := val2.Elem().Field(i)
		if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() {
			val1.Elem().Field(i).Set(fieldVal)
		}
	}

	return ret, nil
}
