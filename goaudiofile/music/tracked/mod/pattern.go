package mod

// Channel is a representation of the MOD file pattern channel bitfield
type Channel [4]uint8

// Period is a coefficient used in the calculation of an instrument's variable playback frequency (i.e.: a note)
type Period uint16

// Instrument returns the instrument number for this pattern channel
func (p Channel) Instrument() uint8 {
	return (p[0] & 0xF0) | (p[2] >> 4)
}

// Period returns the note period for this pattern channel
func (p Channel) Period() Period {
	return (Period(p[0]&0x0F) << 8) | Period(p[1])
}

// Effect returns the effect value for this pattern channel
func (p Channel) Effect() uint8 {
	return (p[2] & 0x0F)
}

// EffectParameter returns the effect parameter value for this pattern channel
func (p Channel) EffectParameter() uint8 {
	return p[3]
}

// Row is an array of all channels for a particular pattern row
type Row []Channel

// Pattern is a representation of a MOD file's single pattern
type Pattern [64]Row

// NewPattern creates a new pattern with a number of channels equal to the `channels` parameter
func NewPattern(channels int) Pattern {
	p := Pattern{}

	for r := range p {
		row := make(Row, channels)
		p[r] = row
	}

	return p
}
