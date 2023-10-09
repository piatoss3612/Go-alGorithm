package bj10828

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Stack []int

func (s *Stack) push() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	*s = append(*s, n)
}

func (s *Stack) pop() {
	if len(*s) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	top := len(*s) - 1
	v := (*s)[top]
	*s = (*s)[:top]
	fmt.Fprintln(writer, v)
}

func (s *Stack) size() {
	fmt.Fprintln(writer, len(*s))
}

func (s *Stack) empty() {
	if len(*s) == 0 {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func (s *Stack) top() {
	if len(*s) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, (*s)[len(*s)-1])
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var s Stack
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		switch command {
		case "push":
			s.push()
		case "pop":
			s.pop()
		case "size":
			s.size()
		case "empty":
			s.empty()
		case "top":
			s.top()
		}
	}
}
