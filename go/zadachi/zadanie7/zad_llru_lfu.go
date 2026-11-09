package zadanie7

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/hash"
	"lab2go/struct_go/spisok2"
	"strconv"
	"strings"
)

func LFUSet(cache *hash.HashTable, lists []spisok2.LList, minFrequency *int, capacity int, key, value string) {
	if capacity == 0 {
		return
	}
	saved, found := hash.Get(cache, key)
	if found {
		colon := strings.LastIndexByte(saved, ':')
		frequency, _ := strconv.Atoi(saved[colon+1:])
		spisok2.DeleteValue(&lists[frequency], key)
	} else if hash.Length(cache) == capacity {
		removed := lists[*minFrequency].Head.Chislo
		spisok2.DeleteFirst(&lists[*minFrequency])
		hash.Delete(cache, removed)
	}
	hash.Insert(cache, key, value+":1")
	spisok2.AddEnd(&lists[1], key)
	*minFrequency = 1
}

func LFUGet(cache *hash.HashTable, lists []spisok2.LList, minFrequency *int, key string) string {
	saved, found := hash.Get(cache, key)
	if !found {
		return "-1"
	}
	colon := strings.LastIndexByte(saved, ':')
	value := saved[:colon]
	frequency, _ := strconv.Atoi(saved[colon+1:])
	spisok2.DeleteValue(&lists[frequency], key)
	spisok2.AddEnd(&lists[frequency+1], key)
	if frequency == *minFrequency && lists[frequency].Head == nil {
		*minFrequency = frequency + 1
	}
	hash.Insert(cache, key, value+":"+strconv.Itoa(frequency+1))
	return value
}

func Run(in *bufio.Reader, prompts bool) int {
	var capacity, q int
	if prompts {
		fmt.Print("Ёмкость кэша и количество запросов: ")
	}
	if !input.ReadInts(in, &capacity, &q) || capacity < 0 || q < 0 {
		fmt.Println("Ошибка: нужны неотрицательные числа")
		return 1
	}
	var cache hash.HashTable
	hash.InitTable(&cache, 8)
	lists := make([]spisok2.LList, q+1)
	minFrequency := 0
	var result strings.Builder

	if prompts {
		fmt.Println("Введите запросы SET ключ значение или GET ключ:")
	}
	for i := 0; i < q; i++ {
		var command, key string
		if _, err := fmt.Fscan(in, &command, &key); err != nil {
			fmt.Println("Ошибка: запросов меньше, чем указано")
			return 1
		}
		switch command {
		case "SET":
			var value string
			if _, err := fmt.Fscan(in, &value); err != nil {
				fmt.Println("Ошибка: нужно значение")
				return 1
			}
			LFUSet(&cache, lists, &minFrequency, capacity, key, value)
		case "GET":
			if result.Len() > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(LFUGet(&cache, lists, &minFrequency, key))
		default:
			fmt.Println("Ошибка: неизвестная команда")
			return 1
		}
	}
	if !prompts {
		var extra string
		if _, err := fmt.Fscan(in, &extra); err == nil {
			fmt.Println("Ошибка: лишние данные после запросов")
			return 1
		}
	}
	fmt.Println("Результат: " + result.String())
	return 0
}
