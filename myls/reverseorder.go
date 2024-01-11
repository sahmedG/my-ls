package myls

func ReverseOrder(files []FileInfo) []FileInfo {
	files2 := []FileInfo{}
	for i := len(files) - 1; i >= 0; i-- {
		temp := files[i]
		files2 = append(files2, temp)
	}
	return files2
}
