package fanunmarshal

import (
	"fmt"
	"sort"
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
	}

	if len(res) != len(testdata) {
		t.Errorf("expected %d unique elements, got %d", len(testdata), len(res))
	}

}

func ExampleIFanUnMarshal_UnMarshalSlice_example2() {

	// setup receiver struct
	type SomeData struct {
		Data  string `json:"data"`
		Other int    `json:"other"`
	}

	// small tiny example
	testdata := [][]byte{
		[]byte(`{ "data": "1", "other": 42}`),
		[]byte(`{ "data": "2", "other": 1337}`),
		[]byte(`{ "data": "3", "other": 161803398875}`),
	}

	// we expect slice of SomeData
	var expected = SomeData{}

	data := New().
		WithWorkers(10).
		WithUseJsonIter().
		UnMarshalSlice(testdata, &expected)

	/** sort it for comparation, just for the example to pass
		take notice that order of the returned slice is not in order you may expect it to be
	**/
	sort.Slice(data, func(i, j int) bool {
		return data[i].(*SomeData).Other < data[j].(*SomeData).Other
	})

	// data is the slice returned from the UnMarshalSlice method
	for _, d := range data {
		fmt.Printf("%+v\n", d)
	}

	// Output: &{Data:1 Other:42}
	// &{Data:2 Other:1337}
	// &{Data:3 Other:161803398875}
}

func ExampleIFanUnMarshal_UnMarshalChan_example1() {

	// setup receiver struct
	type SomeData struct {
		Data  string `json:"data"`
		Other int    `json:"other"`
	}

	// small tiny example
	testdata := [][]byte{
		[]byte(`{ "data": "1", "other": 42}`),
		[]byte(`{ "data": "2", "other": 1337}`),
		[]byte(`{ "data": "3", "other": 161803398875}`),
	}

	// we expect slice of SomeData
	var (
		// setup expected what we receive in the output channel
		expected = SomeData{}
		// setup instance
		fm = New().
			WithWorkers(3).        // length of data
			WithUseJsonIter().     // use jsoniter lib
			DisableAutoScaleDown() // because we're not sending the length of the data into UnMarshalChan we need to disable the scaledown

			// setup our input channel
		pipe = fm.MakeChan(testdata)

		// here we send in a nil as the length of the data, otherwise a pointer to int of length
		outputChan = fm.UnMarshalChan(pipe, &expected, nil)
	)

	/**
		take notice that order of the channel is not in order you may expect it to be
	**/
	var summed = 0
	// 42 + 1337 + 161803398875 = 161803400254
	for d := range outputChan {
		if data, ok := d.(*SomeData); ok {
			summed += data.Other
		}
	}
	fmt.Printf("%d\n", summed)

	// Output: 161803400254
}
