package mnozh

import (
	"lab2go/struct_go/hash"
	"lab2go/struct_go/massiv"
)

type Set struct {
	table hash.HashTable
}

func Init(set *Set) {
	hash.InitTable(&set.table, 8)
}

func Add(set *Set, chislo string) {
	hash.Insert(&set.table, chislo, "")
}

func Del(set *Set, chislo string) {
	hash.Delete(&set.table, chislo)
}

func At(set *Set, chislo string) bool {
	return hash.Contains(&set.table, chislo)
}

func Size(set *Set) int {
	return hash.Length(&set.table)
}

func Elements(set *Set, arr *massiv.Array) {
	hash.Keys(&set.table, arr)
}

func Clear(set *Set) {
	hash.Clear(&set.table)
}
