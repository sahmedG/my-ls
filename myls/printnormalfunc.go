package myls

import (
	"fmt"
	"time"
)

func Printnormal(Name string, modTime time.Time, IsDir, IsSymlink, IsExecFile, IsArchive, IsDevice, IsGraphic, recursive bool) {
	// colors values used with different files and folders types
	colors := ColorCodes{
		Blue:        "\033[34m",
		Green:       "\033[32m",
		Cyan:        "\033[36m",
		Red:         "\033[31m",
		Reset:       "\033[0m",
		DeviceColor: "\033[33;40m",
		Graphic:     "\u001b[35m",
		Bold:        "\u001b[1m",
	}
	// checking each file/folder type and print the details with proper color value
	if IsExecFile && !IsDir && !IsSymlink {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Green, Name, colors.Reset)
	} else if IsDir {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Blue, Name, colors.Reset)
	} else if IsArchive {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Red, Name, colors.Reset)
	} else if IsSymlink {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Cyan, Name, colors.Reset)
	} else if IsDevice {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.DeviceColor, Name, colors.Reset)
	} else if IsGraphic && !IsArchive && !IsDevice && !IsDir {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Graphic, Name, colors.Reset)
	} else if Name == "." || Name == ".." {
		fmt.Printf("%s%s%s%s ", colors.Bold, colors.Blue, Name, colors.Reset)
	} else {
		fmt.Printf("%s%s%s ", colors.Reset, Name, colors.Reset)
	}
}
