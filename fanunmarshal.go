// Package fanunmarshal is a concurrent unmarshaller
//
// use with slices of byte slices [][]byte, for example for data coming from Redis using MGet
// if we need to unmarshal huge structs, this is where the package shines
package fanunmarshal

const (
	// default amount of workers
	DefaultWorkers = 2
	// default scaledown the amount of workers set based on the list size
	AutoScaleDown = true
	// default use the standard json encoding lib
	UseStdLib = true
	// use jsoniter (faster)
	UseJsoniter = false
)

type FanUnMarshal struct {
	amountWorkers uint
	autoScaleDown bool
	useStdLib     bool
}

//New instance
func New() IFanUnMarshal {
	return &FanUnMarshal{
		amountWorkers: DefaultWorkers,
		autoScaleDown: AutoScaleDown,
		useStdLib:     UseStdLib,
	}
}

// WithWorkers set the amount of workers to work on your list
func (f *FanUnMarshal) WithWorkers(workers uint) IFanUnMarshal {
	if workers == 0 {
		workers = DefaultWorkers
	}
	f.amountWorkers = workers
	return f
}

// DisableAutoScaleDown, disable scaling down the max amount of workers based on your list amount
func (f *FanUnMarshal) DisableAutoScaleDown() IFanUnMarshal {
	f.autoScaleDown = false
	return f
}

// WithUseJsonIter use jsoniter lib instead of default std lib json package
func (f *FanUnMarshal) WithUseJsonIter() IFanUnMarshal {
	f.useStdLib = false
	return f
}
