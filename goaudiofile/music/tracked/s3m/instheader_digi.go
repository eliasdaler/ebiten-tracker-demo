package s3m

import "github.com/gotracker/goaudiofile/internal/util"

// SCRSFlags is a bitset for the S3M instrument/sample header definition
type SCRSFlags uint8

const (
	// SCRSFlagsLooped is looping
	SCRSFlagsLooped = SCRSFlags(0x01)
	// SCRSFlagsStereo is stereo
	SCRSFlagsStereo = SCRSFlags(0x02)
	// SCRSFlags16Bit is 16-bit
	SCRSFlags16Bit = SCRSFlags(0x04)
)

// IsLooped returns true if bit 0 is set
func (f SCRSFlags) IsLooped() bool {
	return (f & SCRSFlagsLooped) != 0
}

// IsStereo returns true if bit 1 is set
func (f SCRSFlags) IsStereo() bool {
	return (f & SCRSFlagsStereo) != 0
}

// Is16BitSample returns true if bit 2 is set
func (f SCRSFlags) Is16BitSample() bool {
	return (f & SCRSFlags16Bit) != 0
}

// Packing is a type of sample packing format
type Packing uint8

const (
	// PackingUnpacked is an unpacked S3M PCM sample
	PackingUnpacked = Packing(iota)
	// PackingDP30ADPCM is Digiplayer/ST3 3.00 ADPCM packing
	PackingDP30ADPCM
)

// SCRSDigiplayerHeader is the remaining header for S3M PCM samples
type SCRSDigiplayerHeader struct {
	MemSeg        ParaPointer24
	Length        HiLo32
	LoopBegin     HiLo32
	LoopEnd       HiLo32
	Volume        Volume
	Reserved1D    uint8
	PackingScheme Packing
	Flags         SCRSFlags
	C2Spd         HiLo32
	Reserved24    [4]byte
	Reserved28    [2]byte
	Reserved2A    [2]byte
	Reserved2C    [4]byte
	SampleName    [28]byte
	SCRS          [4]uint8
}

// GetSampleName returns a string representation of the data stored in the SampleName field
func (h *SCRSDigiplayerHeader) GetSampleName() string {
	return util.GetString(h.SampleName[:])
}
