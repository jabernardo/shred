package main

import (
	"fmt"
	"os"

	Aargh "github.com/jabernardo/aargh"
)

// Temporary Help screen
func CommandHelp(app *Aargh.App) {
	fmt.Printf("File shredding utility @%s\n\n", app.Version)

	fmt.Printf("Usage:\n\t %s method [files]\n\n", os.Args[0])

	fmt.Printf("Methods:\n\n")
	fmt.Printf("\t%-10v\t1 Pass\n", "simple")
	fmt.Printf("\t%-10v\t3 Passes\n", "advance")
	fmt.Printf("\t%-10v\t7 Passes\n", "complex")
	fmt.Printf("\t%-10v\t35 Passes\n", "brutal")
}
