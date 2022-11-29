package mod

import "github.com/gotracker/goaudiofile/internal/util"

// ModuleHeader is a representation of the MOD file header
type ModuleHeader struct {
	Name       [20]byte
	Instrument [31]InstrumentHeader
	SongLen    uint8
	RestartPos uint8
	Order      [128]uint8
	Sig        [4]uint8
}

// GetName returns a string representation of the data stored in the Name field
func (mh *ModuleHeader) GetName() string {
	return util.GetString(mh.Name[:])
}
