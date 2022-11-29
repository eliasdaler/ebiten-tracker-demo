package s3m

import (
	"math"

	"github.com/gotracker/goaudiofile/internal/util"
)

// OPL2Multiple is the YM3812/YMF262 Frequency Data Multiplier / MULT (MULTIPLE) value
type OPL2Multiple uint8

// Multiplier returns the actual frequency data multiplier
func (m OPL2Multiple) Multiplier() float64 {
	switch {
	case m == 0x00:
		return 0.5
	case m >= 0x01 && m <= 0x09:
		return float64(m)
	case m >= 0x0A && m <= 0x0B:
		return 10
	case m >= 0x0C && m <= 0x0D:
		return 12
	case m >= 0x0E && m <= 0x0F:
		return 15
	default:
		panic("unhandled value")
	}
}

// OPL2KSL is the YM3812/YMF262 Key Scale Level / OPL2KSL selection value
type OPL2KSL uint8

// Attenuation returns the attenuation per octave, in dB
func (k OPL2KSL) Attenuation() float64 {
	switch k {
	case 0x00:
		return 0
	case 0x01:
		return 1.5
	case 0x02:
		return 3
	case 0x03:
		return 6
	default:
		panic("unhandled value")
	}
}

// OPL2Feedback is the YM3812/YMF262 Feedback / FB modulation depth value
type OPL2Feedback uint8

// Modulation returns the feedback modulation depth
func (f OPL2Feedback) Modulation() float64 {
	switch {
	case f == 0x00:
		return 0
	case f >= 0x01 && f <= 0x07:
		div := 1 << (7 - int(f))
		return (4 * math.Pi) / float64(div)
	default:
		panic("unhandled value")
	}
}

// OPL2Waveform is a selection of the OPL2 waveform type
type OPL2Waveform uint8

const (
	// OPL2WaveformSine is a sine wave
	// Available in all OPL chips.
	OPL2WaveformSine = OPL2Waveform(0 + iota)
	// OPL2WaveformHalfSine is a sine wave that ends before going negative in aplitude
	// Also known as rectified sine wave
	// The negative parts of the wave are flattened to 0.
	// Available in OPLL, OPL2, and OPL3.
	OPL2WaveformHalfSine
	// OPL2WaveformAbsoluteSine is the absolute value of a sine wave
	// The negative parts of the wave are inverted to be positive.
	// Available in OPL2 and OPL3.
	OPL2WaveformAbsoluteSine
	// OPL2WaveformPulseSine is the absolute value of a sine wave that ends at the peak
	// The absolute sine wave, with the downward half of each envelope flattened to 0.
	// Available in OPL2 and OPL3.
	OPL2WaveformPulseSine
	// OPL2WaveformSineEvenPeriods is a sine wave on only even periods (alternating sine)
	// The wavelength is halved compared to the normal sine wave, and every alternating period is flattened to 0.
	// Available in OPL3.
	OPL2WaveformSineEvenPeriods
	// OPL2WaveformAbsoluteSineEvenPeriods is the absolute value of a sine wave and on only even periods
	// An absolute version of the alternating sine, resulting in pairs of humps grouped together like on the back of a camel.
	// Available in OPL3.
	OPL2WaveformAbsoluteSineEvenPeriods
	// OPL2WaveformSquare is a square wave
	// Available in OPL3.
	OPL2WaveformSquare
	// OPL2WaveformDerivedSquare is the logarithmic sawtooth wave
	// Logarithmic sawtooth wave, or derived square wave. This is actually the result of an exponentiation by a negative number.
	// Available in OPL3.
	OPL2WaveformDerivedSquare
)

// OPL2Specs is the specifiers for an OPL2/Adlib instrument
type OPL2Specs struct {
	Modulat0   uint8   // D00
	Carrier0   uint8   // D01
	Modulat1   uint8   // D02
	Carrier1   uint8   // D03
	Modulat2   uint8   // D04
	Carrier2   uint8   // D05
	Modulat3   uint8   // D06
	Carrier3   uint8   // D07
	Modulat4   uint8   // D08
	Carrier4   uint8   // D09
	Global     uint8   // D0A
	Reserved0B [1]byte // D0B
}

const (
	opl2ScaleEnvWithKeys  = uint8(0x10) // D00..D01
	opl2Sustain           = uint8(0x20) // D00..D01
	opl2Vibrato           = uint8(0x40) // D00..D01
	opl2Tremolo           = uint8(0x80) // D00..D01
	opl2AdditiveSynthesis = uint8(0x01) // D0A
)

// ModulatorKeyScaleRateSelect returns true if the modulator's envelope scales with keys
// If enabled, the envelopes of higher notes are played more quickly than those of lower notes.
func (o *OPL2Specs) ModulatorKeyScaleRateSelect() bool {
	return (o.Modulat0 & opl2ScaleEnvWithKeys) != 0
}

// ModulatorSustain returns true if the modulator's envelope sustain is enabled
// If enabled, the volume envelope stays at the sustain stage and does not enter the
// release stage of the envelope until a note-off event is encountered. Otherwise, it
// directly advances from the decay stage to the release stage without waiting for a
// note-off event.
func (o *OPL2Specs) ModulatorSustain() bool {
	return (o.Modulat0 & opl2Sustain) != 0
}

// ModulatorVibrato returns true if the modulator's vibrato is enabled
// If enabled, adds a vibrato effect with a depth of 7 cents (0.07 semitones).
// The rate of this vibrato is a static 6.4Hz.
func (o *OPL2Specs) ModulatorVibrato() bool {
	return (o.Modulat0 & opl2Vibrato) != 0
}

// ModulatorTremolo returns true if the modulator's tremolo is enabled
// If enabled, adds a tremolo effect with a depth of 1dB.
// The rate of this tremolo is a static 3.7Hz.
func (o *OPL2Specs) ModulatorTremolo() bool {
	return (o.Modulat0 & opl2Tremolo) != 0
}

// ModulatorFrequencyMultiplier returns the modulator's frequency multiplier
// Multiplies the frequency of the operator with a value between 0.5
// (pitched one octave down) and 15.
func (o *OPL2Specs) ModulatorFrequencyMultiplier() OPL2Multiple {
	return OPL2Multiple(o.Modulat0 & 0x0F)
}

// ModulatorKeyScaleLevel returns the key scale level
// Attenuates the output level of the operators towards higher pitch by the given amount
// (disabled, 1.5 dB / octave, 3 dB / octave, 6 dB / octave).
func (o *OPL2Specs) ModulatorKeyScaleLevel() OPL2KSL {
	v := o.Modulat1
	bit0 := (v & 0x80) >> 7
	bit1 := (v & 0x40) >> 6
	return OPL2KSL((bit1 << 1) | bit0)
}

// ModulatorVolume returns the modulator's volume
// The overall volume of the operator - if the modulator is in FM mode (i.e.: NOT in
// additive synthesis mode), this will instead be the total pitch depth.
func (o *OPL2Specs) ModulatorVolume() uint8 {
	return 63 - (o.Modulat1 & 0x3F)
}

// ModulatorAttackRate returns the modulator's attack rate
// Specifies how fast the volume envelope fades in from silence to peak volume.
func (o *OPL2Specs) ModulatorAttackRate() uint8 {
	return o.Modulat2 >> 4
}

// ModulatorDecayRate returns the modulator's decay rate
// Specifies how fast the volume envelope reaches the sustain volume after peaking.
func (o *OPL2Specs) ModulatorDecayRate() uint8 {
	return o.Modulat2 & 0x1F
}

// ModulatorSustainLevel returns the modulator's sustain level
// Specifies at which level the volume envelope is held before it is released.
func (o *OPL2Specs) ModulatorSustainLevel() uint8 {
	return 15 - (o.Modulat3 >> 4)
}

// ModulatorReleaseRate returns the modulator's release rate
// Specifies how fast the volume envelope fades out from the sustain level.
func (o *OPL2Specs) ModulatorReleaseRate() uint8 {
	return o.Modulat3 & 0x1F
}

// ModulatorWaveformSelection returns the modulator's waveform selection
func (o *OPL2Specs) ModulatorWaveformSelection() OPL2Waveform {
	return OPL2Waveform(o.Modulat4 & 0x07)
}

// CarrierKeyScaleRateSelect returns true if the carrier's envelope scales with keys
// If enabled, the envelopes of higher notes are played more quickly than those of lower notes.
func (o *OPL2Specs) CarrierKeyScaleRateSelect() bool {
	return (o.Carrier0 & opl2ScaleEnvWithKeys) != 0
}

// CarrierSustain returns true if the carrier's envelope sustain is enabled
// If enabled, the volume envelope stays at the sustain stage and does not enter the
// release stage of the envelope until a note-off event is encountered. Otherwise, it
// directly advances from the decay stage to the release stage without waiting for a
// note-off event.
func (o *OPL2Specs) CarrierSustain() bool {
	return (o.Carrier0 & opl2Sustain) != 0
}

// CarrierVibrato returns true if the carrier's vibrato is enabled
// If enabled, adds a vibrato effect with a depth of 7 cents (0.07 semitones).
// The rate of this vibrato is a static 6.4Hz.
func (o *OPL2Specs) CarrierVibrato() bool {
	return (o.Carrier0 & opl2Vibrato) != 0
}

// CarrierTremolo returns true if the carrier's tremolo is enabled
// If enabled, adds a tremolo effect with a depth of 1dB.
// The rate of this tremolo is a static 3.7Hz.
func (o *OPL2Specs) CarrierTremolo() bool {
	return (o.Carrier0 & opl2Tremolo) != 0
}

// CarrierFrequencyMultiplier returns the carrier's frequency multiplier
// Multiplies the frequency of the operator with a value between 0.5
// (pitched one octave down) and 15.
func (o *OPL2Specs) CarrierFrequencyMultiplier() OPL2Multiple {
	return OPL2Multiple(o.Carrier0 & 0x0F)
}

// CarrierKeyScaleLevel returns the key scale level
// Attenuates the output level of the operators towards higher pitch by the given amount
// (disabled, 1.5 dB / octave, 3 dB / octave, 6 dB / octave).
func (o *OPL2Specs) CarrierKeyScaleLevel() OPL2KSL {
	v := o.Carrier1
	bit0 := (v & 0x80) >> 7
	bit1 := (v & 0x40) >> 6
	return OPL2KSL((bit1 << 1) | bit0)
}

// CarrierVolume returns the carrier's volume
// The overall volume of the operator.
func (o *OPL2Specs) CarrierVolume() uint8 {
	return 63 - (o.Carrier1 & 0x3F)
}

// CarrierAttackRate returns the carrier's attack rate
// Specifies how fast the volume envelope fades in from silence to peak volume.
func (o *OPL2Specs) CarrierAttackRate() uint8 {
	return o.Carrier2 >> 4
}

// CarrierDecayRate returns the carrier's decay rate
// Specifies how fast the volume envelope reaches the sustain volume after peaking.
func (o *OPL2Specs) CarrierDecayRate() uint8 {
	return o.Carrier2 & 0x1F
}

// CarrierSustainLevel returns the carrier's sustain level
// Specifies at which level the volume envelope is held before it is released.
func (o *OPL2Specs) CarrierSustainLevel() uint8 {
	return 15 - (o.Carrier3 >> 4)
}

// CarrierReleaseRate returns the carrier's release rate
// Specifies how fast the volume envelope fades out from the sustain level.
func (o *OPL2Specs) CarrierReleaseRate() uint8 {
	return o.Carrier3 & 0x1F
}

// CarrierWaveformSelection returns the carrier's waveform selection
func (o *OPL2Specs) CarrierWaveformSelection() OPL2Waveform {
	return OPL2Waveform(o.Carrier4 & 0x07)
}

// ModulationFeedback returns the modulation feedback
func (o *OPL2Specs) ModulationFeedback() OPL2Feedback {
	return OPL2Feedback((o.Global >> 1) & 0x7)
}

// AdditiveSynthesis returns true if additive synthesis is enabled
func (o *OPL2Specs) AdditiveSynthesis() bool {
	return (o.Global & opl2AdditiveSynthesis) != 0
}

// SCRSAdlibHeader is the remaining header for S3M adlib instruments
type SCRSAdlibHeader struct {
	Reserved0D [3]byte
	OPL2       OPL2Specs
	Volume     Volume
	// in ST3 doc as poorly documented value 'Dsk'
	// maybe this has a bit for the keyboard split method selection (NTS in YM3812/YMF262)?
	Reserved1D byte
	Reserved1E [2]byte
	C2Spd      HiLo32
	Reserved24 [12]byte
	SampleName [28]byte
	SCRI       [4]uint8
}

// GetSampleName returns a string representation of the data stored in the SampleName field
func (h *SCRSAdlibHeader) GetSampleName() string {
	return util.GetString(h.SampleName[:])
}
