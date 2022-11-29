package it

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/gotracker/goaudiofile/music/tracked/it/block"
)

func readBlock(data []byte, ptr ParaPointer, cmwt uint16) (block.Block, error) {
	ofs := ptr.Offset()
	if ofs > len(data)-4 {
		return nil, io.EOF
	}
	r := bytes.NewBuffer(data[ofs:])

	var blockID uint32
	if err := binary.Read(r, binary.BigEndian, &blockID); err != nil {
		return nil, err
	}

	switch {
	case blockID == 0x504E414D: // PNAM
		return readBlockPNAM(data, ptr, cmwt)
	case blockID>>16 == 0x4658: // FX__
		return readBlockFX00(data, ptr, cmwt)
	default:
		return readBlockUnknown(data, ptr, cmwt)
	}
}

func readBlockPNAM(data []byte, ptr ParaPointer, cmwt uint16) (block.Block, error) {
	p := block.PatternNames{}

	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	if err := binary.Read(r, binary.LittleEndian, &p.Identifier); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.BlockLen); err != nil {
		return nil, err
	}

	var nam block.PatternName
	cNameLen := len(nam)

	for pos := uint32(0); pos < p.BlockLen; {
		nlen := int(p.BlockLen - pos)
		if nlen > cNameLen {
			nlen = cNameLen
		}
		n := make([]byte, nlen)
		if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
			return nil, err
		}
		for len(n) < cNameLen {
			n = append(n, 0)
		}
		copy(nam[:], n)
		p.Name = append(p.Name, nam)
		pos += uint32(nlen)
	}

	return &p, nil
}

func readBlockFX00(data []byte, ptr ParaPointer, cmwt uint16) (block.Block, error) {
	p := block.FX{}

	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	if err := binary.Read(r, binary.LittleEndian, &p.Identifier); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.BlockLen); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.PluginType); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.UniqueID); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.RoutingFlags); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.MixMode); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.GainFactor); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Reserved0B); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.OutputRouting); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Reserved10); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.UserPluginName); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.LibraryName); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.DataLength); err != nil {
		return nil, err
	}

	p.Data = make([]byte, int(p.DataLength))
	if err := binary.Read(r, binary.LittleEndian, &p.Data); err != nil {
		return nil, err
	}

	return &p, nil
}

func readBlockUnknown(data []byte, ptr ParaPointer, cmwt uint16) (block.Block, error) {
	p := block.Unknown{}

	ofs := ptr.Offset()
	r := bytes.NewBuffer(data[ofs:])

	if err := binary.Read(r, binary.LittleEndian, &p.Identifier); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.BlockLen); err != nil {
		return nil, err
	}

	if p.Length()+ofs > len(data) {
		return nil, io.EOF
	}

	return &p, nil
}
