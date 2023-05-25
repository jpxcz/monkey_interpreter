package main

import (
	"fmt"
	relp "monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey pporgaming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	relp.Start(os.Stdin, os.Stdout)
}