package fanunmarshal

type IFanMarshal interface {
	// WithWorkers set the amount of workers to work on your list
	WithWorkers(workers uint) IFanMarshal
	// DisableAutoScaleDown, disable scaling down the max amount of workers based on your list amount
	DisableAutoScaleDown() IFanMarshal
	// WithUseJsonIter use jsoniter lib instead of default std lib json package
	WithUseJsonIter() IFanMarshal

	// UnMarshalSlice unmarshal a slice, returning the slice
	UnMarshalSlice(data [][]byte, expected interface{}) []interface{}

	// MakeChan make a channel based of a slice of byteslices
	MakeChan(data [][]byte) <-chan []byte
	// UnMarshalChan unmarshal a slice, returning the slice
	UnMarshalChan(pipe <-chan []byte, expected interface{}, dataLength *int) <-chan interface{}
}
