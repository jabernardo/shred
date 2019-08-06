package main

import (
	Aargh "github.com/jabernardo/aargh"
)

func main() {
	app := Aargh.New()

	app.Name = "shred"
	app.Version = "0.1"

	// Set help screen
	app.Command("help", CommandHelp)

	// Set default command
	app.Command("default", CommandHelp)

	// Modes
	app.Command("simple", CommandShredSimple)
	app.Command("advance", CommandShredAdvance)
	app.Command("complex", CommandShredComplex)
	app.Command("brutal", CommandShredBrutal)

	app.Run(true)
}
