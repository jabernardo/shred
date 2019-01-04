package shredder

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Dumper struct{}

// New Dumper
func NewDumper() *Dumper {
	return &Dumper{}
}

// Generate some garbage
func (this *Dumper) GetGarbage(length int64) string {
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
	size := info.Size()

	if mode.IsRegular() {
		content := []byte(this.GetGarbage(size))
		err := ioutil.WriteFile(path, content, mode)

		if err != nil {
			fmt.Printf("Can't write to file: %s\n", path)
		}
	}

	return true
}
