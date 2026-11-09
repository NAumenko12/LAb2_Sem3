package zadanie6

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/cuckoo"
	"lab2go/struct_go/hash"
	"lab2go/struct_go/massiv"
	"strconv"
	"strings"
)

func SplitSymbols(text string, symbols *massiv.Array) {
	for _, symbol := range text {
		massiv.AddEnd(symbols, string(symbol))
	}
}

func LongestSubstring(symbols *massiv.Array) (int, int) {
	var positions hash.HashTable
	hash.InitTable(&positions, 8)
	left, bestLeft, maximum := 0, 0, 0
	for right := 0; right < massiv.Length(symbols); right++ {
		symbol := massiv.Get(symbols, right)
		if saved, found := hash.Get(&positions, symbol); found {
			previous, _ := strconv.Atoi(saved)
			if previous >= left {
				left = previous + 1
			}
		}
		hash.Insert(&positions, symbol, strconv.Itoa(right))
		length := right - left + 1
		if length > maximum {
			maximum = length
			bestLeft = left
		}
	}
	return bestLeft, maximum
}

func RunLongestSubstring(text string) {
	var symbols massiv.Array
	massiv.InitArray(&symbols, len(text))
	SplitSymbols(text, &symbols)
	start, length := LongestSubstring(&symbols)
	var substring strings.Builder
	for i := start; i < start+length; i++ {
		substring.WriteString(massiv.Get(&symbols, i))
	}
	fmt.Printf("Цепочки: %d (%s)\n", length, substring.String())
}

func DemoHashTables(line string) {

	var chain hash.HashTable
	var table cuckoo.CuckooTable
	hash.InitTable(&chain, 8)
	for _, key := range strings.Fields(line) {
		hash.Insert(&chain, key, "value")
		if !cuckoo.Insert(&table, key) {
			fmt.Println("Кукушка: цикл при добавлении ключа", key)
		}
	}
	fmt.Println("Цепочки:")
	hash.PrintAddresses(&chain)
	fmt.Println("Кукушка:")
	cuckoo.Print(&table)
}

func Run(in *bufio.Reader) {
	var choice int
	fmt.Print("1 — подстрока без повторов, 2 — демонстрация таблиц: ")
	if !input.ReadInts(in, &choice) {
		return
	}
	input.ReadLine(in)
	switch choice {
	case 1:
		fmt.Println("Введите строку:")
		text, _ := input.ReadLine(in)
		RunLongestSubstring(text)
	case 2:
		fmt.Println("Введите ключи через пробел:")
		line, _ := input.ReadLine(in)
		DemoHashTables(line)
	default:
		fmt.Println("Нет такого режима")
	}
}
