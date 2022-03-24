- [README IS WIP](#readme-is-wip)
- [Introduction](#introduction)
- [Basic usage with slice response](#basic-usage-with-slice-response)
- [Contributing](#contributing)
- [License](#license)

# README IS WIP 

# Introduction
**fanunmarshal** is a concurrent unmarhaller

use with slices of byte slices **[][]byte**, for example for data coming from Redis using MGet

# Basic usage with slice response
```
import github.com/thisisdevelopment/fanunmarshal


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


func withSlice() {
    testdata := initTestData() 
	var expected = SomeData{}
	data := fanunmarshal.New().
		WithWorkers(10).
		WithUseJsonIter().
		UnMarshalSlice(testdata, &expected)

	res := map[string]bool{}
	for _, d := range data {
		somedata := d.(*SomeData)
	}
}

func withChannel() {
    var testdata = initTestData()
	var expected = SomeData{}

	fm := New().
		WithWorkers(10).
		WithUseJsonIter().
		DisableAutoScaleDown()

	pipe := fm.MakeChan(testdata)
	outputChan := fm.UnMarshalChan(pipe, &expected, nil)

    /** OR
    	fm := fanunmarshal.New().
	    	WithWorkers(10).
		    WithUseJsonIter()

        dataLength := len(testData)
    	outputChan := fm.UnMarshalChan(pipe, &expected, &dataLength)

    **/ 



	res := map[string]bool{}
	for d := range outputChan {
		somedata := d.(*SomeData)
	}

}
```







# Contributing 
You can help to deliver a better fanunmarshaller, check out how you can do things [CONTRIBUTING.md](CONTRIBUTING.md)

# License 
© [This is Development BV](https://www.thisisdevelopment.nl), 2022~time.Now()
Released under the [MIT License](https://github.com/thisisdevelopment/fanunmarshal/blob/master/LICENSE)
