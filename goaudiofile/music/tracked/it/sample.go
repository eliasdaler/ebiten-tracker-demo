package it

import (
	"bytes"
	"encoding/binary"

	"github.com/gotracker/goaudiofile/internal/util"
)

// Sample is a sample from the IT format
type Sample struct {
	IMPS             [4]byte
	Filename         [12]byte
	Reserved10       [1]byte
	GlobalVolume     Volume
	Flags            SampleFlags
	Volume           Volume
	Name             [26]byte
	ConvertFlags     ConvertFlags
	DefaultPan       SamplePanValue
	Length           uint32
	LoopBegin        uint32
	LoopEnd          uint32
	C5Speed          uint32
	SustainLoopBegin uint32
	SustainLoopEnd   uint32
	SamplePointer    ParaPointer32
	VibratoSpeed     uint8
	VibratoDepth     uint8
	VibratoSweep     uint8
	VibratoType      uint8
}

// GetName returns a string representation of the data stored in the Name field
func (s *Sample) GetName() string {
	return util.GetString(s.Name[:])
}

// GetFilename returns a string representation of the data stored in the Filename field
func (s *Sample) GetFilename() string {
	return util.GetString(s.Filename[:])
}

func readIMPS(data []byte, ptr ParaPointer, cmwt uint16) (*Sample, error) {
	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	sample := Sample{}
	if err := binary.Read(r, binary.LittleEndian, &sample); err != nil {
		return nil, err
	}

	return &sample, nil
}

func readSampleData(data []byte, ptr ParaPointer, cmwt uint16, out []byte) error {
	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	if _, err := r.Read(out); err != nil {
		return err
	}

	return nil
}
