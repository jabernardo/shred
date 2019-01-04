package main

import (
	Aargh "github.com/jabernardo/aargh"
)

func main() {
	app := Aargh.New()

	// Set help screen
	app.Command("help", CommandHelp)

	// Set default command
	app.Command("help", CommandHelp)

	// Modes
	app.Command("simple", CommandShredSimple)

	app.Run()
}
