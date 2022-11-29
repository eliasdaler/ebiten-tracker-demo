package mod

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

// File is an MOD internal file representation
type File struct {
	Head     ModuleHeader
	Patterns []Pattern
	Samples  []SampleData
}

type formatIntf interface {
	readPattern(*modFormatDetails, io.Reader) (*Pattern, error)
	rectifyOrderList(*modFormatDetails, [128]uint8) ([128]uint8, error)
}

type modFormatDetails struct {
	channels int
	format   formatIntf
}

var (
	signatureLookup = make(map[string]modFormatDetails)
)

// Read reads a MOD file from the reader `r` and creates an internal MOD File representation
func Read(r io.Reader) (*File, error) {
	f := File{}

	if err := binary.Read(r, binary.LittleEndian, &f.Head); err != nil {
		return nil, err
	}

	sig := util.GetString(f.Head.Sig[:])
	var ffmt *modFormatDetails
	s, ok := signatureLookup[sig]
	if ok {
		ffmt = &s
	}

	if ffmt == nil || ffmt.channels == 0 {
		return nil, errors.New("invalid file format")
	}

	processor := ffmt.format
	if processor == nil {
		return nil, errors.New("could not identify format reader")
	}

	numPatterns := 0
	orderList, err := processor.rectifyOrderList(ffmt, f.Head.Order)
	if err != nil {
		return nil, err
	}
	for i, o := range orderList {
		if i < int(f.Head.SongLen) {
			f.Head.Order[i] = o
		}
		// we count all patterns, even if we're not in the 'song' range
		// hidden/'deleted' patterns can exist...
		if numPatterns <= int(o) {
			numPatterns = int(o) + 1
		}
	}

	f.Patterns = make([]Pattern, numPatterns)
	for i := 0; i < numPatterns; i++ {
		pattern, err := processor.readPattern(ffmt, r)
		if err != nil {
			return nil, err
		}
		if pattern == nil {
			continue
		}
		f.Patterns[i] = *pattern
	}

	f.Samples = make([]SampleData, len(f.Head.Instrument))
	for instNum, inst := range f.Head.Instrument {
		samp := make([]byte, inst.Len.Value())
		if err := binary.Read(r, binary.LittleEndian, &samp); err != nil {
			return nil, err
		}
		f.Samples[instNum] = samp
	}

	return &f, nil
}
