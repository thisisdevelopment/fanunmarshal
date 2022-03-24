package fanunmarshal

import (
	jsonstd "encoding/json"
	"sync"

	jsoniter "github.com/json-iterator/go"

	"github.com/barkimedes/go-deepcopy"
)

var jsonntr = jsoniter.ConfigCompatibleWithStandardLibrary

func (f *fanMarshal) MakeChan(data [][]byte) <-chan []byte {
	var out = make(chan []byte)
	go func() {
		for _, d := range data {
			out <- d
		}
		close(out)
	}()

	return out

}

func (f *fanMarshal) fanIn(chs ...<-chan interface{}) <-chan interface{} {

	var (
		out = make(chan interface{})
		wg  = &sync.WaitGroup{}
	)
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(chx <-chan interface{}) {
			for exp := range chx {
				out <- exp
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func (f *fanMarshal) unmarshal(ch <-chan []byte, expected interface{}) <-chan interface{} {
	var out = make(chan interface{})
	go func() {
		for d := range ch {
			exp, _ := deepcopy.Anything(expected)
			fn := jsonstd.Unmarshal
			if !f.useStdLib {
				fn = jsonntr.Unmarshal
			}
			if err := fn(d, exp); err != nil {
				panic(err)
			}

			out <- exp
		}
		close(out)
	}()
	return out
}
