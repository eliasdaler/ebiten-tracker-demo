package mod

import "github.com/gotracker/goaudiofile/internal/util"

// InstrumentHeader is a representation of the MOD file instrument header
type InstrumentHeader struct {
	Name      [22]byte
	Len       WordLength
	FineTune  uint8
	Volume    uint8
	LoopStart WordLength
	LoopEnd   WordLength
}

// GetName returns a string representation of the data stored in the Name field
func (i *InstrumentHeader) GetName() string {
	return util.GetString(i.Name[:])
}

// SampleData is the data associated to the instrument
type SampleData []uint8
