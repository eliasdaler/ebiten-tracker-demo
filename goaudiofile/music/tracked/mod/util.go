package mod

// WordLength is a count of WORD (uint16) sized values, stored in BigEndian format
type WordLength uint16

// Value returns the actual length described by this WordLength
func (m WordLength) Value() int {
	v := BE16ToLE16(uint16(m))
	return int(v) << 1
}
