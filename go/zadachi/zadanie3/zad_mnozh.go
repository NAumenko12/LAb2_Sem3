package zadanie3

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/massiv"
	"lab2go/struct_go/mnozh"
	"os"
	"strings"
)

func ReadSetValues(values string, set *mnozh.Set) bool {
	if strings.TrimSpace(values) == "" {
		return true
	}
	for _, part := range strings.Split(values, ",") {
		number := strings.TrimSpace(part)
		if number == "" {
			return false
		}
		for _, c := range number {
			if c < '0' || c > '9' {
				return false
			}
		}
		number = strings.TrimLeft(number, "0")
		if number == "" {
			number = "0"
		}
		mnozh.Add(set, number)
	}
	return true
}

func ReadSets(text string) ([]mnozh.Set, bool) {
	text = strings.TrimSpace(text)
	if len(text) < 2 || text[0] != '{' || text[len(text)-1] != '}' {
		return nil, false
	}
	text = strings.TrimSpace(text[1 : len(text)-1])
	var sets []mnozh.Set
	for text != "" {
		if text[0] != '{' {
			return nil, false
		}
		values, rest, ok := strings.Cut(text[1:], "}")
		if !ok {
			return nil, false
		}
		var set mnozh.Set
		mnozh.Init(&set)
		if !ReadSetValues(values, &set) {
			return nil, false
		}
		sets = append(sets, set)
		text = strings.TrimSpace(rest)
		if text == "" {
			break
		}
		if text[0] != ',' {
			return nil, false
		}
		text = strings.TrimSpace(text[1:])
		if text == "" {
			return nil, false
		}
	}
	return sets, true
}

func CountCommon(first, second *mnozh.Set) int {
	var elements massiv.Array
	massiv.InitArray(&elements, mnozh.Size(first))
	mnozh.Elements(first, &elements)
	common := 0
	for i := 0; i < massiv.Length(&elements); i++ {
		if mnozh.At(second, massiv.Get(&elements, i)) {
			common++
		}
	}
	return common
}

func PrintSetInline(set *mnozh.Set) {
	var elements massiv.Array
	massiv.InitArray(&elements, mnozh.Size(set))
	mnozh.Elements(set, &elements)
	fmt.Print("{")
	for i := 0; i < massiv.Length(&elements); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(massiv.Get(&elements, i))
	}
	fmt.Print("}")
}

func PrintIntersectionPair(first, second *mnozh.Set, common int) {
	fmt.Print("{")
	PrintSetInline(first)
	fmt.Print(", ")
	PrintSetInline(second)
	fmt.Print("}, количество общих элементов: ", common)
	if common > 0 {
		var elements massiv.Array
		massiv.InitArray(&elements, mnozh.Size(first))
		mnozh.Elements(first, &elements)
		fmt.Print(" (")
		printed := 0
		for i := 0; i < massiv.Length(&elements); i++ {
			value := massiv.Get(&elements, i)
			if mnozh.At(second, value) {
				if printed > 0 {
					fmt.Print(", ")
				}
				fmt.Print(value)
				printed++
			}
		}
		fmt.Print(")")
	}
	fmt.Println()
}

func Run(in *bufio.Reader) int {
	fmt.Println("Введите множества в одной строке, например {{1,2,3},{2,3,4}}:")
	text, ok := input.ReadLine(in)
	if !ok {
		return 1
	}
	return RunQuery(text)
}

func RunQuery(text string) int {
	sets, ok := ReadSets(text)
	if !ok {
		fmt.Fprintln(os.Stderr, "Ошибка: нужна запись вида {{1,2,3},{2,3,4}}")
		return 1
	}
	if len(sets) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: нужно хотя бы два множества")
		return 1
	}

	maximum := -1
	for i := 0; i < len(sets); i++ {
		for j := i + 1; j < len(sets); j++ {
			common := CountCommon(&sets[i], &sets[j])
			if common > maximum {
				maximum = common
			}
		}
	}
	for i := 0; i < len(sets); i++ {
		for j := i + 1; j < len(sets); j++ {
			if CountCommon(&sets[i], &sets[j]) == maximum {
				PrintIntersectionPair(&sets[i], &sets[j], maximum)
			}
		}
	}
	return 0
}
