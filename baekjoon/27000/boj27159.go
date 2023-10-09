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
	N       int
)

// 난이도: Bronze 3
// 메모리: 912KB
// 시간: 4ms
// 분류: 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
}

func Solve() {
	head := 0
	tail := 0
	sum := 0

	for i := 1; i <= N; i++ {
		x := scanInt()
		if x-1 == tail {
			tail = x
		} else {
			head, tail = x, x
			sum += head
		}
	}

	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
