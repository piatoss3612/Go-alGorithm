package bj18258

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue []int

func (q *Queue) push() {
	n := scanInt()
	*q = append(*q, n)
}

func (q *Queue) pop() {
	if len(*q) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	v := (*q)[0]
	*q = (*q)[1:]
	fmt.Fprintln(writer, v)
}

func (q *Queue) size() {
	fmt.Fprintln(writer, len(*q))
}

func (q *Queue) empty() {
	if len(*q) == 0 {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func (q *Queue) front() {
	if len(*q) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, (*q)[0])
	}
}

func (q *Queue) back() {
	if len(*q) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, (*q)[len(*q)-1])
	}
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	var q Queue
	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		switch cmd {
		case "push":
			q.push()
		case "pop":
			q.pop()
		case "size":
			q.size()
		case "empty":
			q.empty()
		case "front":
			q.front()
		case "back":
			q.back()
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
