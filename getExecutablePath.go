package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	file, _ := os.Executable()
	var index int
	if runtime.GOOS == "windows" {
		index = strings.LastIndexByte(file, '\\')
	} else {
		index = strings.LastIndexByte(file, '/')
	}
	pathname := file[:index+1]
	fmt.Printf("file: %s\npath: %s\n", file, pathname)
}
