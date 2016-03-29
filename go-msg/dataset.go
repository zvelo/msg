package msg

import (
	"fmt"
	"reflect"
	"strings"
)

// errInvalidDataSetType indicates an invalid int was used as a DataSetType enum
// value
type errInvalidDataSetType DataSetType

func (e errInvalidDataSetType) Error() string {
	return fmt.Sprintf("invalid dataset type: %d", int32(e))
}

// Various errors
var (
	ErrNilDataSet   = fmt.Errorf("dataset was nil")
	ErrInvalidField = fmt.Errorf("dataset type does not exist in dataset definition")
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
