package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N          int
	HalliGalli map[string]int
)

// 난이도: Bronze 2
// 메모리: 2520KB
// 시간: 32ms
// 분류: 구현, 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	HalliGalli = make(map[string]int)
	for i := 1; i <= N; i++ {
		HalliGalli[scanString()] += scanInt()
	}
}

func Solve() {
	for _, cnt := range HalliGalli {
		if cnt == 5 {
			fmt.Fprintln(writer, "YES")
			return
		}
	}
	fmt.Fprintln(writer, "NO")
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
