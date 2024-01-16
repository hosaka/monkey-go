package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hosaka/monkey-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to Monkey\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
