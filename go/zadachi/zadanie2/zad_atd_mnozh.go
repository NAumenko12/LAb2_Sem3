package zadanie2

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/massiv"
	"lab2go/struct_go/mnozh"
	"os"
	"strings"
)

func LoadSet(set *mnozh.Set, filename string) error {
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return nil
	}

	var values []string
	if trimmed[0] == '[' || strings.HasSuffix(filename, ".json") {
		if trimmed[0] != '[' {
			return fmt.Errorf("ожидался JSON-массив строк")
		}
		if err := json.Unmarshal(trimmed, &values); err != nil {
			return err
		}
	} else {
		for _, line := range strings.Split(string(data), "\n") {
			value := strings.TrimSuffix(line, "\r")
			if value != "" {
				values = append(values, value)
			}
		}
	}
	for _, value := range values {
		mnozh.Add(set, value)
	}
	return nil
}

func SaveSet(set *mnozh.Set, filename string) error {
	var elements massiv.Array
	massiv.InitArray(&elements, mnozh.Size(set))
	mnozh.Elements(set, &elements)

	values := make([]string, massiv.Length(&elements))
	for i := range values {
		values[i] = massiv.Get(&elements, i)
	}

	var data bytes.Buffer
	encoder := json.NewEncoder(&data)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(values); err != nil {
		return err
	}
	return os.WriteFile(filename, data.Bytes(), 0644)
}

func RunSetQuery(filename, query string) int {
	fields := strings.Fields(query)
	command := ""
	if len(fields) > 0 {
		command = fields[0]
	}
	switch command {
	case "SETADD", "SETDEL", "SET_AT", "SETPRINT":
	default:
		fmt.Fprintln(os.Stderr, "Ошибка: неизвестная команда "+command)
		return 1
	}

	expected := 2
	if command == "SETPRINT" {
		expected = 1
	}
	if len(fields) < expected {
		fmt.Fprintln(os.Stderr, "Ошибка: нужно значение элемента")
		return 1
	}
	if len(fields) > expected {
		fmt.Fprintln(os.Stderr, "Ошибка: слишком много аргументов команды")
		return 1
	}
	value := ""
	if expected == 2 {
		value = fields[1]
		for i := 0; i < len(value); i++ {
			if value[i] < ' ' {
				fmt.Fprintln(os.Stderr, "Ошибка: управляющий символ в элементе")
				return 1
			}
		}
	}

	var set mnozh.Set
	mnozh.Init(&set)
	if err := LoadSet(&set, filename); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: не удалось прочитать файл", filename, err)
		return 1
	}
	sizeBefore := mnozh.Size(&set)

	switch command {
	case "SETADD":
		mnozh.Add(&set, value)
		if mnozh.Size(&set) == sizeBefore {
			fmt.Println(value, "уже есть в множестве")
		}
	case "SETDEL":
		mnozh.Del(&set, value)
		if mnozh.Size(&set) == sizeBefore {
			fmt.Println(value, "нет в множестве")
		}
	case "SET_AT":
		if mnozh.At(&set, value) {
			fmt.Println("-> TRUE")
		} else {
			fmt.Println("-> FALSE")
		}
	case "SETPRINT":
		var elements massiv.Array
		massiv.InitArray(&elements, mnozh.Size(&set))
		mnozh.Elements(&set, &elements)
		fmt.Print("{")
		for i := 0; i < massiv.Length(&elements); i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(massiv.Get(&elements, i))
		}
		fmt.Println("}")
	}

	if mnozh.Size(&set) != sizeBefore {
		if err := SaveSet(&set, filename); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка: не удалось записать файл", filename)
			return 1
		}
		fmt.Print("-> ", value)
		if command == "SETDEL" {
			fmt.Print(" удалён")
		}
		fmt.Println()
	}
	return 0
}

func RunArgs(args []string) int {
	filename, query := "", ""
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if (arg != "--file" && arg != "--query") || i+1 >= len(args) {
			fmt.Fprintln(os.Stderr, "Ошибка: неизвестный или неполный параметр", arg)
			return 1
		}
		i++
		if arg == "--file" {
			filename = args[i]
		} else {
			query = args[i]
		}
	}
	if filename == "" || query == "" {
		fmt.Fprintln(os.Stderr, "Запуск: lab2 --file set.json --query 'SETADD x | SETDEL x | SET_AT x | SETPRINT'")
		return 1
	}
	return RunSetQuery(filename, query)
}

func Run(in *bufio.Reader) {
	fmt.Print("Файл множества: ")
	filename, _ := input.ReadLine(in)
	fmt.Print("Команда (SETADD x / SETDEL x / SET_AT x / SETPRINT): ")
	query, _ := input.ReadLine(in)
	RunSetQuery(filename, query)
}
