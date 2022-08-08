package main

import (
	"fmt"
	"interpreter-go/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", usr.Username)
	fmt.Println("Welcome to Monkey REPL")
	repl.Start(os.Stdin, os.Stdout)
}
