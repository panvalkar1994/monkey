package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/panvalkar1994/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err!=nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey Programming Language!\n", user.Username)
	fmt.Println("Feel free to type any commands")
	repl.Start(os.Stdin, os.Stdout)
}