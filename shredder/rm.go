package shredder

import (
	"fmt"
	"os"
	"path/filepath"
)

type Shredder struct{}

// Create new instance
func New() *Shredder {
	return &Shredder{}
}

// Remove file or directory
// path string
func (this *Shredder) Remove(path string) int8 {
	this.OverrideFiles(this.GetContents(path, true, false))

	return -1 // Nothing to remove
}

// Override files
func (this *Shredder) OverrideFiles(paths []string) {
	dumper := NewDumper()

	for _, path := range paths {
		dumper.OverrideFile(path)
		fmt.Println(path)
	}
}

// Get contents
func (this *Shredder) GetContents(path string, files bool, folders bool) []string {
	var paths []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if (info.IsDir() && folders) || (!info.IsDir() && files) {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}
