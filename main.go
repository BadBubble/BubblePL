package main

import (
	"BubblePL/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
