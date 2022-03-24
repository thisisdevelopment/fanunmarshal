package fanunmarshal

// UnMarshalSlice unmarshal a slice of []byte slices, returning the a slice
func (f *fanMarshal) UnMarshalSlice(data [][]byte, expected interface{}) []interface{} {
	// make channel from slice
	var (
		pipe          = make(chan []byte)
		workers       = []<-chan interface{}{}
		amountWorkers = f.amountWorkers
	)

	if f.autoScaleDown && len(data) < int(amountWorkers) {
		amountWorkers = uint(len(data))
	}

	for i := 1; i <= int(amountWorkers); i++ {
		workers = append(workers, f.unmarshal(pipe, expected))
	}

	go func() {
		for _, d := range data {
			pipe <- d
		}
		close(pipe)
	}()

	var out = []interface{}{}
	for res := range f.fanIn(workers...) {
		out = append(out, res)
	}
	return out
}
func (f *fanMarshal) UnMarshalChan(pipe <-chan []byte, expected interface{}, dataLength *int) <-chan interface{} {
	// make channel from slice
	var (
		workers       = []<-chan interface{}{}
		amountWorkers = f.amountWorkers
		out           = make(chan interface{})
	)

	if dataLength == nil && f.autoScaleDown {
		panic("to use auto scaledown with a channel you have to pass in the expected dataLength or use DisableAutoScaleDown()")
	}

	if dataLength != nil && *dataLength > 0 {
		if f.autoScaleDown && *dataLength < int(amountWorkers) {
			amountWorkers = uint(*dataLength)
		}
	}

	for i := 1; i <= int(amountWorkers); i++ {
		workers = append(workers, f.unmarshal(pipe, expected))
	}

	go func() {
		for res := range f.fanIn(workers...) {
			out <- res
		}
		close(out)
	}()

	return out
}
