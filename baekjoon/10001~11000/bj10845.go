package bj10845

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Queue []int

func (q *Queue) push() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var q Queue
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		switch command {
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
