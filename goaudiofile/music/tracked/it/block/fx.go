package block

import "github.com/gotracker/goaudiofile/internal/util"

type FXPluginType [4]byte

var (
	PluginTypeVST = FXPluginType{'P', 's', 't', 'V'}
	PluginTypeDMO = FXPluginType{'O', 'M', 'X', 'D'}
)

type FXGainFactor uint8

// Value returns the value as a multiplier
func (g FXGainFactor) Value() float64 {
	return float64(g) / 10
}

type FXUserPluginName [32]byte

func (p FXUserPluginName) String() string {
	return util.GetString(p[:])
}

type FXLibraryName [64]byte

func (l FXLibraryName) String() string {
	return util.GetString(l[:])
}

// FX is a FX__ block
type FX struct {
	blockBase
	PluginType     FXPluginType     // Plugin type ("PtsV" for VST, "OMXD" for DMO plugins)
	UniqueID       [4]byte          // Plugin unique ID
	RoutingFlags   uint8            // Routing Flags
	MixMode        uint8            // Mix Mode
	GainFactor     FXGainFactor     // Gain Factor * 10 (9 = 90%, 10 = 100%, 11 = 110%, etc.)
	Reserved0B     uint8            // Reserved
	OutputRouting  uint32           // Output Routing (0 = send to master 0x80 + x = send to plugin x)
	Reserved10     [16]byte         // Reserved
	UserPluginName FXUserPluginName // User-chosen plugin name (Windows code page)
	LibraryName    FXLibraryName    // Library name (Original DLL name / DMO identifier - UTF-8 starting from OpenMPT 1.22.07.01, Windows code page in older versions)
	DataLength     uint32           // Length of plugin-specific data (parameters or opaque chunk)
	Data           []byte           // Plugin-specific data
}

// FourCC returns the big-endian representation of the block identifier
func (b *FX) FourCC() uint32 {
	return b.blockBase.FourCC()
}

// Length returns the size of the whole block
func (b *FX) Length() int {
	return b.blockBase.Length()
}
