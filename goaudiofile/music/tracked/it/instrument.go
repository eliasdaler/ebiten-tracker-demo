package it

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

var (
	// ErrInvalidInstrumentFormat is an error for when an invalid instrument format is encountered
	ErrInvalidInstrumentFormat = errors.New("invalid instrument format")
)

// IMPIIntf is an interface to the IT instruments
type IMPIIntf interface{}

func readIMPI(data []byte, ptr ParaPointer, cmwt uint16) (IMPIIntf, error) {
	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	switch {
	case cmwt < 0x200:
		return readIMPIOld(r)
	default:
		return readIMPINew(r)
	}
}

func readIMPIOld(r io.Reader) (*IMPIInstrumentOld, error) {
	inst := IMPIInstrumentOld{}

	if err := binary.Read(r, binary.LittleEndian, &inst); err != nil {
		return nil, err
	}

	if util.GetString(inst.IMPI[:]) != "IMPI" {
		return nil, ErrInvalidInstrumentFormat
	}

	return &inst, nil
}

func readIMPINew(r io.Reader) (*IMPIInstrument, error) {
	inst := IMPIInstrument{}

	if err := binary.Read(r, binary.LittleEndian, &inst); err != nil {
		return nil, err
	}

	if util.GetString(inst.IMPI[:]) != "IMPI" {
		return nil, ErrInvalidInstrumentFormat
	}

	return &inst, nil
}
