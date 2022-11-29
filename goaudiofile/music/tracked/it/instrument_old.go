package it

import "github.com/gotracker/goaudiofile/internal/util"

// IMPIInstrumentOld is the format of the IMPI Instrument for tracker compatibility versions < 0x0200
type IMPIInstrumentOld struct {
	IMPI               [4]byte
	Filename           [12]byte
	Nul10              uint8
	Flags              IMPIOldFlags
	VolumeLoopStart    uint8
	VolumeLoopEnd      uint8
	SustainLoopStart   uint8
	SustainLoopEnd     uint8
	Fadeout            uint16
	NewNoteAction      NewNoteAction
	DuplicateNoteCheck DuplicateNoteCheck
	TrackerVersion     uint16
	SampleCount        uint8
	Reserved1F         uint8
	Name               [26]byte
	Reserved3A         [6]uint8
	NoteSampleKeyboard [120]NoteSample
	VolumeEnvelope     [200]uint8
	NodePoints         [25]NodePoint16
}

// GetName returns a string representation of the data stored in the Name field
func (i *IMPIInstrumentOld) GetName() string {
	return util.GetString(i.Name[:])
}

// GetFilename returns a string representation of the data stored in the Filename field
func (i *IMPIInstrumentOld) GetFilename() string {
	return util.GetString(i.Filename[:])
}
