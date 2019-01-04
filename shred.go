package main

import (
	"fmt"

	shredder "github.com/jabernardo/shred/shredder"

	aargh "github.com/jabernardo/aargh"
)

func main() {
	app := aargh.New()
	shred := shredder.New()

	shred.Remove("/home/sudoaldrich/Workspace/todo")

	app.Run()

	fmt.Println("Hello")
}
