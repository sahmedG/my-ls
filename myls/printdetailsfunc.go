package myls

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

func PrintDetails(path, name string, modTime time.Time, isDir, isSymlink, isExecFile, isArchive, isDevice bool, sizeBytes int64, Major, Minor uint32) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get file permissions in string format (e.g., "drwxrwxr-x")
	perm := FormatPermissions(info.Mode())

	// Get file owner and group
	owner, err := user.LookupId(strconv.Itoa(int(info.Sys().(*syscall.Stat_t).Uid)))
	if err != nil {
		log.Fatal(err)
	}
	group, err := user.LookupGroupId(strconv.Itoa(int(info.Sys().(*syscall.Stat_t).Gid)))
	if err != nil {
		log.Fatal(err)
	}
	// color values
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

	// Format the modification time
	modTimeStr := modTime.Format("Jan _2 15:04")

	// Print the formatted information each colored according to each type
	if isDir {
		fmt.Printf("%s %2d %-2s %-2s %7d %s %s %s %s",
			perm,
			info.Sys().(*syscall.Stat_t).Nlink,
			owner.Username,
			group.Name,
			info.Size(),
			modTimeStr,
			colors.Blue,
			name,
			colors.Reset,
		)
	} else if isExecFile && !isDir && !isSymlink {
		fmt.Printf("%s %2d %-2s %-2s %7d %s %s %s %s",
			perm,
			info.Sys().(*syscall.Stat_t).Nlink,
			owner.Username,
			group.Name,
			info.Size(),
			modTimeStr,
			colors.Green,
			name,
			colors.Reset,
		)
	} else if isSymlink {
		fmt.Printf("%s %2d %-2s %-2s %7d %s %s %s %s",
			perm,
			info.Sys().(*syscall.Stat_t).Nlink,
			owner.Username,
			group.Name,
			info.Size(),
			modTimeStr,
			colors.Cyan,
			name,
			colors.Reset,
		)

	} else if isArchive {
		fmt.Printf("%s %2d %-2s %-2s %7d %s %s %s %s",
			perm,
			info.Sys().(*syscall.Stat_t).Nlink,
			owner.Username,
			group.Name,
			info.Size(),
			modTimeStr,
			colors.Red,
			name,
			colors.Reset,
		)

	} else if isDevice {
		fmt.Printf("%s %2d %-2s %-2s %d, %d %s %s %s %s",
			perm,
			info.Sys().(*syscall.Stat_t).Nlink,
			owner.Username,
			group.Name,
			Major,
			Minor,
			modTimeStr,
			colors.DeviceColor,
			name,
			colors.Reset,
		)

	} else {
		if name == "." || name == ".." {
			fmt.Printf("%s %2d %-2s %-2s %7d %s %s %s %s",
				perm,
				info.Sys().(*syscall.Stat_t).Nlink,
				owner.Username,
				group.Name,
				info.Size(),
				modTimeStr,
				colors.Blue,
				name,
				colors.Reset,
			)
		} else {
			fmt.Printf("%s %2d %-2s %-2s %7d %s %s",
				perm,
				info.Sys().(*syscall.Stat_t).Nlink,
				owner.Username,
				group.Name,
				info.Size(),
				modTimeStr,
				name,
			)
		}

	}
	if isSymlink {
		target, err := os.Readlink(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" -> %s", target)
	}
	fmt.Println()
}
