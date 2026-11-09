package main

import (
	"bufio"
	"lab2go/zadachi/zadanie3"
	"os"
)

func main() {
	os.Exit(zadanie3.Run(bufio.NewReader(os.Stdin)))
}
