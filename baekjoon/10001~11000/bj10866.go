package bj10866

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

type Deque []int

func (d *Deque) push_front() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	if len(*d) == 0 {
		*d = append(*d, n)
	} else {
		*d = append([]int{n}, *d...)
	}
}

func (d *Deque) push_back() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	*d = append(*d, n)
}

func (d *Deque) pop_front() {
	if len(*d) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		p := (*d)[0]
		*d = (*d)[1:]
		fmt.Fprintln(writer, p)
	}
}

func (d *Deque) pop_back() {
	l := len(*d)
	if l == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		p := (*d)[l-1]
		*d = (*d)[:l-1]
		fmt.Fprintln(writer, p)
	}
}

func (d *Deque) size() {
	fmt.Fprintln(writer, len(*d))
}

func (d *Deque) empty() {
	if len(*d) == 0 {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func (d *Deque) front() {
	if len(*d) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, (*d)[0])
	}
}

func (d *Deque) back() {
	if len(*d) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, (*d)[len(*d)-1])
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var d Deque
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		switch command {
		case "push_front":
			d.push_front()
		case "push_back":
			d.push_back()
		case "pop_front":
			d.pop_front()
		case "pop_back":
			d.pop_back()
		case "size":
			d.size()
		case "empty":
			d.empty()
		case "front":
			d.front()
		case "back":
			d.back()
		}
	}
}
