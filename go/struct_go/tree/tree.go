package tree

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func InitTree(tree *Tree) {
	tree.Root = nil
}

func Add(root **Node, chislo int) {
	if *root == nil {
		*root = &Node{Data: chislo}
		return
	}
	if chislo < (*root).Data {
		Add(&(*root).Left, chislo)
	} else if chislo > (*root).Data {
		Add(&(*root).Right, chislo)
	}
}

func Delete(root **Node, chislo int) {
	if *root == nil {
		return
	}
	if chislo < (*root).Data {
		Delete(&(*root).Left, chislo)
		return
	}
	if chislo > (*root).Data {
		Delete(&(*root).Right, chislo)
		return
	}
	if (*root).Left != nil && (*root).Right != nil {
		rep := (*root).Right
		for rep.Left != nil {
			rep = rep.Left
		}
		(*root).Data = rep.Data
		Delete(&(*root).Right, rep.Data)
		return
	}
	if (*root).Left != nil {
		*root = (*root).Left
	} else {
		*root = (*root).Right
	}
}

func Find(root *Node, chislo int) *Node {
	for root != nil && root.Data != chislo {
		if chislo < root.Data {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func Height(root *Node) int {
	if root == nil {
		return 0
	}
	left := Height(root.Left)
	right := Height(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func PrintPryam(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, " ")
	PrintPryam(root.Left)
	PrintPryam(root.Right)
}
func PrintSimetrich(root *Node) {
	if root == nil {
		return
	}
	PrintSimetrich(root.Left)
	fmt.Print(root.Data, " ")
	PrintSimetrich(root.Right)
}
func PrintObrat(root *Node) {
	if root == nil {
		return
	}
	PrintObrat(root.Left)
	PrintObrat(root.Right)
	fmt.Print(root.Data, " ")
}

func PrintLevel(root *Node, level int) bool {
	if root == nil {
		return false
	}
	if level == 0 {
		fmt.Print(root.Data, " ")
		return true
	}
	left := PrintLevel(root.Left, level-1)
	right := PrintLevel(root.Right, level-1)
	return left || right
}

func PrintBFS(root *Node) {
	level := 0
	for PrintLevel(root, level) {
		level++
	}
}

func PrintAddresses(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("Значение: %d | адрес: %p | left: %p | right: %p\n", root.Data, root, root.Left, root.Right)
	PrintAddresses(root.Left)
	PrintAddresses(root.Right)
}
