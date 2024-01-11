package myls

import "os"

// Check if a file is executable
func IsExecFile(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	mode := info.Mode()
	return mode&0111 != 0
}
