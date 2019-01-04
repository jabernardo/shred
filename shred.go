package main

import (
	"fmt"

	Dumper "github.com/jabernardo/shred/dumper"

	Aargh "github.com/jabernardo/aargh"
)

func main() {
	app := Aargh.New()

	dumper := Dumper.New(Dumper.ModeBrutal)

	dumper.Dump("/home/sudoaldrich/Workspace/hello.rs")

	app.Run()

	fmt.Println("Hello")
}
