package zadanie1

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/massiv"
	"lab2go/struct_go/stack"
	"lab2go/struct_go/tree"
	"strconv"
)

func TreeDFS(root *tree.Node) {
	if root == nil {
		fmt.Println("дерево пустое")
		return
	}
	var st stack.Stack
	stack.InitStack(&st)
	stack.Push(&st, strconv.Itoa(root.Data))
	for !stack.IsEmptyStack(&st) {
		value, _ := strconv.Atoi(stack.Top(&st))
		stack.Pop(&st)
		node := tree.Find(root, value)
		fmt.Print(node.Data, " ")
		if node.Right != nil {
			stack.Push(&st, strconv.Itoa(node.Right.Data))
		}
		if node.Left != nil {
			stack.Push(&st, strconv.Itoa(node.Left.Data))
		}
	}
}

func RunTreeDFS(in *bufio.Reader, prompts bool) int {
	var t tree.Tree
	tree.InitTree(&t)
	if prompts {
		fmt.Println("Введите числа дерева, 0 — конец:")
	}
	for {
		var token string
		if _, err := fmt.Fscan(in, &token); err != nil {
			break
		}
		value, err := strconv.ParseInt(token, 10, 32)
		if err != nil {
			if prompts {
				break
			}
			fmt.Println("Ошибка: нужны целые числа для дерева")
			return 1
		}
		if value == 0 {
			break
		}
		tree.Add(&t.Root, int(value))
	}
	fmt.Print("Обход в глубину  ")
	TreeDFS(t.Root)
	fmt.Println()
	return 0
}

func GraphDFS(neighbors []massiv.Array, start int) {
	visited := make([]bool, len(neighbors))
	var st stack.Stack
	stack.InitStack(&st)
	stack.Push(&st, strconv.Itoa(start))
	for !stack.IsEmptyStack(&st) {
		vertex, _ := strconv.Atoi(stack.Top(&st))
		stack.Pop(&st)
		if visited[vertex] {
			continue
		}
		visited[vertex] = true
		fmt.Print(vertex, " ")
		for i := 0; i < massiv.Length(&neighbors[vertex]); i++ {
			stack.Push(&st, massiv.Get(&neighbors[vertex], i))
		}
	}
	fmt.Println()
}

func RunGraphDFS(in *bufio.Reader, prompts bool) int {
	var n, m int
	if prompts {
		fmt.Print("Введите количество вершин и рёбер: ")
	}
	if !input.ReadInts(in, &n, &m) || n < 1 || m < 0 {
		fmt.Println("Ошибка: неверное количество вершин или рёбер")
		return 1
	}
	neighbors := make([]massiv.Array, n+1)
	for i := range neighbors {
		massiv.InitArray(&neighbors[i], 2)
	}
	if prompts {
		fmt.Println("Введите рёбра парами u v:")
	}
	for i := 0; i < m; i++ {
		var u, v int
		if !input.ReadInts(in, &u, &v) || u < 1 || u > n || v < 1 || v > n {
			fmt.Println("Ошибка: неверное ребро")
			return 1
		}
		massiv.AddEnd(&neighbors[u], strconv.Itoa(v))
		massiv.AddEnd(&neighbors[v], strconv.Itoa(u))
	}
	if prompts {
		fmt.Print("Введите начальную вершину: ")
	}
	var start int
	if !input.ReadInts(in, &start) || start < 1 || start > n {
		if prompts {
			fmt.Println("Ошибка: неверная начальная вершина")
		} else {
			fmt.Println("Ошибка: неверная начальная вершина или лишние данные")
		}
		return 1
	}
	if !prompts {
		var extra string
		if _, err := fmt.Fscan(in, &extra); err == nil {
			fmt.Println("Ошибка: неверная начальная вершина или лишние данные")
			return 1
		}
	}
	fmt.Print("Обход в глубину: ")
	GraphDFS(neighbors, start)
	return 0
}

func Run(in *bufio.Reader) {
	var choice int
	fmt.Print("1 — дерево, 2 — граф: ")
	if !input.ReadInts(in, &choice) {
		return
	}
	switch choice {
	case 1:
		RunTreeDFS(in, true)
	case 2:
		RunGraphDFS(in, true)
	default:
		fmt.Println("Нет такого режима")
	}
}
