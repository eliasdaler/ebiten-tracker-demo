package xm

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

// ModuleHeader is a representation of the XM file header
type ModuleHeader struct {
	IDText          [17]uint8
	Name            [20]uint8
	Reserved1A      uint8
	TrackerName     [20]uint8
	VersionNumber   uint16
	HeaderSize      uint32
	SongLength      uint16
	RestartPosition uint16
	NumChannels     uint16
	NumPatterns     uint16
	NumInstruments  uint16
	Flags           HeaderFlags
	DefaultSpeed    uint16
	DefaultTempo    uint16
	OrderTable      [256]uint8
}

// GetIDText returns a string representation of the data stored in the IDText field
func (mh *ModuleHeader) GetIDText() string {
	return util.GetString(mh.IDText[:])
}

// GetName returns a string representation of the data stored in the Name field
func (mh *ModuleHeader) GetName() string {
	return util.GetString(mh.Name[:])
}

// GetTrackerName returns a string representation of the data stored in the TrackerName field
func (mh *ModuleHeader) GetTrackerName() string {
	return util.GetString(mh.Name[:])
}

// HeaderFlags is the set of flags for an XM header
type HeaderFlags uint16

const (
	// HeaderFlagLinearSlides activates the linear frequency table (off = Amiga frequency table)
	HeaderFlagLinearSlides = HeaderFlags(0x0001)
	// HeaderFlagExtendedFilterRange activates the extended filter range
	HeaderFlagExtendedFilterRange = HeaderFlags(0x1000)
)

// IsLinearSlides returns true if the song plays with linear note slides (or if false, with Amiga note slides)
func (f HeaderFlags) IsLinearSlides() bool {
	return (f & HeaderFlagLinearSlides) != 0
}

// IsExtendedFilterRange returns true if the song has extended filter ranges enabled
func (f HeaderFlags) IsExtendedFilterRange() bool {
	return (f & HeaderFlagExtendedFilterRange) != 0
}

func readHeaderPartial(r io.Reader) (*ModuleHeader, error) {
	xmh := ModuleHeader{}

	if err := binary.Read(r, binary.LittleEndian, &xmh.IDText); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.Name); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.Reserved1A); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.TrackerName); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.VersionNumber); err != nil {
		return nil, err
	}

	sz := uint32(0)
	if err := binary.Read(r, binary.LittleEndian, &xmh.HeaderSize); err != nil {
		return nil, err
	}
	sz += 4

	if err := binary.Read(r, binary.LittleEndian, &xmh.SongLength); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.RestartPosition); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.NumChannels); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.NumPatterns); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.NumInstruments); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.Flags); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.DefaultSpeed); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &xmh.DefaultTempo); err != nil {
		return nil, err
	}
	if sz += 2; sz >= xmh.HeaderSize {
		return &xmh, nil
	}

	for i := range xmh.OrderTable {
		if err := binary.Read(r, binary.LittleEndian, &xmh.OrderTable[i]); err != nil {
			return nil, err
		}
		if sz++; sz >= xmh.HeaderSize {
			return &xmh, nil
		}
	}

	return &xmh, nil
}

func readHeader(r io.Reader) (*ModuleHeader, error) {
	xmh, err := readHeaderPartial(r)
	if err != nil {
		return nil, err
	}

	if xmh.NumChannels < 1 || xmh.NumChannels > 32 {
		return nil, errors.New("invalid number of channels - possibly corrupt file")
	}

	if xmh.NumPatterns > 256 {
		return nil, errors.New("invalid number of patterns - possibly corrupt file")
	}

	if xmh.NumInstruments > 128 {
		return nil, errors.New("invalid number of instruments - possibly corrupt file")
	}

	return xmh, nil
}
