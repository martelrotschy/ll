package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	dirPath := "."
	if len(os.Args) > 1 {
		dirPath = os.Args[1]
	}

	// Add current directory (.) and parent directory (..)
	files := []os.DirEntry{
		&dirEntry{name: ".", isDir: true},
		&dirEntry{name: "..", isDir: true},
	}

	// Read the directory entries
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Append the directory entries to the files list
	files = append(files, dirEntries...)

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		// File name
		fmt.Printf("%s", file.Name())

		// File size
		fmt.Printf("  %6d", fileInfo.Size())

		// File permissions
		mode := fileInfo.Mode()
		fmt.Printf("  %s", mode)

		// File modification time
		modTime := fileInfo.ModTime()
		fmt.Printf("  %s", modTime.Format("Jan _2 15:04"))

		// File type (directory or file)
		if mode.IsDir() {
			fmt.Printf("  <DIR>\n")
		} else {
			fmt.Println()
		}
	}
}

type dirEntry struct {
	name  string
	isDir bool
}

func (d *dirEntry) Name() string {
	return d.name
}

func (d *dirEntry) IsDir() bool {
	return d.isDir
}

func (d *dirEntry) Type() os.FileMode {
	if d.isDir {
		return os.ModeDir
	}
	return 0
}

func (d *dirEntry) Info() (os.FileInfo, error) {
	return d, nil
}

func (d *dirEntry) Size() int64 {
	return 0
}

func (d *dirEntry) Mode() os.FileMode {
	return os.ModeDir | 0777
}

func (d *dirEntry) ModTime() time.Time {
	return time.Time{}
}

func (d *dirEntry) Sys() interface{} {
	return nil
}