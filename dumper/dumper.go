package dumper

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Dumper struct {
	Level int
}

const (
	ModeSimple  int = 1
	ModeAdvance     = 3
	ModeComplex     = 7
	ModeBrutal      = 35
)

// New Dumper
func New(level int) *Dumper {
	new := &Dumper{}
	new.Level = level

	return new
}

// Generate some garbage
func (this *Dumper) Random(length uint64) string {
	var chars string = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seeded *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	garbage := make([]byte, length)

	for i := range garbage {
		garbage[i] = chars[seeded.Intn(len(chars))]
	}

	return string(garbage)
}

// Override file contents
func (this *Dumper) OverrideFile(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	mode := info.Mode()
	size := uint64(info.Size())

	if mode.IsRegular() {
		content := []byte(this.Random(size))
		err := ioutil.WriteFile(path, content, mode)

		if err != nil {
			return false
		}
	}

	return true
}

// Get contents
func (this *Dumper) GetContents(path string, files_only bool) []string {
	var paths []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && files_only {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}

// Dump path
func (this *Dumper) Dump(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		fmt.Println("No path: ", path)
		return false
	}

	if info.IsDir() {
		for _, file := range this.GetContents(path, true) {
			for i := 0; i < this.Level; i++ {
				this.OverrideFile(file)
			}

			fmt.Println(file)
		}
	} else {
		this.OverrideFile(path)
		fmt.Println(path)
	}

	name_length := len(info.Name())
	name_size := uint64(name_length)
	temp_name := filepath.Dir(path) + string(os.PathSeparator) + this.Random(name_size)

	// Rename path
	err = os.Rename(path, temp_name)

	if err != nil {
		return false
	}

	// Remove file

	err = os.RemoveAll(temp_name)

	if err != nil {
		return false
	}

	return true
}
