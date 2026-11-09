package main

import (
	"bufio"
	"lab2go/zadachi/zadanie4"
	"os"
)

func main() {
	zadanie4.Run(bufio.NewReader(os.Stdin), true)
}
