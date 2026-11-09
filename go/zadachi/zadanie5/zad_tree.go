package zadanie5

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/huffman"
	"lab2go/struct_go/massiv"
)

func PrintHuffmanTable(codes *massiv.Array, frequency [256]int) {
	fmt.Println("Символ — частота — код:")
	for i := 0; i < massiv.Length(codes); i++ {
		entry := massiv.Get(codes, i)
		symbol := entry[0]
		if symbol == ' ' {
			fmt.Print("пробел")
		} else if symbol > 32 && symbol < 127 {
			fmt.Printf("%c", symbol)
		} else {
			fmt.Print("байт ", int(symbol))
		}
		fmt.Printf(" — %d — %s\n", frequency[symbol], entry[2:])
	}
}

func Run(in *bufio.Reader) {
	fmt.Println("Введите строку:")
	text, _ := input.ReadLine(in)
	run(text, "", in)
}

func RunQuery(text, decodeBits string) int {
	return run(text, decodeBits, nil)
}

func run(text, decodeBits string, in *bufio.Reader) int {
	if text == "" {
		fmt.Println("Строка пустая")
		return 1
	}

	frequency := huffman.CountFrequencies(text)
	root := huffman.BuildTree(frequency)
	var codes massiv.Array
	massiv.InitArray(&codes, 256)
	huffman.BuildCodes(root, "", &codes)
	PrintHuffmanTable(&codes, frequency)

	bits := huffman.Encode(text, &codes)
	decoded, _ := huffman.Decode(root, bits)
	fmt.Println("Код:", bits)
	fmt.Println("Расшифровка:", decoded)
	if in != nil {
		fmt.Println("Введите код для расшифровки по этому дереву (пустая строка — пропустить):")
		decodeBits, _ = input.ReadLine(in)
	}
	if decodeBits != "" {
		decoded, ok := huffman.Decode(root, decodeBits)
		if !ok {
			fmt.Println("Ошибка: код содержит не 0/1 или обрывается внутри символа")
			return 1
		}
		fmt.Println("Результат:", decoded)
	}
	return 0
}
