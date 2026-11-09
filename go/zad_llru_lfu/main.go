package main

import (
	"bufio"
	"lab2go/zadachi/zadanie7"
	"os"
)

func main() {
	zadanie7.Run(bufio.NewReader(os.Stdin), true)
}
