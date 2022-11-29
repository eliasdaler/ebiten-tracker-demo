package s3m

// PackedPattern is the S3M packed pattern definition
type PackedPattern struct {
	Length uint16
	Data   []byte
}

// PatternFlags is a flagset (and channel id) for data in the channel
type PatternFlags uint8

const (
	// PatternFlagCommand is the flag that denotes existence of a command on the channel
	PatternFlagCommand = PatternFlags(0x80)
	// PatternFlagVolume is the flag that denotes existence of a volume on the channel
	PatternFlagVolume = PatternFlags(0x40)
	// PatternFlagNote is the flag that denotes existence of a note on the channel
	PatternFlagNote = PatternFlags(0x20)
)

// HasCommand returns true if there exists a command on the channel
func (w PatternFlags) HasCommand() bool {
	return (w & PatternFlagCommand) != 0
}

// HasVolume returns true if there exists a volume on the channel
func (w PatternFlags) HasVolume() bool {
	return (w & PatternFlagVolume) != 0
}

// HasNote returns true if there exists a note on the channel
func (w PatternFlags) HasNote() bool {
	return (w & PatternFlagNote) != 0
}

// Channel returns the channel ID for this channel
func (w PatternFlags) Channel() uint8 {
	return uint8(w) & 0x1F
}

// C2SPD defines the C-2 (or in some players cases C-4) note sampling rate
type C2SPD uint16

// Volume defines a volume value
type Volume uint8

const (
	// DefaultC2Spd is the default C2SPD for S3M files
	DefaultC2Spd = C2SPD(8363)

	// DefaultVolume is the default volume for many things in S3M files
	DefaultVolume = Volume(64)

	// EmptyVolume is a volume value that uses the volume value on the instrument
	EmptyVolume = Volume(255)
)

// Semitone is a specific note in a 12-step scale of notes / octaves
type Semitone uint8

// Key is a note key component
type Key uint8

const (
	// KeyC is C
	KeyC = Key(0 + iota)
	// KeyCSharp is C#
	KeyCSharp
	// KeyD is D
	KeyD
	// KeyDSharp is D#
	KeyDSharp
	// KeyE is E
	KeyE
	// KeyF is F
	KeyF
	// KeyFSharp is F#
	KeyFSharp
	// KeyG is G
	KeyG
	// KeyGSharp is G#
	KeyGSharp
	// KeyA is A
	KeyA
	// KeyASharp is A#
	KeyASharp
	// KeyB is B
	KeyB
	//KeyInvalid1 is invalid
	KeyInvalid1
	//KeyInvalid2 is invalid
	KeyInvalid2
	//KeyInvalid3 is invalid
	KeyInvalid3
	//KeyInvalid4 is invalid
	KeyInvalid4
)

// IsInvalid returns true if the key is invalid
func (k Key) IsInvalid() bool {
	switch k {
	case KeyInvalid1, KeyInvalid2, KeyInvalid3, KeyInvalid4:
		return true
	default:
		return false
	}
}

// Octave is the octave the key is in
type Octave uint8

// Note is a combination of key and octave
type Note uint8

const (
	// EmptyNote denotes an empty note
	EmptyNote = Note(255)
	// StopNote denotes a stop for the instrument
	StopNote = Note(254)
)

// Key returns the key component of the note
func (n Note) Key() Key {
	return Key(n & 0x0F)
}

// Octave returns the octave component of the note
func (n Note) Octave() Octave {
	return Octave((n & 0xF0) >> 4)
}

// IsStop returns true if the note is a stop
func (n Note) IsStop() bool {
	return n == StopNote
}

// IsInvalid returns true if the note is invalid in any way (or is a stop)
func (n Note) IsInvalid() bool {
	return n == EmptyNote || n.IsStop() || n.Key().IsInvalid()
}

// Semitone returns the semitone value for the note
func (n Note) Semitone() Semitone {
	key := Semitone(n.Key())
	octave := Semitone(n.Octave())
	return Semitone(octave*12 + key)
}
