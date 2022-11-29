package it

// IMPMFlags is a set of flags describing various features in the IT file
type IMPMFlags uint16

const (
	// IMPMFlagStereo :: On = Stereo, Off = Mono (panning enablement flag)
	IMPMFlagStereo = IMPMFlags(1 << 0)
	// IMPMFlagVol0Optimizations :: If on, no mixing occurs if the volume at mixing time is 0 (redundant v1.04+)
	IMPMFlagVol0Optimizations = IMPMFlags(1 << 1)
	// IMPMFlagUseInstruments :: On = Use instruments, Off = Use samples
	IMPMFlagUseInstruments = IMPMFlags(1 << 2)
	// IMPMFlagLinearSlides :: On = Linear slides, Off = Amiga slides
	IMPMFlagLinearSlides = IMPMFlags(1 << 3)
	// IMPMFlagOldEffects :: On = Old Effects, Off = IT Effects
	IMPMFlagOldEffects = IMPMFlags(1 << 4)
	// IMPMFlagEFGLinking :: On = Link Effect G's memory with Effect E/F
	IMPMFlagEFGLinking = IMPMFlags(1 << 5)
	// IMPMFlagMidiPitchController :: Use MIDI pitch controller, Pitch depth given by PitchWheelDepth
	IMPMFlagMidiPitchController = IMPMFlags(1 << 6)
	// IMPMFlagReqEmbedMidi :: Request embedded MIDI configuration
	IMPMFlagReqEmbedMidi = IMPMFlags(1 << 7)
)

// IsStereo returns true if stereo (panning) is enabled
func (f IMPMFlags) IsStereo() bool {
	return (f & IMPMFlagStereo) != 0
}

// IsVol0Optimizations returns true if vol-0 optimization is enabled
func (f IMPMFlags) IsVol0Optimizations() bool {
	return (f & IMPMFlagVol0Optimizations) != 0
}

// IsUseInstruments returns true if use-instruments (instead of samples) is enabled
func (f IMPMFlags) IsUseInstruments() bool {
	return (f & IMPMFlagUseInstruments) != 0
}

// IsLinearSlides returns true if linear slides is enabled
func (f IMPMFlags) IsLinearSlides() bool {
	return (f & IMPMFlagLinearSlides) != 0
}

// IsOldEffects returns true if old-style effects are enabled
func (f IMPMFlags) IsOldEffects() bool {
	return (f & IMPMFlagOldEffects) != 0
}

// IsEFGLinking returns true if effect E/F/G linking is enabled
func (f IMPMFlags) IsEFGLinking() bool {
	return (f & IMPMFlagEFGLinking) != 0
}

// IsMidiPitchController returns true if midi pitch controller is enabled
func (f IMPMFlags) IsMidiPitchController() bool {
	return (f & IMPMFlagMidiPitchController) != 0
}

// IsReqEmbedMidi returns true if request embedded midi configuration is enabled
func (f IMPMFlags) IsReqEmbedMidi() bool {
	return (f & IMPMFlagReqEmbedMidi) != 0
}

// IMPMSpecialFlags is a set of flags describing various special features in the IT file
type IMPMSpecialFlags uint16

const (
	// IMPMSpecialFlagMessageAttached :: On = song message attached
	IMPMSpecialFlagMessageAttached = IMPMSpecialFlags(1 << 0)
	// IMPMSpecialFlagHistoryIncluded :: On = history data included (maybe)
	IMPMSpecialFlagHistoryIncluded = IMPMSpecialFlags(1 << 1)
	// IMPMSpecialFlagHighlightDataIncluded :: On = highlight data included
	IMPMSpecialFlagHighlightDataIncluded = IMPMSpecialFlags(1 << 2)
	// IMPMSpecialFlagEmbedMidi :: MIDI configuration embedded
	IMPMSpecialFlagEmbedMidi = IMPMSpecialFlags(1 << 3)
	// IMPMSpecialFlagReservedBit4 :: Reserved
	IMPMSpecialFlagReservedBit4 = IMPMSpecialFlags(1 << 4)
	// IMPMSpecialFlagReservedBit5 :: Reserved
	IMPMSpecialFlagReservedBit5 = IMPMSpecialFlags(1 << 5)
	// IMPMSpecialFlagReservedBit6 :: Reserved
	IMPMSpecialFlagReservedBit6 = IMPMSpecialFlags(1 << 6)
	// IMPMSpecialFlagReservedBit7 :: Reserved
	IMPMSpecialFlagReservedBit7 = IMPMSpecialFlags(1 << 7)
	// IMPMSpecialFlagReservedBit8 :: Reserved
	IMPMSpecialFlagReservedBit8 = IMPMSpecialFlags(1 << 8)
	// IMPMSpecialFlagReservedBit9 :: Reserved
	IMPMSpecialFlagReservedBit9 = IMPMSpecialFlags(1 << 9)
	// IMPMSpecialFlagReservedBit10 :: Reserved
	IMPMSpecialFlagReservedBit10 = IMPMSpecialFlags(1 << 10)
	// IMPMSpecialFlagReservedBit11 :: Reserved
	IMPMSpecialFlagReservedBit11 = IMPMSpecialFlags(1 << 11)
	// IMPMSpecialFlagReservedBit12 :: Reserved
	IMPMSpecialFlagReservedBit12 = IMPMSpecialFlags(1 << 12)
	// IMPMSpecialFlagReservedBit13 :: Reserved
	IMPMSpecialFlagReservedBit13 = IMPMSpecialFlags(1 << 13)
	// IMPMSpecialFlagReservedBit14 :: Reserved
	IMPMSpecialFlagReservedBit14 = IMPMSpecialFlags(1 << 14)
	// IMPMSpecialFlagReservedBit15 :: Reserved
	IMPMSpecialFlagReservedBit15 = IMPMSpecialFlags(1 << 15)
)

// IsMessageAttached returns true if there is a special message attached to the file
func (sf IMPMSpecialFlags) IsMessageAttached() bool {
	return (sf & IMPMSpecialFlagMessageAttached) != 0
}

// IsHistoryIncluded returns true if there is a history block following the pattern parapointres in the file
func (sf IMPMSpecialFlags) IsHistoryIncluded() bool {
	return (sf & IMPMSpecialFlagHistoryIncluded) != 0
}

// IsHighlightDataIncluded returns true if there is a highlight data attached to the file
func (sf IMPMSpecialFlags) IsHighlightDataIncluded() bool {
	return (sf & IMPMSpecialFlagHighlightDataIncluded) != 0
}

// IsEmbedMidi returns true if embedded midi configuration is enabled
func (sf IMPMSpecialFlags) IsEmbedMidi() bool {
	return (sf & IMPMSpecialFlagEmbedMidi) != 0
}

// SampleFlags defines the flags associated to the IT format sample header
type SampleFlags uint8

const (
	// SampleFlagSampleExists :: On = sample associated with header
	SampleFlagSampleExists = SampleFlags(1 << 0)
	// SampleFlag16Bit :: On = 16 bit, Off = 8 bit
	SampleFlag16Bit = SampleFlags(1 << 1)
	// SampleFlagStereo :: On = stereo, Off = mono
	SampleFlagStereo = SampleFlags(1 << 2)
	// SampleFlagCompressed :: On = compressed samples
	SampleFlagCompressed = SampleFlags(1 << 3)
	// SampleFlagUseLoop :: On = Use loop
	SampleFlagUseLoop = SampleFlags(1 << 4)
	// SampleFlagUseSustainLoop :: On = Use sustain loop
	SampleFlagUseSustainLoop = SampleFlags(1 << 5)
	// SampleFlagPingPongLoop :: On = Ping Pong loop, Off = Forwards loop
	SampleFlagPingPongLoop = SampleFlags(1 << 6)
	// SampleFlagPingPongSustainLoop :: On = Ping Pong Sustain loop, Off = Forwards Sustain loop
	SampleFlagPingPongSustainLoop = SampleFlags(1 << 7)
)

// DoesSampleExist returns true if the sample header has a sample associated with it.
func (sf SampleFlags) DoesSampleExist() bool {
	return (sf & SampleFlagSampleExists) != 0
}

// Is16Bit returns true if the sample is 16-bit
func (sf SampleFlags) Is16Bit() bool {
	return (sf & SampleFlag16Bit) != 0
}

// IsStereo returns true if the sample is stereo
func (sf SampleFlags) IsStereo() bool {
	return (sf & SampleFlagStereo) != 0
}

// IsCompressed returns true if the sample data is compressed
func (sf SampleFlags) IsCompressed() bool {
	return (sf & SampleFlagCompressed) != 0
}

// IsLoopEnabled returns true if the sample loop is enabled
func (sf SampleFlags) IsLoopEnabled() bool {
	return (sf & SampleFlagUseLoop) != 0
}

// IsSustainLoopEnabled returns true if the sample sustain loop is enabled
func (sf SampleFlags) IsSustainLoopEnabled() bool {
	return (sf & SampleFlagUseSustainLoop) != 0
}

// IsLoopPingPong returns true if the sample loop mode is ping-pong
func (sf SampleFlags) IsLoopPingPong() bool {
	return (sf & SampleFlagPingPongLoop) != 0
}

// IsSustainLoopPingPong returns true if the sample sustain loop mode is ping-pong
func (sf SampleFlags) IsSustainLoopPingPong() bool {
	return (sf & SampleFlagPingPongSustainLoop) != 0
}

// ConvertFlags defines the flags associated to the IT format sample data conversion settings
type ConvertFlags uint8

const (
	// ConvertFlagSignedSamples :: On = Samples are signed. Off = Samples are unsigned
	ConvertFlagSignedSamples = ConvertFlags(1 << 0)
	// ConvertFlagBigEndian :: On = 16 bit big endian, Off = 16 bit little endian
	ConvertFlagBigEndian = ConvertFlags(1 << 1)
	// ConvertFlagSampleDelta :: On = stored as delta values, Off = stored as PCM values
	ConvertFlagSampleDelta = ConvertFlags(1 << 2)
	// ConvertFlagByteDelta :: On = stored as byte delta values
	ConvertFlagByteDelta = ConvertFlags(1 << 3)
	// ConvertFlagTXWave12Bit :: On = stored as TX-Wave 12-bit values
	ConvertFlagTXWave12Bit = ConvertFlags(1 << 4)
	// ConvertFlagChannelSelectionPrompt :: On = Left/Right/All Stereo prompt
	ConvertFlagChannelSelectionPrompt = ConvertFlags(1 << 5)
	// ConvertFlagReserved6 :: Reserved
	ConvertFlagReserved6 = ConvertFlags(1 << 6)
	// ConvertFlagReserved7 :: Reserved
	ConvertFlagReserved7 = ConvertFlags(1 << 7)
)

// IsSignedSamples returns true if the sample data is in signed values
func (sf ConvertFlags) IsSignedSamples() bool {
	return (sf & ConvertFlagSignedSamples) != 0
}

// IsBigEndian returns true if the sample is 16-bit big endian
func (sf ConvertFlags) IsBigEndian() bool {
	return (sf & ConvertFlagBigEndian) != 0
}

// IsSampleDelta returns true if the sample is stored in sample delta format
func (sf ConvertFlags) IsSampleDelta() bool {
	return (sf & ConvertFlagSampleDelta) != 0
}

// IsByteDelta returns true if the sample data is stored in byte delta format
func (sf ConvertFlags) IsByteDelta() bool {
	return (sf & ConvertFlagByteDelta) != 0
}

// IsTXWave12Bit returns true if the sample loop is stored in TX-Wave 12-bit format
func (sf ConvertFlags) IsTXWave12Bit() bool {
	return (sf & ConvertFlagTXWave12Bit) != 0
}

// IsChannelSelectPrompt returns true if the channel selection prompt is enabled
func (sf ConvertFlags) IsChannelSelectPrompt() bool {
	return (sf & ConvertFlagChannelSelectionPrompt) != 0
}

// IMPIOldFlags is the flagset for IMPIInstrumentOld instruments
type IMPIOldFlags uint8

const (
	// IMPIOldFlagUseVolumeEnvelope :: On = Use volume envelope
	IMPIOldFlagUseVolumeEnvelope = IMPIOldFlags(1 << 0)
	// IMPIOldFlagUseVolumeLoop :: On = Use volume loop
	IMPIOldFlagUseVolumeLoop = IMPIOldFlags(1 << 1)
	// IMPIOldFlagUseSustainVolumeLoop :: On = Use sustain volume loop
	IMPIOldFlagUseSustainVolumeLoop = IMPIOldFlags(1 << 2)
)

// EnvelopeFlags is the flagset for new instrument envelopes
type EnvelopeFlags uint8

const (
	// EnvelopeFlagEnvelopeOn :: On = Use envelope
	EnvelopeFlagEnvelopeOn = EnvelopeFlags(1 << 0)
	// EnvelopeFlagLoopOn :: On = Use loop
	EnvelopeFlagLoopOn = EnvelopeFlags(1 << 1)
	// EnvelopeFlagSustainLoopOn :: On = Use sustain loop
	EnvelopeFlagSustainLoopOn = EnvelopeFlags(1 << 2)
)

// ChannelDataFlags is a set of flags specifying what data is available in the packed pattern and/or in this channel data
type ChannelDataFlags uint8

const (
	// ChannelDataFlagNote :: On = Note is available
	ChannelDataFlagNote = ChannelDataFlags(1 << 0)
	// ChannelDataFlagInstrument :: On = Instrument is available
	ChannelDataFlagInstrument = ChannelDataFlags(1 << 1)
	// ChannelDataFlagVolPan :: On = VolPan is available
	ChannelDataFlagVolPan = ChannelDataFlags(1 << 2)
	// ChannelDataFlagCommand :: On = Command/CommandData is available
	ChannelDataFlagCommand = ChannelDataFlags(1 << 3)
	// ChannelDataFlagUseLastNote :: On = Use the previous Note value in the channel
	ChannelDataFlagUseLastNote = ChannelDataFlags(1 << 4)
	// ChannelDataFlagUseLastInstrument :: On = Use the previous Instrument value in the channel
	ChannelDataFlagUseLastInstrument = ChannelDataFlags(1 << 5)
	// ChannelDataFlagUseLastVolPan :: On = Use the previous VolPan value in the channel
	ChannelDataFlagUseLastVolPan = ChannelDataFlags(1 << 6)
	// ChannelDataFlagUseLastCommand :: On = Use the previous Command/CommandData value in the channel
	ChannelDataFlagUseLastCommand = ChannelDataFlags(1 << 7)
)

// HasNote returns true if there is a Note value present
func (cdf ChannelDataFlags) HasNote() bool {
	return (cdf & ChannelDataFlagNote) != 0
}

// HasInstrument returns true if there is a Instrument value present
func (cdf ChannelDataFlags) HasInstrument() bool {
	return (cdf & ChannelDataFlagInstrument) != 0
}

// HasVolPan returns true if there is a VolPan value present
func (cdf ChannelDataFlags) HasVolPan() bool {
	return (cdf & ChannelDataFlagVolPan) != 0
}

// HasCommand returns true if there is a Command value present
func (cdf ChannelDataFlags) HasCommand() bool {
	return (cdf & ChannelDataFlagCommand) != 0
}

// IsUseLastNote returns true if the previous Note value should be used
func (cdf ChannelDataFlags) IsUseLastNote() bool {
	return (cdf & ChannelDataFlagUseLastNote) != 0
}

// IsUseLastInstrument returns true if the previous Instrument value should be used
func (cdf ChannelDataFlags) IsUseLastInstrument() bool {
	return (cdf & ChannelDataFlagUseLastInstrument) != 0
}

// IsUseLastVolPan returns true if the previous VolPan value should be used
func (cdf ChannelDataFlags) IsUseLastVolPan() bool {
	return (cdf & ChannelDataFlagUseLastVolPan) != 0
}

// IsUseLastCommand returns true if the previous Command/CommandData value should be used
func (cdf ChannelDataFlags) IsUseLastCommand() bool {
	return (cdf & ChannelDataFlagUseLastCommand) != 0
}
