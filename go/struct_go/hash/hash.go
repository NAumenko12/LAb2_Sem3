package hash

import (
	"fmt"
	"lab2go/struct_go/massiv"
)

type HNode struct {
	key   string
	value string
	next  *HNode
}

type HashTable struct {
	data  []*HNode
	size  int
	mesto int
}

func InitTable(table *HashTable, mesto int) {
	if mesto < 1 {
		mesto = 1
	}
	table.data = make([]*HNode, mesto)
	table.size = 0
	table.mesto = mesto
}

func Hash(key string, mesto int) int {
	var h uint32 = 0
	for i := 0; i < len(key); i++ {
		h = h*31 + uint32(key[i])
	}
	return int(h % uint32(mesto))
}

func Resize(table *HashTable) {
	new_mesto := table.mesto * 2
	new_data := make([]*HNode, new_mesto)
	for i := 0; i < table.mesto; i++ {
		cur_node := table.data[i]
		for cur_node != nil {
			next_node := cur_node.next
			index := Hash(cur_node.key, new_mesto)
			cur_node.next = new_data[index]
			new_data[index] = cur_node
			cur_node = next_node
		}
	}
	table.data = new_data
	table.mesto = new_mesto
}

func Insert(table *HashTable, key string, value string) {
	if table.data == nil {
		InitTable(table, 8)
	}
	index := Hash(key, table.mesto)
	for cur_node := table.data[index]; cur_node != nil; cur_node = cur_node.next {
		if cur_node.key == key {
			cur_node.value = value
			return
		}
	}
	table.data[index] = &HNode{key, value, table.data[index]}
	table.size++
	if table.size >= table.mesto-table.mesto/4 {
		Resize(table)
	}
}

func Get(table *HashTable, key string) (string, bool) {
	if table.data == nil {
		return "", false
	}
	for cur_node := table.data[Hash(key, table.mesto)]; cur_node != nil; cur_node = cur_node.next {
		if cur_node.key == key {
			return cur_node.value, true
		}
	}
	return "", false
}

func Contains(table *HashTable, key string) bool {
	_, ok := Get(table, key)
	return ok
}

func Delete(table *HashTable, key string) {
	if table.data == nil {
		return
	}
	index := Hash(key, table.mesto)
	cur_node := table.data[index]
	var prev_node *HNode = nil
	for cur_node != nil && cur_node.key != key {
		prev_node = cur_node
		cur_node = cur_node.next
	}
	if cur_node == nil {
		return
	}
	if prev_node == nil {
		table.data[index] = cur_node.next
	} else {
		prev_node.next = cur_node.next
	}
	table.size--
}

func Length(table *HashTable) int {
	return table.size
}

func Mesto(table *HashTable) int {
	return table.mesto
}

func Keys(table *HashTable, arr *massiv.Array) {
	for i := 0; i < table.mesto; i++ {
		for cur_node := table.data[i]; cur_node != nil; cur_node = cur_node.next {
			massiv.AddEnd(arr, cur_node.key)
		}
	}
}

func ChainStats(table *HashTable) (int, float64) {
	max_len, not_empty := 0, 0
	for i := 0; i < table.mesto; i++ {
		length := 0
		for cur_node := table.data[i]; cur_node != nil; cur_node = cur_node.next {
			length++
		}
		if length > 0 {
			not_empty++
		}
		if length > max_len {
			max_len = length
		}
	}
	if not_empty == 0 {
		return 0, 0
	}
	return max_len, float64(table.size) / float64(not_empty)
}

func Print(table *HashTable) {
	for i := 0; i < table.mesto; i++ {
		for cur_node := table.data[i]; cur_node != nil; cur_node = cur_node.next {
			fmt.Println(cur_node.key, "->", cur_node.value)
		}
	}
}

func PrintAddresses(table *HashTable) {
	for i := 0; i < table.mesto; i++ {
		fmt.Printf("[%d]", i)
		for cur_node := table.data[i]; cur_node != nil; cur_node = cur_node.next {
			fmt.Printf(" -> %s:%s (%p)", cur_node.key, cur_node.value, cur_node)
		}
		fmt.Println()
	}
	fmt.Printf("size = %d mesto = %d\n", table.size, table.mesto)
}

func Clear(table *HashTable) {
	for i := 0; i < table.mesto; i++ {
		table.data[i] = nil
	}
	table.size = 0
}
