package huffman

import "lab2go/struct_go/massiv"

type Node struct {
	Symbol byte
	Weight int
	Left   *Node
	Right  *Node
}

func IsLeaf(node *Node) bool {
	return node.Left == nil && node.Right == nil
}

func CountFrequencies(text string) [256]int {
	var frequency [256]int
	for i := 0; i < len(text); i++ {
		frequency[text[i]]++
	}
	return frequency
}

func BuildTree(frequency [256]int) *Node {
	var nodes [256]*Node
	count := 0
	for symbol, weight := range frequency {
		if weight > 0 {
			nodes[count] = &Node{Symbol: byte(symbol), Weight: weight}
			count++
		}
	}
	if count == 0 {
		return nil
	}

	for count > 1 {
		min1, min2 := -1, -1
		for i := 0; i < count; i++ {
			if min1 == -1 || nodes[i].Weight < nodes[min1].Weight {
				min2, min1 = min1, i
			} else if min2 == -1 || nodes[i].Weight < nodes[min2].Weight {
				min2 = i
			}
		}
		left, right := nodes[min1], nodes[min2]
		parent := &Node{Weight: left.Weight + right.Weight, Left: left, Right: right}
		if min1 > min2 {
			min1, min2 = min2, min1
		}
		nodes[min1] = parent
		for i := min2; i < count-1; i++ {
			nodes[i] = nodes[i+1]
		}
		count--
	}
	return nodes[0]
}

func BuildCodes(node *Node, code string, codes *massiv.Array) {
	if node == nil {
		return
	}
	if IsLeaf(node) {
		if code == "" {
			code = "0"
		}
		massiv.AddEnd(codes, string([]byte{node.Symbol})+":"+code)
		return
	}
	BuildCodes(node.Left, code+"0", codes)
	BuildCodes(node.Right, code+"1", codes)
}

func Encode(text string, codes *massiv.Array) string {
	bits := make([]byte, 0, len(text))
	for i := 0; i < len(text); i++ {
		for j := 0; j < massiv.Length(codes); j++ {
			entry := massiv.Get(codes, j)
			if entry[0] == text[i] {
				bits = append(bits, entry[2:]...)
				break
			}
		}
	}
	return string(bits)
}

func Decode(root *Node, bits string) (string, bool) {
	if root == nil {
		return "", bits == ""
	}
	text := make([]byte, 0, len(bits))
	if IsLeaf(root) {
		for i := 0; i < len(bits); i++ {
			if bits[i] != '0' {
				return "", false
			}
			text = append(text, root.Symbol)
		}
		return string(text), true
	}
	cur_node := root
	for i := 0; i < len(bits); i++ {
		if bits[i] == '0' {
			cur_node = cur_node.Left
		} else if bits[i] == '1' {
			cur_node = cur_node.Right
		} else {
			return "", false
		}
		if IsLeaf(cur_node) {
			text = append(text, cur_node.Symbol)
			cur_node = root
		}
	}
	return string(text), cur_node == root
}
