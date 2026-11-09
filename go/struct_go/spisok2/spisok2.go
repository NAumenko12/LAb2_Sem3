package spisok2

import "fmt"

type Node struct {
	Chislo string
	Next   *Node
	Prev   *Node
}

type LList struct {
	Head *Node
	Tail *Node
}

func InitList(list *LList) {
	list.Head = nil
	list.Tail = nil
}

func Print(flist *Node) {
	for flist != nil {
		fmt.Println(flist.Chislo)
		flist = flist.Next
	}
}
func PrintBac(last *Node) {
	for last != nil {
		fmt.Println(last.Chislo)
		last = last.Prev
	}
}
func AddAfter(list *LList, flist *Node, chislo string) {
	new_node := &Node{chislo, flist.Next, flist}
	if flist.Next != nil {
		flist.Next.Prev = new_node
	} else {
		list.Tail = new_node
	}
	flist.Next = new_node
}
func AddFirst(list *LList, chislo string) {
	new_node := &Node{chislo, list.Head, nil}
	if list.Head != nil {
		list.Head.Prev = new_node
	} else {
		list.Tail = new_node
	}
	list.Head = new_node
}
func AddEnd(list *LList, chislo string) {
	new_node := &Node{chislo, nil, list.Tail}
	if list.Tail != nil {
		list.Tail.Next = new_node
	} else {
		list.Head = new_node
	}
	list.Tail = new_node
}
func AddBefore(list *LList, flist *Node, chislo string) {
	new_node := &Node{chislo, flist, flist.Prev}
	if flist.Prev != nil {
		flist.Prev.Next = new_node
	} else {
		list.Head = new_node
	}
	flist.Prev = new_node
}
func AddIndex(list *LList, chislo string, index int) {
	if index == 0 {
		AddFirst(list, chislo)
		return
	}
	cur_node := list.Head
	for i := 0; i < index-1; i++ {
		cur_node = cur_node.Next
	}
	AddAfter(list, cur_node, chislo)
}

func DeleteFirst(list *LList) {
	list.Head = list.Head.Next
	if list.Head != nil {
		list.Head.Prev = nil
	} else {
		list.Tail = nil
	}
}
func DeleteEnd(list *LList) {
	list.Tail = list.Tail.Prev
	if list.Tail != nil {
		list.Tail.Next = nil
	} else {
		list.Head = nil
	}
}
func DeleteAfter(list *LList, flist *Node) {
	delete_node := flist.Next
	flist.Next = delete_node.Next
	if delete_node.Next != nil {
		delete_node.Next.Prev = flist
	} else {
		list.Tail = flist
	}
}
func DeleteBefore(list *LList, flist *Node) {
	delete_node := flist.Prev
	flist.Prev = delete_node.Prev
	if delete_node.Prev != nil {
		delete_node.Prev.Next = flist
	} else {
		list.Head = flist
	}
}
func DeleteIndex(list *LList, index int) {
	cur_node := list.Head
	for i := 0; i < index; i++ {
		cur_node = cur_node.Next
	}
	DeleteNode(list, cur_node)
}
func DeleteNode(list *LList, flist *Node) {
	if flist.Prev != nil {
		flist.Prev.Next = flist.Next
	} else {
		list.Head = flist.Next
	}
	if flist.Next != nil {
		flist.Next.Prev = flist.Prev
	} else {
		list.Tail = flist.Prev
	}
}
func DeleteValue(list *LList, chislo string) {
	cur_node := list.Head
	for cur_node != nil && cur_node.Chislo != chislo {
		cur_node = cur_node.Next
	}
	if cur_node == nil {
		return
	}
	DeleteNode(list, cur_node)
}
func Find(flist *Node, chislo string) {
	for flist != nil {
		if flist.Chislo == chislo {
			fmt.Printf("%s найден, адрес: %p\n", chislo, flist)
			return
		}
		flist = flist.Next
	}
	fmt.Println(chislo, "не найден")
}
func PrintAddresses(list *LList) {
	fmt.Printf("head: %p | tail: %p\n", list.Head, list.Tail)
	cur_node := list.Head
	for cur_node != nil {
		fmt.Printf("%s | адрес: %p | next: %p | prev: %p\n", cur_node.Chislo, cur_node, cur_node.Next, cur_node.Prev)
		cur_node = cur_node.Next
	}
	fmt.Println()
}
