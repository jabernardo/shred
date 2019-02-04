package main

import (
	"fmt"
	"os"

	Aargh "github.com/jabernardo/aargh"
	Dumper "github.com/jabernardo/shred/dumper"
)

// Shred files
func CommandShredSimple(app *Aargh.App) {
	if len(app.Arguments) == 0 {
		fmt.Println("No input files")
		os.Exit(1)
	}

	dumper := Dumper.New(Dumper.ModeSimple)

	for _, input := range app.Arguments {
		dumper.Dump(input)
	}
}

// Advance mode
func CommandShredAdvance(app *Aargh.App) {
	if len(app.Arguments) == 0 {
		fmt.Println("No input files")
		os.Exit(1)
	}

	dumper := Dumper.New(Dumper.ModeAdvance)

	for _, input := range app.Arguments {
		dumper.Dump(input)
	}
}

// Complex mode
func CommandShredComplex(app *Aargh.App) {
	if len(app.Arguments) == 0 {
		fmt.Println("No input files")
		os.Exit(1)
	}

	dumper := Dumper.New(Dumper.ModeComplex)

	for _, input := range app.Arguments {
		dumper.Dump(input)
	}
}

// Brutal mode
func CommandShredBrutal(app *Aargh.App) {
	if len(app.Arguments) == 0 {
		fmt.Println("No input files")
		os.Exit(1)
	}

	dumper := Dumper.New(Dumper.ModeBrutal)

	for _, input := range app.Arguments {
		dumper.Dump(input)
	}
}

