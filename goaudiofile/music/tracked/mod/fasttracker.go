// fasttracker / fasttracker 2

package mod

import (
	"encoding/binary"
	"errors"
	"io"
)

type fmtFT struct {
	formatIntf
}

var (
	fasttracker = &fmtFT{}
)

func (f *fmtFT) readPattern(ffmt *modFormatDetails, r io.Reader) (*Pattern, error) {
	if r == nil {
		return nil, errors.New("r is nil")
	}

	p := NewPattern(ffmt.channels)
	for _, row := range p {
		for c := 0; c < ffmt.channels; c++ {
			if err := binary.Read(r, binary.LittleEndian, &row[c]); err != nil {
				return nil, err
			}
		}
	}

	return &p, nil
}

func (f *fmtFT) rectifyOrderList(ffmt *modFormatDetails, in [128]uint8) ([128]uint8, error) {
	return in, nil
}

func init() {
	// fasttracker
	signatureLookup["2CHN"] = modFormatDetails{2, fasttracker}
	signatureLookup["4CHN"] = modFormatDetails{4, fasttracker}
	signatureLookup["6CHN"] = modFormatDetails{6, fasttracker}
	signatureLookup["8CHN"] = modFormatDetails{8, fasttracker}

	// fasttracker 2
	signatureLookup["10CH"] = modFormatDetails{10, fasttracker}
	signatureLookup["11CH"] = modFormatDetails{11, fasttracker}
	signatureLookup["12CH"] = modFormatDetails{12, fasttracker}
	signatureLookup["13CH"] = modFormatDetails{13, fasttracker}
	signatureLookup["14CH"] = modFormatDetails{14, fasttracker}
	signatureLookup["15CH"] = modFormatDetails{15, fasttracker}
	signatureLookup["16CH"] = modFormatDetails{16, fasttracker}
	signatureLookup["17CH"] = modFormatDetails{17, fasttracker}
	signatureLookup["18CH"] = modFormatDetails{18, fasttracker}
	signatureLookup["19CH"] = modFormatDetails{19, fasttracker}
	signatureLookup["20CH"] = modFormatDetails{20, fasttracker}
	signatureLookup["21CH"] = modFormatDetails{21, fasttracker}
	signatureLookup["22CH"] = modFormatDetails{22, fasttracker}
	signatureLookup["23CH"] = modFormatDetails{23, fasttracker}
	signatureLookup["24CH"] = modFormatDetails{24, fasttracker}
	signatureLookup["25CH"] = modFormatDetails{25, fasttracker}
	signatureLookup["26CH"] = modFormatDetails{26, fasttracker}
	signatureLookup["27CH"] = modFormatDetails{27, fasttracker}
	signatureLookup["28CH"] = modFormatDetails{28, fasttracker}
	signatureLookup["29CH"] = modFormatDetails{29, fasttracker}
	signatureLookup["30CH"] = modFormatDetails{30, fasttracker}
	signatureLookup["31CH"] = modFormatDetails{31, fasttracker}
	signatureLookup["32CH"] = modFormatDetails{32, fasttracker}
}
