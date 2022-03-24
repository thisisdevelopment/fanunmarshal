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

func New() IFanMarshal {
	return &fanMarshal{
		amountWorkers: DefaultWorkers,
		autoScaleDown: AutoScaleDown,
		useStdLib:     UseStdLib,
	}
}

func (f *fanMarshal) WithWorkers(workers uint) IFanMarshal {
	if workers == 0 {
		workers = DefaultWorkers
	}
	f.amountWorkers = workers
	return f
}

func (f *fanMarshal) DisableAutoScaleDown() IFanMarshal {
	f.autoScaleDown = false
	return f
}

func (f *fanMarshal) WithUseJsonIter() IFanMarshal {
	f.useStdLib = false
	return f
}
