package myls

import "time"
// struct to hold files/folders details
type FileInfo struct {
	Name       string
	ModTime    time.Time
	IsDir      bool
	IsSymlink  bool
	IsExecFile bool
	IsArchive  bool
	IsDevice   bool
	IsGraphic  bool
	SizeBytes  int64
	Major      uint32
	Minor      uint32
}

// color struct
type ColorCodes struct {
	Blue        string
	Green       string
	Cyan        string
	Red         string
	Reset       string
	DeviceColor string
	Graphic     string
	Bold        string
}


// type Flagsoptions struct {
// 	ShowHidden  bool
// 	Recursive   bool
// 	SortByTime  bool
// 	ShowDetails bool
// 	Reverse     bool
// }
