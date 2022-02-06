package main

import (
	"fmt"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

// TODO
// write code to fully support unicode
