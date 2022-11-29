package xm

import (
	"bytes"
	"encoding/binary"
	"io"
)

// File is an XM internal file representation
type File struct {
	Head        ModuleHeader
	Patterns    []Pattern
	Instruments []InstrumentHeader
}

// Read reads an XM file from the reader `r` and creates an internal File representation
func Read(r io.Reader) (*File, error) {
	xmh, err := readHeader(r)
	if err != nil {
		return nil, err
	}

	f := File{
		Head: *xmh,
	}

	for i := uint16(0); i < xmh.NumPatterns; i++ {
		p := Pattern{}

		ph, err := readPatternHeader(r, xmh.VersionNumber)
		if err != nil {
			return nil, err
		}

		p.Header = *ph

		ppd := make([]byte, int(ph.PackedPatternDataSize))
		if err := binary.Read(r, binary.LittleEndian, &ppd); err != nil {
			return nil, err
		}

		p.PackedData = ppd

		if err := p.unpack(int(xmh.NumChannels)); err != nil {
			return nil, err
		}

		f.Patterns = append(f.Patterns, p)
	}

	for i := uint16(0); i < xmh.NumInstruments; i++ {
		ih, err := readInstrumentHeader(r)
		if err != nil {
			return nil, err
		}

		f.Instruments = append(f.Instruments, *ih)
	}

	return &f, err
}

// Pattern is an XM internal file representation and converted/unpacked pattern set
type Pattern struct {
	PatternFileFormat

	Data []PatternRow
}

func (p *Pattern) unpack(numChannels int) error {
	numRows := int(p.Header.NumRows)

	if len(p.PackedData) == 0 {
		// empty pattern
		p.Data = make([]PatternRow, numRows)
		for i := range p.Data {
			p.Data[i] = make(PatternRow, numChannels)
		}
		return nil
	}

	// it's not empty, so let's unpack it!

	p.Data = make([]PatternRow, numRows)
	packed := bytes.NewReader(p.PackedData)
	for i := range p.Data {
		row := make(PatternRow, numChannels)
		p.Data[i] = row
		for c := 0; c < numChannels; c++ {
			ch := &row[c]
			if err := binary.Read(packed, binary.LittleEndian, &ch.Flags); err != nil {
				return err
			}

			// is the first byte a bitfield instead of note?
			if ch.IsValid() {
				// it is!
				// note present?
				if ch.HasNote() {
					if err := binary.Read(packed, binary.LittleEndian, &ch.Note); err != nil {
						return err
					}
				}
			} else {
				// it isn't... assume it's a note and that we have everything present
				ch.Note = uint8(ch.Flags)
				ch.Flags = ChannelFlagsAll
			}

			// instrument present?
			if ch.HasInstrument() {
				if err := binary.Read(packed, binary.LittleEndian, &ch.Instrument); err != nil {
					return err
				}
			}

			// volume present?
			if ch.HasVolume() {
				if err := binary.Read(packed, binary.LittleEndian, &ch.Volume); err != nil {
					return err
				}
			}

			// effect present?
			if ch.HasEffect() {
				if err := binary.Read(packed, binary.LittleEndian, &ch.Effect); err != nil {
					return err
				}
			}

			// effect parameter present?
			if ch.HasEffectParameter() {
				if err := binary.Read(packed, binary.LittleEndian, &ch.EffectParameter); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
