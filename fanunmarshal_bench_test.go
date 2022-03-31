package fanunmarshal_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/thisisdevelopment/fanunmarshal"
)

type testData []testObj

type testObj struct {
	ID            string   `json:"_id"`
	Index         int64    `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int64    `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

type friend struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func generateJsonL() (byteData [][]byte) {
	d := testData{}
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
			var expected = testObj{}
			json.Unmarshal(d, &expected)
		}
	}
}

func BenchmarkWithLibSlice_stdlib_10(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := fanunmarshal.New().WithWorkers(10)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = testObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_stdlib_100(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := fanunmarshal.New().WithWorkers(100)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = testObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_jsoniter_10(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := fanunmarshal.New().WithWorkers(10).WithUseJsonIter()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = testObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}

func BenchmarkWithLibSlice_jsoniter_100(b *testing.B) {

	b.StopTimer()
	data := generateJsonL()

	fm := fanunmarshal.New().WithWorkers(100).WithUseJsonIter()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var expected = testObj{}
		data := fm.UnMarshalSlice(data, &expected)
		_ = data
	}
}
