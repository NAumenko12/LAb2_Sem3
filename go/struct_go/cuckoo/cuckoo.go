package cuckoo

import "fmt"

type CuckooCell struct {
	Key  string
	Used bool
}

type CuckooTable struct {
	First  [11]CuckooCell
	Second [11]CuckooCell
}

func Hash1(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash*31 + int(key[i])) % 11
	}
	return hash
}

func Hash2(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		value := int(key[i])
		hash = (hash*17 + value + value/11) % 11
	}
	return hash
}

func Contains(table *CuckooTable, key string) bool {
	first := table.First[Hash1(key)]
	second := table.Second[Hash2(key)]
	return (first.Used && first.Key == key) || (second.Used && second.Key == key)
}

func Insert(table *CuckooTable, key string) bool {
	if Contains(table, key) {
		return true
	}
	backup := *table
	current := key
	for i := 0; i < 20; i++ {
		first := Hash1(current)
		if !table.First[first].Used {
			table.First[first] = CuckooCell{Key: current, Used: true}
			return true
		}
		displaced := table.First[first].Key
		table.First[first].Key = current
		current = displaced

		second := Hash2(current)
		if !table.Second[second].Used {
			table.Second[second] = CuckooCell{Key: current, Used: true}
			return true
		}
		displaced = table.Second[second].Key
		table.Second[second].Key = current
		current = displaced
	}
	*table = backup
	return false
}

func Delete(table *CuckooTable, key string) {
	first := Hash1(key)
	second := Hash2(key)
	if table.First[first].Used && table.First[first].Key == key {
		table.First[first].Used = false
	}
	if table.Second[second].Used && table.Second[second].Key == key {
		table.Second[second].Used = false
	}
}

func Print(table *CuckooTable) {
	fmt.Println("Индекс | первая таблица | вторая таблица")
	for i := 0; i < len(table.First); i++ {
		first, second := "-", "-"
		if table.First[i].Used {
			first = table.First[i].Key
		}
		if table.Second[i].Used {
			second = table.Second[i].Key
		}
		fmt.Printf("%d | %s | %s\n", i, first, second)
	}
}
