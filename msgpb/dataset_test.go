package msgpb

import (
	"fmt"
	"strings"
	"testing"
)

func testDS(t *testing.T, ds *Dataset, expectNil bool) {
	t.Helper()

	// iterate through each valid dataset type
	for dstID := range DatasetType_name {
		dst := DatasetType(dstID)

		i, err := ds.FieldByType(dst)
		if err != nil {
			t.Error("DatasetByType returned error", err)
		}

		if i == nil && !expectNil {
			t.Error("did not expect nil dataset field")
		}

		if i != nil && expectNil {
			t.Errorf("got unexpected nil dataset field %#v", i)
		}

		if expectNil {
			return
		}

		switch dst {
		case CATEGORIZATION:
			r, ok := i.(*Dataset_Categorization)
			if !ok {
				t.Error("type of i not *Dataset_Categorization")
			}

			if r != ds.Categorization {
				t.Error("t != ds.Categorization")
			}
		case MALICIOUS:
			r, ok := i.(*Dataset_Malicious)
			if !ok {
				t.Error("type of i not *Dataset_Malicious")
			}

			if r != ds.Malicious {
				t.Error("t != ds.Malicious")
			}
		case ECHO:
			r, ok := i.(*Dataset_Echo)
			if !ok {
				t.Error("type of i not *Dataset_Echo")
			}

			if r != ds.Echo {
				t.Error("t != ds.Echo")

			}
		case LANGUAGE:
			r, ok := i.(*Dataset_Language)
			if !ok {
				t.Error("type of i not *Dataset_Language")
			}

			if r != ds.Language {
				t.Error("t != ds.Language")

			}
		default:
			t.Errorf("unexpected dataset type: %s", dst)
		}
	}
}

func TestDatasetByType(t *testing.T) {
	testDS(t, &Dataset{
		Categorization: &Dataset_Categorization{},
		Malicious:      &Dataset_Malicious{},
		Echo:           &Dataset_Echo{},
		Language:       &Dataset_Language{},
	}, false)
}

func TestNilDatasetByType(t *testing.T) {
	testDS(t, &Dataset{}, true)
}

func TestDatasetByTypeErr(t *testing.T) {
	ds := &Dataset{}
	i, err := ds.FieldByType(DatasetType(-1))

	if err == nil {
		t.Error("DatasetByType didn't return error for invalid dataset type")
	}

	if i != nil {
		t.Error("expected DatasetByType to return nil interface when err != nil ")
	}

	e, ok := err.(errInvalidDatasetType)
	if !ok {
		t.Error("error was not of type errInvalidDatasetType")
	}

	const errMsg0 = "invalid dataset type: -1"
	if e.Error() != errMsg0 || err.Error() != errMsg0 {
		t.Error("error did not have expected message")
	}

	ds = nil
	i, err = ds.FieldByType(CATEGORIZATION)

	if err == nil {
		t.Error("DatasetByType didn't return error for nil dataset")
	}

	if i != nil {
		t.Error("expected DatasetByType to return nil interface when err != nil ")
	}

	if err != ErrNilDataset {
		t.Error("unexpected error type")
	}
}

func ExampleDataset_FieldByType() {
	ds := &Dataset{
		Categorization: &Dataset_Categorization{},
	}

	i, _ := ds.FieldByType(CATEGORIZATION)

	c := i.(*Dataset_Categorization)

	fmt.Printf("c == ds.Categorization => %v\n", c == ds.Categorization)
	// Output: c == ds.Categorization => true
}

func testDatasetType(t *testing.T, dstValue int32, dstName string) {
	dst, err := NewDatasetType(strings.ToLower(dstName))
	if err != nil {
		t.Error("unexpected error")
	}

	if DatasetType(dstValue) != dst {
		t.Error("unexpected dataset type value")
	}
}

func TestNewDatasetType(t *testing.T) {
	for dstValue, dstName := range DatasetType_name {
		testDatasetType(t, dstValue, dstName)
		testDatasetType(t, dstValue, strings.ToLower(dstName))
		testDatasetType(t, dstValue, strings.ToUpper(dstName))
	}

	dst, err := NewDatasetType("invalid dataset type name")
	if err == nil {
		t.Error("NewDatasetType should have returned an error")
	}

	if err != ErrInvalidDatasetType {
		t.Error("NewDatasetType returned the wrong error")
	}

	if dst != DatasetType(-1) {
		t.Error("NewDatasetType returned wrong response")
	}
}

func TestMergeDatasets(t *testing.T) {
	d1 := &Dataset{
		Categorization: &Dataset_Categorization{
			Value: []Category{1},
		},
	}
	d2 := &Dataset{
		Malicious: &Dataset_Malicious{
			Category: []Category{MAL_4},
		},
	}

	if d1.Categorization == nil {
		t.Error("d1.Categorization should not be nil")
	}

	if d2.Malicious == nil {
		t.Error("d2.Malicious should not be nil")
	}

	if d2.Categorization != nil {
		t.Error("d2.Categorization should be nil")
	}

	if d1.Malicious != nil {
		t.Error("d1.Malicious should be nil")
	}

	d3, err := MergeDatasets(d1, d2)
	if err != nil {
		t.Errorf("MergeDatasets returned an error: %s", err)
	}

	if d3.Categorization == nil {
		t.Error("d3.Categorization should not be nil")
	}

	if d3.Categorization.Value[0] != 1 {
		t.Error("d3.Categorization.Values[0] should equal 1")
	}

	if d3.Malicious == nil {
		t.Error("d3.Malicious should not be nil")
	}

	if d3.Malicious.Category[0] != MAL_4 {
		t.Error("d3.Malicious.Category wasn't MAL_4")
	}
}

func TestMergeDatasetsOneEmpty(t *testing.T) {
	d1 := &Dataset{}
	d2 := &Dataset{
		Malicious: &Dataset_Malicious{
			Category: []Category{MAL_4},
		},
	}

	if d1.Categorization != nil {
		t.Error("d1.Categorization should be nil")
	}

	if d2.Malicious == nil {
		t.Error("d2.Malicious should not be nil")
	}

	if d2.Categorization != nil {
		t.Error("d2.Categorization should be nil")
	}

	if d1.Malicious != nil {
		t.Error("d1.Malicious should be nil")
	}

	d3, err := MergeDatasets(d1, d2)
	if err != nil {
		t.Errorf("MergeDatasets returned an error: %s", err)
	}

	if d3.Categorization != nil {
		t.Error("d3.Categorization should be nil")
	}

	if d3.Malicious == nil {
		t.Error("d3.Malicious should not be nil")
	}

	if d3.Malicious.Category[0] != MAL_4 {
		t.Error("d3.Malicious.Category wasn't MAL_4")
	}
}

func TestMergeDatasetsBothEmpty(t *testing.T) {
	d1 := &Dataset{}
	d2 := &Dataset{}

	if d1.Categorization != nil {
		t.Error("d1.Categorization should be nil")
	}

	if d1.Malicious != nil {
		t.Error("d1.Malicious should be nil")
	}

	if d2.Categorization != nil {
		t.Error("d2.Categorization should be nil")
	}

	if d2.Malicious != nil {
		t.Error("d2.Malicious should be nil")
	}

	d3, err := MergeDatasets(d1, d2)
	if err != nil {
		t.Errorf("MergeDatasets returned an error: %s", err)
	}

	if d3.Categorization != nil {
		t.Error("d3.Categorization should be nil")
	}

	if d3.Malicious != nil {
		t.Error("d3.Malicious should be nil")
	}
}

func TestMergeDatasetsOneNil(t *testing.T) {
	var d1 *Dataset
	d2 := &Dataset{
		Malicious: &Dataset_Malicious{
			Category: []Category{MAL_4},
		},
	}

	if d2.Malicious == nil {
		t.Error("d2.Malicious should not be nil")
	}

	if d2.Categorization != nil {
		t.Error("d2.Categorization should be nil")
	}

	d3, err := MergeDatasets(d1, d2)
	if err != nil {
		t.Errorf("MergeDatasets returned an error: %s", err)
	}

	if d3.Categorization != nil {
		t.Error("d3.Categorization should be nil")
	}

	if d3.Malicious == nil {
		t.Error("d3.Malicious should not be nil")
	}

	if d3.Malicious.Category[0] != MAL_4 {
		t.Error("d3.Malicious.Category wasn't MAL_4")
	}
}

func TestMergeDatasetsBothNil(t *testing.T) {
	var d1 *Dataset
	var d2 *Dataset
	d3, err := MergeDatasets(d1, d2)
	if err != nil {
		t.Errorf("MergeDatasets returned an error: %s", err)
	}

	if d3.Categorization != nil {
		t.Error("d3.Categorization should be nil")
	}

	if d3.Malicious != nil {
		t.Error("d3.Malicious should be nil")
	}
}
