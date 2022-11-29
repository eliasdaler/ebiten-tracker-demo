// startrekker / star tracker

package mod

import (
	"encoding/binary"
	"errors"
	"io"
)

type fmtST struct {
	formatIntf
}

var (
	startrekker = &fmtST{}
)

func (f *fmtST) readPattern(ffmt *modFormatDetails, r io.Reader) (*Pattern, error) {
	if r == nil {
		return nil, errors.New("r is nil")
	}

	p := NewPattern(ffmt.channels)
	for _, row := range p {
		for c := 0; c < 4; c++ {
			if err := binary.Read(r, binary.LittleEndian, &row[c]); err != nil {
				return nil, err
			}
		}
	}
	// weird format...
	if ffmt.channels == 8 {
		for _, row := range p {
			for c := 4; c < 8; c++ {
				if err := binary.Read(r, binary.LittleEndian, &row[c]); err != nil {
					return nil, err
				}
			}
		}
	}

	return &p, nil
}

func (f *fmtST) rectifyOrderList(ffmt *modFormatDetails, in [128]uint8) ([128]uint8, error) {
	// really weird format...
	if ffmt.channels == 8 {
		out := [128]uint8{}
		for i, o := range in {
			out[i] = o / 2
		}
		return out, nil
	}
	return in, nil
}

func init() {
	// fasttracker
	signatureLookup["FLT4"] = modFormatDetails{4, startrekker}
	signatureLookup["FLT8"] = modFormatDetails{8, startrekker}
}
