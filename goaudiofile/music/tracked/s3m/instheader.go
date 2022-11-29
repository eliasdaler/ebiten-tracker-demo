package s3m

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"

	"github.com/gotracker/goaudiofile/internal/util"
)

// HiLo32 is a 32 bit value where the first and second 16 bits are stored separately
type HiLo32 struct {
	Lo uint16
	Hi uint16
}

// SCRSType is the type of the SCRS instrument/sample
type SCRSType uint8

const (
	// SCRSTypeNone is a None type
	SCRSTypeNone = SCRSType(0 + iota)
	// SCRSTypeDigiplayer is a Digiplayer/S3M PCM sample
	SCRSTypeDigiplayer
	// SCRSTypeOPL2Melody is an Adlib/OPL2 melody instrument
	SCRSTypeOPL2Melody
	// SCRSTypeOPL2BassDrum is an Adlib/OPL2 bass drum instrument
	SCRSTypeOPL2BassDrum
	// SCRSTypeOPL2Snare is an Adlib/OPL2 snare drum instrument
	SCRSTypeOPL2Snare
	// SCRSTypeOPL2Tom is an Adlib/OPL2 tom drum instrument
	SCRSTypeOPL2Tom
	// SCRSTypeOPL2Cymbal is an Adlib/OPL2 cymbal instrument
	SCRSTypeOPL2Cymbal
	// SCRSTypeOPL2HiHat is an Adlib/OPL2 hi-hat instrument
	SCRSTypeOPL2HiHat
)

// SCRSHeader is the S3M instrument/sample header definition
type SCRSHeader struct {
	Type     SCRSType
	Filename [12]byte
}

// GetFilename returns a string representation of the data stored in the Filename field
func (h *SCRSHeader) GetFilename() string {
	return util.GetString(h.Filename[:])
}

// SCRSAncillaryHeader is the generic interface of the Type-specific header
type SCRSAncillaryHeader interface{}

// SCRSNoneHeader is the remaining header for S3M none-type instrument
type SCRSNoneHeader struct {
	Reserved0D [19]byte
	Volume     Volume
	Reserved1D [3]byte
	C2Spd      HiLo32
	Reserved24 [12]byte
	SampleName [28]byte
	Reserved4C [4]uint8
}

// GetSampleName returns a string representation of the data stored in the SampleName field
func (h *SCRSNoneHeader) GetSampleName() string {
	return util.GetString(h.SampleName[:])
}

// SCRS is a full header for an S3M instrument
type SCRS struct {
	Head      SCRSHeader
	Ancillary SCRSAncillaryHeader
}

// ReadSCRS reads an SCRS from the input stream
func ReadSCRS(r io.Reader) (*SCRS, error) {
	sh := SCRS{}
	if err := binary.Read(r, binary.LittleEndian, &sh.Head); err != nil {
		return nil, err
	}

	switch sh.Head.Type {
	case SCRSTypeNone:
		sh.Ancillary = &SCRSNoneHeader{}
	case SCRSTypeDigiplayer:
		sh.Ancillary = &SCRSDigiplayerHeader{}
	case SCRSTypeOPL2Melody, SCRSTypeOPL2BassDrum, SCRSTypeOPL2Snare, SCRSTypeOPL2Tom, SCRSTypeOPL2Cymbal, SCRSTypeOPL2HiHat:
		sh.Ancillary = &SCRSAdlibHeader{}
	default:
		return nil, errors.Errorf("unknown SCRS instrument type %0.2x", sh.Head.Type)
	}

	if sh.Ancillary != nil {
		if err := binary.Read(r, binary.LittleEndian, sh.Ancillary); err != nil {
			return nil, err
		}
	}

	return &sh, nil
}
