package main

import (
	"fmt"
	"os"
	"sort"
	"Myls/myls"
)

func main() {
	//all flags that are accepted
	showHidden := false
	recursive := false
	sortByTime := false
	showDetails := false
	reverse := false
	args := os.Args[1:]

	// Process command-line options
	sort.Strings(args)
	for len(args) > 0 && args[0][0] == '-' {
		if len(args[0][1:]) == 0 {
			return
		}
		for _, ch := range args[0][1:] {
			switch ch {
			case 'a':
				showHidden = true
			case 'R':
				recursive = true
			case 'r':
				reverse = true
			case 't':
				sortByTime = true
			case 'l':
				showDetails = true
			default:
				fmt.Printf("Unknown flag: %s\n", args[0])
				return
			}
		}
		args = args[1:] // trim after each flag
	}
	// get current working directory
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}
	// start reading inputs
	if len(args) > 0 {
		// loop through all input values
		for _, ch := range args {
			path = ch
			err := myls.ListFiles(path, showHidden, recursive, sortByTime, showDetails, reverse,len(args))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	} else {
		err := myls.ListFiles(path, showHidden, recursive, sortByTime, showDetails, reverse,0)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
