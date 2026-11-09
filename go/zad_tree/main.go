package main

import (
	"bufio"
	"lab2go/zadachi/zadanie5"
	"os"
)

func main() {
	zadanie5.Run(bufio.NewReader(os.Stdin))
}
