package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, Q    int
	back    []int
	forward []int
	current int
)

// 난이도: Gold 5
// 메모리: 5688KB
// 시간: 20ms
// 분류: 구현, 자료구조, 덱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Result()
}

func Input() {
	N, Q = scanInt(), scanInt()
	for i := 1; i <= Q; i++ {
		Query()
	}
}

func Query() {
	cmd := scanCmd()
	switch cmd {
	case 'B':
		if len(back) > 0 {
			forward = append([]int{current}, forward...)
			current = back[len(back)-1]
			back = back[:len(back)-1]
		}
	case 'F':
		if len(forward) > 0 {
			back = append(back, current)
			current = forward[0]
			forward = forward[1:]
		}
	case 'A':
		visit := scanInt()
		forward = []int{}
		if current != 0 {
			back = append(back, current)
		}
		current = visit
	case 'C':
		Compress()
	}
}

func Compress() {
	cmp := []int{}
	for len(back) > 0 {
		b := back[0]
		back = back[1:]

		if len(cmp) > 0 && cmp[len(cmp)-1] == b {
			continue
		}

		cmp = append(cmp, b)
	}
	back = cmp
}

func Result() {
	fmt.Fprintln(writer, current)
	if len(back) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		for i := len(back) - 1; i >= 0; i-- {
			fmt.Fprintf(writer, "%d ", back[i])
		}
		fmt.Fprintln(writer)
	}
	if len(forward) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		for _, v := range forward {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	}
}

func scanCmd() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
