package fanunmarshal_test

import (
	"fmt"
	"sort"

	fum "github.com/thisisdevelopment/fanunmarshal"
)

func ExampleFanUnMarshal_UnMarshalSlice() {
	// small tiny example
	testdata := [][]byte{
		[]byte(`{ "data": "1", "other": 42}`),
		[]byte(`{ "data": "2", "other": 1337}`),
		[]byte(`{ "data": "3", "other": 161803398875}`),
	}

	// setup receiver struct
	type SomeData struct {
		Data  string `json:"data"`
		Other int    `json:"other"`
	}

	// we expect slice of SomeData
	var expected = SomeData{}

	data := fum.New().
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

func ExampleFanUnMarshal_UnMarshalChan() {

	// small tiny example
	testdata := [][]byte{
		[]byte(`{ "data": "1", "other": 42}`),
		[]byte(`{ "data": "2", "other": 1337}`),
		[]byte(`{ "data": "3", "other": 161803398875}`),
	}

	// setup receiver struct
	type SomeData struct {
		Data  string `json:"data"`
		Other int    `json:"other"`
	}

	// we expect slice of SomeData
	// setup expected what we receive in the output channel
	var expected = SomeData{}

	// setup instance
	var fm = fum.New().
		WithWorkers(3).        // length of data
		WithUseJsonIter().     // use jsoniter lib
		DisableAutoScaleDown() // because we're not sending the length of the data into UnMarshalChan we need to disable the scaledown

	// setup our input channel
	var pipe = fm.MakeChan(testdata)

	// here we send in a nil as the length of the data, otherwise a pointer to int of length
	var outputChan = fm.UnMarshalChan(pipe, &expected, nil)

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
