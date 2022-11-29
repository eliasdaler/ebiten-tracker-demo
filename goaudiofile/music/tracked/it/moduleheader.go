package it

import (
	"encoding/binary"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

// ModuleHeader is the initial header definition of an IT file
type ModuleHeader struct {
	IMPM                 [4]byte
	Name                 [26]byte
	PHighlight           uint16
	OrderCount           uint16
	InstrumentCount      uint16
	SampleCount          uint16
	PatternCount         uint16
	TrackerVersion       uint16
	TrackerCompatVersion uint16
	Flags                IMPMFlags
	SpecialFlags         IMPMSpecialFlags
	GlobalVolume         FineVolume
	MixingVolume         FineVolume
	InitialSpeed         uint8
	InitialTempo         uint8
	PanningSeparation    PanSeparation
	PitchWheelDepth      uint8
	MessageLength        uint16
	MessageOffset        ParaPointer32
	Reserved3C           [4]uint8
	ChannelPan           [64]PanValue
	ChannelVol           [64]Volume
}

// GetName returns a string representation of the data stored in the Name field
func (mh *ModuleHeader) GetName() string {
	return util.GetString(mh.Name[:])
}

// ReadModuleHeader reads a ModuleHeader from the input stream
func ReadModuleHeader(r io.Reader) (*ModuleHeader, error) {
	var mh ModuleHeader
	if err := binary.Read(r, binary.LittleEndian, &mh); err != nil {
		return nil, err
	}

	return &mh, nil
}
