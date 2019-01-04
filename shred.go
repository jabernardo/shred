package main

import (
	"fmt"

	Dumper "github.com/jabernardo/shred/dumper"

	Aargh "github.com/jabernardo/aargh"
)

func main() {
	app := Aargh.New()

	dumper := Dumper.New()
	dumper.Dump("/home/sudoaldrich/Workspace/todo")

	app.Run()

	fmt.Println("Hello")
}
