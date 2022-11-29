package xm

import (
	"encoding/binary"
	"errors"
	"io"
)

// PatternHeader is the XM packed pattern header definition
type PatternHeader struct {
	PatternHeaderLength   uint32
	PackingType           uint8
	NumRows               uint16
	PackedPatternDataSize uint16
}

// ChannelData is the XM unpacked pattern channel data definition
type ChannelData struct {
	Flags           ChannelFlags
	Note            uint8
	Instrument      uint8
	Volume          uint8
	Effect          uint8
	EffectParameter uint8
}

// HasNote returns true when the channel includes note data
func (f ChannelData) HasNote() bool {
	return f.Flags.HasNote()
}

// HasInstrument returns true when the channel includes instrument data
func (f ChannelData) HasInstrument() bool {
	return f.Flags.HasInstrument()
}

// HasVolume returns true when the channel includes volume data
func (f ChannelData) HasVolume() bool {
	return f.Flags.HasVolume()
}

// HasEffect returns true when the channel includes effect data
func (f ChannelData) HasEffect() bool {
	return f.Flags.HasEffect()
}

// HasEffectParameter returns true when the channel includes effect parameter data
func (f ChannelData) HasEffectParameter() bool {
	return f.Flags.HasEffectParameter()
}

// IsValid returns true when the channel flags are valid
func (f ChannelData) IsValid() bool {
	return f.Flags.IsValid()
}

// ChannelFlags describes what is valid in a channel
type ChannelFlags uint8

const (
	// ChannelFlagHasNote signifies that the channel data includes a note
	ChannelFlagHasNote = ChannelFlags(0x01)
	// ChannelFlagHasInstrument signifies that the channel data includes an instrument
	ChannelFlagHasInstrument = ChannelFlags(0x02)
	// ChannelFlagHasVolume signifies that the channel data includes a volume
	ChannelFlagHasVolume = ChannelFlags(0x04)
	// ChannelFlagHasEffect signifies that the channel data includes an effect
	ChannelFlagHasEffect = ChannelFlags(0x08)
	// ChannelFlagHasEffectParameter signifies that the channel data includes an effect parameter
	ChannelFlagHasEffectParameter = ChannelFlags(0x10)
	// ChannelFlagValid signifies that the channel flags are valid
	ChannelFlagValid = ChannelFlags(0x80)

	// ChannelFlagsAll is all channel flags at once
	ChannelFlagsAll = ChannelFlags(0xFF)
)

// HasNote returns true when the channel includes note data
func (f ChannelFlags) HasNote() bool {
	return (f & ChannelFlagHasNote) != 0
}

// HasInstrument returns true when the channel includes instrument data
func (f ChannelFlags) HasInstrument() bool {
	return (f & ChannelFlagHasInstrument) != 0
}

// HasVolume returns true when the channel includes volume data
func (f ChannelFlags) HasVolume() bool {
	return (f & ChannelFlagHasVolume) != 0
}

// HasEffect returns true when the channel includes effect data
func (f ChannelFlags) HasEffect() bool {
	return (f & ChannelFlagHasEffect) != 0
}

// HasEffectParameter returns true when the channel includes effect parameter data
func (f ChannelFlags) HasEffectParameter() bool {
	return (f & ChannelFlagHasEffectParameter) != 0
}

// IsValid returns true when the channel flags are valid
func (f ChannelFlags) IsValid() bool {
	return (f & ChannelFlagValid) != 0
}

// PatternRow is the XM unpacked pattern channel data list for a single pattern row
type PatternRow []ChannelData

// PatternFileFormat is the XM pattern definition in file format
type PatternFileFormat struct {
	Header     PatternHeader
	PackedData []byte
}

func readPatternHeaderPartial(r io.Reader, fileVersion uint16) (*PatternHeader, error) {
	ph := PatternHeader{}

	sz := uint32(0)
	if err := binary.Read(r, binary.LittleEndian, &ph.PatternHeaderLength); err != nil {
		return nil, err
	}
	if sz += 4; sz >= ph.PatternHeaderLength {
		return &ph, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &ph.PackingType); err != nil {
		return nil, err
	}
	if sz++; sz >= ph.PatternHeaderLength {
		return &ph, nil
	}

	if fileVersion == 0x0102 {
		var rowCount uint8
		if err := binary.Read(r, binary.LittleEndian, &rowCount); err != nil {
			return nil, err
		}

		ph.NumRows = uint16(rowCount) + 1
		if sz++; sz >= ph.PatternHeaderLength {
			return &ph, nil
		}

	} else {
		if err := binary.Read(r, binary.LittleEndian, &ph.NumRows); err != nil {
			return nil, err
		}
		if sz += 2; sz >= ph.PatternHeaderLength {
			return &ph, nil
		}
	}

	if err := binary.Read(r, binary.LittleEndian, &ph.PackedPatternDataSize); err != nil {
		return nil, err
	}
	if sz += 2; sz >= ph.PatternHeaderLength {
		return &ph, nil
	}

	return &ph, nil
}

func readPatternHeader(r io.Reader, fileVersion uint16) (*PatternHeader, error) {
	ph, err := readPatternHeaderPartial(r, fileVersion)
	if err != nil {
		return nil, err
	}

	//if ph.NumRows == 0 {
	//	ph.NumRows = 64
	//}

	if ph.PackingType != 0 {
		return nil, errors.New("unexpected pattern packing type - possibly corrupt file")
	}

	if ph.NumRows < 1 || ph.NumRows > 256 {
		return nil, errors.New("pattern row count out of range - possibly corrupt file")
	}

	return ph, nil
}
