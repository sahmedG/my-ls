package myls

import (
	"io/fs"
	"os"
)

// Format file permissions into a string
func FormatPermissions(mode os.FileMode) string {
	// values for different files and folders types
	const (
		directory = "d"
		link      = "l"
		block     = "b"
		char      = "c"
		other     = "-"
	)
	permString := other

	if mode&fs.ModeDir != 0 { //check if type in directory
		permString = directory
	} else if mode&fs.ModeSymlink != 0 { // check if type symlink
		permString = link
	} else if mode&os.ModeCharDevice != 0 { // check if character special file
		permString = char
	} else if mode&os.ModeDevice !=0 { // check if block special file
		permString = block
	}

	// Owner permissions
	permString += FormatPermission(mode, 6)
	// Group permissions
	permString += FormatPermission(mode, 3)
	// Other permissions
	permString += FormatPermission(mode, 0)

	return permString
}
