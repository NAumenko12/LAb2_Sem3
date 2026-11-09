package main

import (
	"bufio"
	"lab2go/zadachi/zadanie1"
	"os"
)

func main() {
	zadanie1.Run(bufio.NewReader(os.Stdin))
}
