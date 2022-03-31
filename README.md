- [README IS WIP](#readme-is-wip)
- [Introduction](#introduction)
- [Basic usage with slice response](#basic-usage-with-slice-response)
- [Benchmarks](#benchmarks)
- [Contributing](#contributing)
- [License](#license)

[![go report card](https://goreportcard.com/badge/github.com/thisisdevelopment/fanunmarshal "go report card")](https://goreportcard.com/report/github.com/thisisdevelopment/fanunmarshal)
[![CircleCI](https://circleci.com/gh/thisisdevelopment/fanunmarshal.svg?style=svg)](https://circleci.com/gh/thisisdevelopment/fanunmarshal)
[![GoDoc](https://godoc.org/github.com/thisisdevelopment/fanunmarshal?status.svg)](https://godoc.org/github.com/thisisdevelopment/fanunmarshal)

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




# Benchmarks

Total json size 1.4mb. 1000 lines

Standard sequencial unmarshalling

```
fanunmarshal on  master [!?] via 🐹 v1.18 took 34s 
❯ go test -benchtime=1s -bench=. -count=10 -cpu=4
goos: darwin
goarch: arm64
pkg: github.com/thisisdevelopment/fanunmarshal

BenchmarkPlainUnMarshal-4                    141           8334686 ns/op
BenchmarkPlainUnMarshal-4                    141           8339650 ns/op
BenchmarkPlainUnMarshal-4                    141           8338255 ns/op
BenchmarkPlainUnMarshal-4                    142           8351871 ns/op
BenchmarkPlainUnMarshal-4                    142           8369517 ns/op
BenchmarkPlainUnMarshal-4                    142           8347529 ns/op
BenchmarkPlainUnMarshal-4                    142           8333782 ns/op
BenchmarkPlainUnMarshal-4                    142           8403951 ns/op
BenchmarkPlainUnMarshal-4                    142           8371589 ns/op
BenchmarkPlainUnMarshal-4                    142           8389560 ns/op
```
Using fanunmarshall 10 workers stdlib json
```
BenchmarkWithLibSlice_stdlib_10-4            336           3573922 ns/op
BenchmarkWithLibSlice_stdlib_10-4            337           3572657 ns/op
BenchmarkWithLibSlice_stdlib_10-4            338           3607899 ns/op
BenchmarkWithLibSlice_stdlib_10-4            336           3579343 ns/op
BenchmarkWithLibSlice_stdlib_10-4            336           3576072 ns/op
BenchmarkWithLibSlice_stdlib_10-4            336           3568089 ns/op
BenchmarkWithLibSlice_stdlib_10-4            333           3555157 ns/op
BenchmarkWithLibSlice_stdlib_10-4            333           3539808 ns/op
BenchmarkWithLibSlice_stdlib_10-4            336           3645658 ns/op
BenchmarkWithLibSlice_stdlib_10-4            333           3553782 ns/op
```
Using fanunmarshall 10 workers jsoniter lib
```
BenchmarkWithLibSlice_jsoniter_10-4          501           2469929 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          474           2497836 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          469           2650896 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          502           2641072 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          512           2712920 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          430           2592961 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          433           2679012 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          436           2694356 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          423           2571725 ns/op
BenchmarkWithLibSlice_jsoniter_10-4          480           2772881 ns/op

PASS
```


# Contributing 
You can help to deliver a better fanunmarshaller, check out how you can do things [CONTRIBUTING.md](CONTRIBUTING.md)

# License 
© [This is Development BV](https://www.thisisdevelopment.nl), 2022~time.Now()
Released under the [MIT License](https://github.com/thisisdevelopment/fanunmarshal/blob/master/LICENSE)
