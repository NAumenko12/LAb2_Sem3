package stack

import "fmt"

type Node struct {
	chislo string
	next   *Node
}

type Stack struct {
	head *Node
}

func InitStack(stack *Stack) {
	stack.head = nil
}
func IsEmptyStack(stack *Stack) bool {
	return stack.head == nil
}
func Push(stack *Stack, chislo string) {
	new_node := &Node{chislo, stack.head}
	stack.head = new_node
}
func Pop(stack *Stack) {
	if IsEmptyStack(stack) {
		fmt.Print("no")
		return
	}
	stack.head = stack.head.next
}
func Top(stack *Stack) string {
	return stack.head.chislo
}
func Print(stack *Stack) {
	for cur_node := stack.head; cur_node != nil; cur_node = cur_node.next {
		fmt.Println(cur_node.chislo)
	}
}
func PrintAddresses(stack *Stack) {
	cur_node := stack.head
	for cur_node != nil {
		fmt.Printf("%s | адрес узла: %p | next: %p\n", cur_node.chislo, cur_node, cur_node.next)
		cur_node = cur_node.next
	}
	fmt.Println()
}
