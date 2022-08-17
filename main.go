package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	var files []string

	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jpg") {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	for idx, file := range files {
		os.Rename(file, strconv.Itoa(idx+1)+".jpg")
	}
}
