package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadLine(in *bufio.Reader) (string, bool) {
	line, err := in.ReadString('\n')
	return strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), err == nil || len(line) > 0
}

func ReadInts(in *bufio.Reader, numbers ...*int) bool {
	for _, number := range numbers {
		var text string
		if _, err := fmt.Fscan(in, &text); err != nil {
			return false
		}
		value, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			return false
		}
		*number = int(value)
	}
	return true
}
