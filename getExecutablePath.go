package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Executable()
	index := strings.LastIndexByte(file, '\\')
	pathname := file[:index+1]
	fmt.Printf("file: %s\npath: %s\n", file, pathname)
}
