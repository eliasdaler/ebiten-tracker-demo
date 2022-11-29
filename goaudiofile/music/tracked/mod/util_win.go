//go:build windows
// +build windows

package mod

import "syscall"

// BE16ToLE16 converts a big-endian uint16 to a little-endian uint16
func BE16ToLE16(be uint16) uint16 {
	return syscall.Ntohs(be)
}
