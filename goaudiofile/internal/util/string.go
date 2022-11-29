package util

import "bytes"

// GetString converts a fixed length byte array with embedded nulls into a string
func GetString(data []byte) string {
	n := bytes.Index(data, []byte{0})
	if n == -1 {
		n = len(data)
	}
	s := string(data[:n])
	return s
}
