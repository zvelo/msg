package msg

import (
	"fmt"
	"testing"
)

func testDS(t *testing.T, ds *DataSet) {
	// iterate through each valid dataset type
	for dstID := range DataSetType_name {
		dst := DataSetType(dstID)

		i, err := ds.FieldByType(dst)
		if err != nil {
			t.Error("DataSetByType returned error", err)
		}

		switch dst {
		case DataSetType_CATEGORIZATION:
			r, ok := i.(*DataSet_Categorization)
			if !ok {
				t.Error("type of i not *DataSet_Categorization")
			}

			if r != ds.Categorization {
				t.Error("t != ds.Categorization")
			}
		case DataSetType_ADFRAUD:
			r, ok := i.(*DataSet_AdFraud)
			if !ok {
				t.Error("type of i not *DataSet_AdFraud")
			}

			if r != ds.Adfraud {
				t.Error("t != ds.Adfraud")
			}
		case DataSetType_MALICIOUS:
			r, ok := i.(*DataSet_Malicious)
			if !ok {
				t.Error("type of i not *DataSet_Malicious")
			}

			if r != ds.Malicious {
				t.Error("t != ds.Malicious")
			}
		default:
			t.Errorf("unexpected dataset type: %s", dst)
		}
	}
}

func TestDataSetByType(t *testing.T) {
	testDS(t, &DataSet{
		Categorization: &DataSet_Categorization{},
		Adfraud:        &DataSet_AdFraud{},
		Malicious:      &DataSet_Malicious{},
	})
}

func TestNilDataSetByType(t *testing.T) {
	testDS(t, &DataSet{})
}

func TestDataSetByTypeErr(t *testing.T) {
	ds := &DataSet{}
	i, err := ds.FieldByType(DataSetType(-1))

	if err == nil {
		t.Error("DataSetByType didn't return error for invalid dataset type")
	}

	if i != nil {
		t.Error("expected DataSetByType to return nil interface when err != nil ")
	}

	e, ok := err.(ErrInvalidDataSetType)
	if !ok {
		t.Error("error was not of type ErrInvalidDataSetType")
	}

	const errMsg0 = "invalid dataset type: -1"
	if e.Error() != errMsg0 || err.Error() != errMsg0 {
		t.Error("error did not have expected message")
	}

	ds = nil
	i, err = ds.FieldByType(DataSetType_CATEGORIZATION)

	if err == nil {
		t.Error("DataSetByType didn't return error for nil dataset")
	}

	if i != nil {
		t.Error("expected DataSetByType to return nil interface when err != nil ")
	}

	if err != ErrNilDataSet {
		t.Error("unexpected error type")
	}
}

func ExampleDataSet_FieldByType() {
	ds := &DataSet{
		Categorization: &DataSet_Categorization{},
	}

	i, _ := ds.FieldByType(DataSetType_CATEGORIZATION)

	c := i.(*DataSet_Categorization)

	fmt.Printf("c == ds.Categorization => %v\n", c == ds.Categorization)
	// Output: c == ds.Categorization => true
}
