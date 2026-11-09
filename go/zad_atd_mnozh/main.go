package main

import (
	"bufio"
	"lab2go/zadachi/zadanie2"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		os.Exit(zadanie2.RunArgs(os.Args[1:]))
	}
	zadanie2.Run(bufio.NewReader(os.Stdin))
}
