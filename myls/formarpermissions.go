package myls

import "os"

// Format a single permission set (owner, group, or other)
func FormatPermission(mode os.FileMode, shift uint) string {
	// values for different system permissions
	const (
		read    = "r"
		write   = "w"
		execute = "x"
		none    = "-"

	)
	perm := ""
	if mode&(1<<(shift+2)) != 0 { // check for read permission
		perm += read
	} else {
		perm += none
	}
	if mode&(1<<(shift+1)) != 0 { // check for write permission
		perm += write
	} else {
		perm += none
	}
	if mode&(1<<shift) != 0 { // check for executable permission
		perm += execute
	} else {
		perm += none
	}

	return perm
}
