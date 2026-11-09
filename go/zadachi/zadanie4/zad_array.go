package zadanie4

import (
	"bufio"
	"fmt"
	"lab2go/internal/input"
	"lab2go/struct_go/massiv"
	"math"
	"strconv"
	"strings"
)

func MaxPyramid(blocks *massiv.Array) int {
	massiv.Sort(blocks)
	total := 0
	i := 0
	for i < massiv.Length(blocks) {
		width, _, _ := strings.Cut(massiv.Get(blocks, i), " ")
		best := 0
		for i < massiv.Length(blocks) {
			currentWidth, heightText, _ := strings.Cut(massiv.Get(blocks, i), " ")
			if currentWidth != width {
				break
			}
			height, _ := strconv.Atoi(heightText)
			if height > best {
				best = height
			}
			i++
		}
		if best > math.MaxInt32-total {
			return -1
		}
		total += best
	}
	return total
}

func Run(in *bufio.Reader, prompts bool) int {
	var n int
	if prompts {
		fmt.Print("Количество блоков: ")
	}
	if !input.ReadInts(in, &n) || n < 1 || n > 100000 {
		fmt.Println("Ошибка: количество блоков от 1 до 100000")
		return 1
	}
	var blocks massiv.Array
	massiv.InitArray(&blocks, n)
	if prompts {
		fmt.Println("Введите ширину и высоту каждого блока:")
	}
	for i := 0; i < n; i++ {
		var width, height int
		if !input.ReadInts(in, &width, &height) || width < 1 || height < 1 {
			fmt.Println("Ошибка: нужны положительные ширина и высота")
			return 1
		}
		massiv.AddEnd(&blocks, strconv.Itoa(width)+" "+strconv.Itoa(height))
	}
	if !prompts {
		var extra string
		if _, err := fmt.Fscan(in, &extra); err == nil {
			fmt.Println("Ошибка: лишние данные после блоков")
			return 1
		}
	}
	height := MaxPyramid(&blocks)
	if height < 0 {
		fmt.Println("Ошибка: сумма высот не помещается в int")
		return 1
	}
	fmt.Println("Максимальная высота:", height)
	return 0
}
