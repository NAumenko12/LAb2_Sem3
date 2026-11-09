package main

import (
	"bufio"
	"fmt"
	"lab2go/zadachi/zadanie1"
	"lab2go/zadachi/zadanie2"
	"lab2go/zadachi/zadanie3"
	"lab2go/zadachi/zadanie4"
	"lab2go/zadachi/zadanie5"
	"lab2go/zadachi/zadanie6"
	"lab2go/zadachi/zadanie7"
	"os"
	"strings"
)

func PrintHelp() {
	fmt.Print("Запуск: lab2 --query 'КОМАНДА данные'\n" +
		"DFS TREE 7 3 9 1 5\n" +
		"DFS GRAPH 3 2 1 2 2 3 1\n" +
		"INTER {{1,2,3},{2,3,4}}\n" +
		"PYRAMID 3 3 1 2 2 3 3\n" +
		"HUFFMAN abracadabra\n" +
		"SUBSTRING abcabcbb\n" +
		"CUCKOO a l w\n" +
		"LFU 2 4 SET 1 10 SET 2 20 GET 1 GET 3\n" +
		"lab2 --file set.json --query 'SETADD 5'\n" +
		"lab2 --query 'HUFFMAN aaaa' --bits 0000\n")
}

func RunQuery(query, filename, bits string) int {
	query = strings.TrimLeft(query, " \t\r\n\v\f")
	command, data := query, ""
	if i := strings.IndexAny(query, " \t\r\n\v\f"); i >= 0 {
		command, data = query[:i], query[i+1:]
	}
	switch command {
	case "":
		fmt.Fprintln(os.Stderr, "Ошибка: пустой запрос")
		return 1
	case "SETADD", "SETDEL", "SET_AT", "SETPRINT":
		if filename == "" {
			fmt.Fprintln(os.Stderr, "Ошибка: укажи --file для множества")
			return 1
		}
		return zadanie2.RunSetQuery(filename, query)
	case "DFS":
		in := bufio.NewReader(strings.NewReader(data))
		var kind string
		fmt.Fscan(in, &kind)
		switch kind {
		case "TREE":
			return zadanie1.RunTreeDFS(in, false)
		case "GRAPH":
			return zadanie1.RunGraphDFS(in, false)
		default:
			fmt.Fprintln(os.Stderr, "Ошибка: укажи DFS TREE или DFS GRAPH")
			return 1
		}
	case "INTER":
		return zadanie3.RunQuery(data)
	case "PYRAMID":
		return zadanie4.Run(bufio.NewReader(strings.NewReader(data)), false)
	case "HUFFMAN":
		return zadanie5.RunQuery(data, bits)
	case "SUBSTRING":
		zadanie6.RunLongestSubstring(data)
	case "CUCKOO":
		if data == "" {
			data = "apple banana cherry"
		}
		zadanie6.DemoHashTables(data)
	case "LFU":
		return zadanie7.Run(bufio.NewReader(strings.NewReader(data)), false)
	default:
		fmt.Fprintln(os.Stderr, "Ошибка: неизвестная команда "+command)
		return 1
	}
	return 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || (len(args) == 1 && args[0] == "--help") {
		PrintHelp()
		return
	}

	query, filename, bits := "", "", ""
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if i+1 >= len(args) || (arg != "--query" && arg != "--file" && arg != "--bits") {
			fmt.Fprintln(os.Stderr, "Ошибка: неизвестный или неполный параметр", arg)
			os.Exit(1)
		}
		i++
		switch arg {
		case "--query":
			query = args[i]
		case "--file":
			filename = args[i]
		case "--bits":
			bits = args[i]
		}
	}
	os.Exit(RunQuery(query, filename, bits))
}
