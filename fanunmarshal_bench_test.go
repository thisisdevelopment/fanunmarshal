package fanunmarshal

import (
	"encoding/json"
	"os"
	"testing"
)

func generateJsonL() (byteData [][]byte) {
	d := TestData{}
	data, _ := os.ReadFile("./testdata/generated.json")
	json.Unmarshal(data, &d)

	for _, obj := range d {
		objb, _ := json.Marshal(obj)
		byteData = append(byteData, objb)
	}
	return

}

func BenchmarkPlainUnMarshal(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, d := range data {
			var expected = TestObj{}
			json.Unmarshal(d, &expected)
		}
	}
}

func BenchmarkWithLibSlice_stdlib_10(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := New().WithWorkers(10)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = TestObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_stdlib_100(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := New().WithWorkers(100)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = TestObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_jsoniter_10(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := New().WithWorkers(10).WithUseJsonIter()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = TestObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_jsoniter_100(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := New().WithWorkers(100).WithUseJsonIter()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = TestObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}
