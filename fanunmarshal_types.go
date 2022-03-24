package fanunmarshal

type IFanUnMarshal interface {
	// WithWorkers set the amount of workers to work on your list
	WithWorkers(workers uint) IFanUnMarshal
	// DisableAutoScaleDown, disable scaling down the max amount of workers based on your list amount
	DisableAutoScaleDown() IFanUnMarshal
	// WithUseJsonIter use jsoniter lib instead of default std lib json package
	WithUseJsonIter() IFanUnMarshal
	// UnMarshalSlice unmarshal a slice, returning the slice
	UnMarshalSlice(data [][]byte, expected interface{}) []interface{}
	// MakeChan make a channel based of a slice of byteslices
	MakeChan(data [][]byte) <-chan []byte
	// UnMarshalChan unmarshal a slice, returning the slice
	UnMarshalChan(pipe <-chan []byte, expected interface{}, dataLength *int) <-chan interface{}
}
