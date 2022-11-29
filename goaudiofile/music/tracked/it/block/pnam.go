package block

import "github.com/gotracker/goaudiofile/internal/util"

type PatternName [32]byte

func (n *PatternName) String() string {
	return util.GetString((*n)[:])
}

// PatternNames is a PNAM block
type PatternNames struct {
	blockBase
	Name []PatternName
}

// FourCC returns the big-endian representation of the block identifier
func (b *PatternNames) FourCC() uint32 {
	return b.blockBase.FourCC()
}

// Length returns the size of the whole block
func (b *PatternNames) Length() int {
	return b.blockBase.Length()
}
