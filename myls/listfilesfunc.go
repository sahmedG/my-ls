package myls

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
	"syscall"
)

func ListFiles(path string, showHidden, recursive, sortByTime, showDetails, reverse bool, length int) error {

	file, err := os.Lstat(path)
	if err != nil {
		return err
	}

	Totalsize := 0
	isArchive := false
	isGraphic := false
	isDevice := false

	if file.Mode()&os.ModeSymlink != 0 && showDetails {
		isDevice = file.Mode()&os.ModeDevice != 0
		isDir := file.IsDir()
		isSymlink := file.Mode()&fs.ModeSymlink == fs.ModeSymlink
		isExecFile := IsExecFile(path)
		sizeBytes := file.Size()
		file := FileInfo{
			Name:       path,
			ModTime:    file.ModTime(),
			IsDir:      isDir,
			IsSymlink:  isSymlink,
			IsExecFile: isExecFile,
			IsArchive:  isArchive,
			IsDevice:   isDevice,
			IsGraphic:  isGraphic,
			SizeBytes:  sizeBytes,
		}
		PrintDetails(path, file.Name, file.ModTime, file.IsDir, file.IsSymlink, file.IsExecFile, file.IsArchive, file.IsDevice, file.SizeBytes, 0, 0)
	} else {
		check, err := os.Stat(path)
		if err != nil {
			return err
		}
		if recursive {
			fmt.Printf("%s%s%s\n", "\033[0m", path+":", "\033[0m")
		}
		if showDetails && showHidden && length == 2 {
			fmt.Printf("%s%s%s\n", "\033[0m", path+":", "\033[0m")
		}
		archiveExtensions := []string{".tar", ".gz", ".zip", ".tar.gz", ".deb", ".rpm"}
		GraphicExtensions := []string{".jpg", ".png", ".jpeg", ".gif"}
		var files []FileInfo
		var dirs []FileInfo
		var dirEntries []fs.DirEntry

		if check.IsDir() && !showHidden {
			if path[0] != '.' {
				dirEntries, _ = os.ReadDir(path)
			}
		} else if check.IsDir() && showHidden {
			dirEntries, _ = os.ReadDir(path)
			sizeBytes := check.Size()
			file := FileInfo{
				Name:       ".",
				ModTime:    check.ModTime(),
				IsDir:      false,
				IsSymlink:  false,
				IsExecFile: false,
				IsArchive:  false,
				IsDevice:   false,
				SizeBytes:  sizeBytes,
			}
			files = append(files, file)
			file = FileInfo{
				Name:       "..",
				ModTime:    check.ModTime(),
				IsDir:      false,
				IsSymlink:  false,
				IsExecFile: false,
				IsArchive:  false,
				IsDevice:   false,
				SizeBytes:  sizeBytes,
			}
			files = append(files, file)
		} else {
			isDir := check.IsDir()
			isSymlink := check.Mode()&fs.ModeSymlink == fs.ModeSymlink
			isExecFile := IsExecFile(path)
			sizeBytes := check.Size()

			file := FileInfo{
				Name:       check.Name(),
				ModTime:    check.ModTime(),
				IsDir:      isDir,
				IsSymlink:  isSymlink,
				IsExecFile: isExecFile,
				IsArchive:  false,
				IsDevice:   false,
				SizeBytes:  sizeBytes,
			}

			files = append(files, file)

			if !showDetails && !recursive && !reverse && !showHidden && !sortByTime {
				Printnormal(file.Name, file.ModTime, file.IsDir, file.IsSymlink, file.IsExecFile, file.IsArchive, file.IsDevice, file.IsGraphic, recursive)
			} else if showDetails && !recursive && !reverse && !sortByTime {
				PrintDetails(path, file.Name, file.ModTime, file.IsDir, file.IsSymlink, file.IsExecFile, file.IsArchive, file.IsDevice, file.SizeBytes, 0, 0)
				fmt.Println()
				return nil
			}
			fmt.Println()
			return nil
		}

		// Loop through the files and categorize them into files and directories
		for _, dirEntry := range dirEntries {
			fileName := dirEntry.Name()
			// Check if the file is a hidden file or if it's a directory
			if !showHidden && fileName[0] == '.' {
				continue // Skip hidden files if not requested
			}

			filePath := path + "/" + fileName
			info, err := os.Lstat(filePath)
			if err != nil {
				continue
			}
			for _, ext := range archiveExtensions {
				if strings.HasSuffix(fileName, ext) {
					isArchive = true
					break
				} else {
					isArchive = false
				}
			}
			for _, ext := range GraphicExtensions {
				if strings.HasSuffix(fileName, ext) {
					isGraphic = true
					break
				} else {
					isGraphic = false
				}
			}
			var Major_size uint32
			var Minor_size uint32

			isDevice = info.Mode()&os.ModeDevice != 0
			if isDevice {
				stat := syscall.Stat_t{}
				err = syscall.Stat(path+fileName, &stat)
				if err != nil {
					fmt.Println("Error", err)
				}

				Major_size = Major(stat.Rdev)
				Minor_size = Minor(stat.Rdev)
			}
			isDir := info.IsDir()
			isSymlink := info.Mode()&fs.ModeSymlink == fs.ModeSymlink
			isExecFile := IsExecFile(filePath)
			sizeBytes := info.Size()
			file := FileInfo{
				Name:       fileName,
				ModTime:    info.ModTime(),
				IsDir:      isDir,
				IsSymlink:  isSymlink,
				IsExecFile: isExecFile,
				IsArchive:  isArchive,
				IsDevice:   isDevice,
				IsGraphic:  isGraphic,
				SizeBytes:  sizeBytes,
				Major:      Major_size,
				Minor:      Minor_size,
			}
			files = append(files, file)
		}

		for _, name := range files {
			fileName := name.Name
			filePath := path + "/" + fileName
			var info fs.FileInfo
			info, err = os.Stat(filePath)
			if err != nil {
				continue
			}
			infouser := info.Sys()
			if stat, ok := infouser.(*syscall.Stat_t); ok {
				Totalsize += int((stat.Blocks + 4096 - 1) / 4096 * (4096 / 1024))
			}
		}
		// sort slice to show the files and folders in proper order under -a flag
		if showHidden {
			for i := 0; i < len(files); i++ {
				for j := i + 1; j < len(files)-1; j++ {
					if len(files[i].Name) > 1 && len(files[j].Name) > 1 && files[i].Name[0] == 46 && files[j].Name[0] != 46 {
						if files[i].Name[1] > files[j].Name[0] {
							files[i], files[j] = files[j], files[i]
						}
					} else if files[i].Name > files[j].Name {
						files[i], files[j] = files[j], files[i]
					}
				}
			}
		}
		if showDetails {
			fmt.Println("total", Totalsize)
		}

		// Sort the slices based on modification time
		if sortByTime {
			sort.SliceStable(files, func(i, j int) bool {
				return files[i].ModTime.After(files[j].ModTime)
			})
			sort.SliceStable(dirs, func(i, j int) bool {
				return dirs[i].ModTime.After(dirs[j].ModTime)
			})
		}
		// Sort the slices in reverse order
		if reverse {
			files = ReverseOrder(files)
		}
		for i := 0; i < len(files); i++ {
			if showDetails {
				PrintDetails(path+"/"+files[i].Name, files[i].Name, files[i].ModTime, files[i].IsDir, files[i].IsSymlink, files[i].IsExecFile, files[i].IsArchive, files[i].IsDevice, files[i].SizeBytes, files[i].Major, files[i].Minor)
			} else {
				Printnormal(files[i].Name, files[i].ModTime, files[i].IsDir, files[i].IsSymlink, files[i].IsExecFile, files[i].IsArchive, files[i].IsDevice, files[i].IsGraphic, recursive)
			}

		}
		if recursive {
			fmt.Println()
		}
		if recursive {
			for i := 0; i < len(files); i++ {
				if files[i].IsDir {
					subdir := path + "/" + files[i].Name
					fmt.Println()
					ListFiles(subdir, showHidden, recursive, sortByTime, showDetails, reverse, length) // Recursively list subdirectory
				}
			}
		}
	}
	if !showDetails && !recursive {
		fmt.Println()
	}
	return nil
}

func Major(dev uint64) uint32 {

	major := uint32((dev & 0x00000000000fff00) >> 8)

	major |= uint32((dev & 0xfffff00000000000) >> 32)

	return major

}

func Minor(dev uint64) uint32 {

	minor := uint32((dev & 0x00000000000000ff) >> 0)

	minor |= uint32((dev & 0x00000ffffff00000) >> 12)

	return minor

}
