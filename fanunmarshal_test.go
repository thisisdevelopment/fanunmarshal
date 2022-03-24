package fanunmarshal

import (
	"testing"
)

type SomeData struct {
	Data  string `json:"data"`
	Other int    `json:"other"`
}

func initTestData() (testdata [][]byte) {
	testdata = append(testdata, []byte(`{ "data": "1", "other": 42}`))
	testdata = append(testdata, []byte(`{ "data": "2", "other": 42}`))
	testdata = append(testdata, []byte(`{ "data": "3", "other": 42}`))
	testdata = append(testdata, []byte(`{ "data": "4", "other": 42}`))
	testdata = append(testdata, []byte(`{ "data": "5", "other": 42}`))
	testdata = append(testdata, []byte(`{ "data": "6", "other": 42}`))
	return testdata
}

func TestNewWithSlice(t *testing.T) {

	var testdata = initTestData()
	var expected = SomeData{}

	data := New().
		WithWorkers(10).
		WithUseJsonIter().
		UnMarshalSlice(testdata, &expected)

	res := map[string]bool{}
	for _, d := range data {
		a := d.(*SomeData)
		res[a.Data] = true
		// spew.Dump(d)
	}

	if len(res) != len(testdata) {
		t.Errorf("expected %d unique elements, got %d", len(testdata), len(res))
	}

}

func TestNewWithChannel(t *testing.T) {

	var testdata = initTestData()
	var expected = SomeData{}

	fm := New().
		WithWorkers(10).
		WithUseJsonIter().
		DisableAutoScaleDown()

	pipe := fm.MakeChan(testdata)
	outputChan := fm.UnMarshalChan(pipe, &expected, nil)

	res := map[string]bool{}
	for d := range outputChan {
		a := d.(*SomeData)
		res[a.Data] = true
		// spew.Dump(d)
	}

	if len(res) != len(testdata) {
		t.Errorf("expected %d unique elements, got %d", len(testdata), len(res))
	}

}
