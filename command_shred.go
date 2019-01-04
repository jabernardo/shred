package main

import (
	"fmt"
	"os"

	Aargh "github.com/jabernardo/aargh"
	Dumper "github.com/jabernardo/shred/dumper"
)

// Simple shred
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
