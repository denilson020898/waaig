package main

import (
	"fmt"
	"os"
	"os/user"
	"waaig/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, MotherFather!\nThis is an unnamed stupid programming language!\n", user.Username)

	fmt.Printf("I insist to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
