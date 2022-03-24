// Package fanunmarshal is a concurrent unmarshaller
//
// use with slices of byte slices [][]byte, for example for data coming from Redis using MGet
package fanunmarshal

const (
	// default amount of workers
	DefaultWorkers = 2
	AutoScaleDown  = true
	UseStdLib      = true
	UseJsoniter    = false
)

type fanMarshal struct {
	amountWorkers uint
	autoScaleDown bool
	useStdLib     bool
}

// New instance
func New() IFanMarshal {
	return &fanMarshal{
		amountWorkers: DefaultWorkers,
		autoScaleDown: AutoScaleDown,
		useStdLib:     UseStdLib,
	}
}

// WithWorkers set the amount of workers to work on your list
func (f *fanMarshal) WithWorkers(workers uint) IFanMarshal {
	if workers == 0 {
		workers = DefaultWorkers
	}
	f.amountWorkers = workers
	return f
}

// DisableAutoScaleDown, disable scaling down the max amount of workers based on your list amount
func (f *fanMarshal) DisableAutoScaleDown() IFanMarshal {
	f.autoScaleDown = false
	return f
}

// WithUseJsonIter use jsoniter lib instead of default std lib json package
func (f *fanMarshal) WithUseJsonIter() IFanMarshal {
	f.useStdLib = false
	return f
}
