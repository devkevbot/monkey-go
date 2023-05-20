package main

import (
	"os"

	"github.com/devkevbot/monkey-go/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
