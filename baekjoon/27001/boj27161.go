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
	N, X    int
	S       string
)

// 난이도: Bronze 1
// 메모리: 932KB
// 시간: 8ms
// 분류: 구현, 시뮬레이션
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
	time := 0
	reverse := false

	for i := 1; i <= N; i++ {
		S, X = scanString(), scanInt()

		if reverse {
			time = timeReverse(time)
		} else {
			time = timeSteram(time)
		}

		rv := isReverse(S)
		sc := isSync(time, X)

		switch {
		case rv && sc == true:
			fmt.Fprintln(writer, time, "NO")
		case rv == true:
			reverse = !reverse
			fmt.Fprintln(writer, time, "NO")
		case sc == true:
			fmt.Fprintln(writer, time, "YES")
		default:
			fmt.Fprintln(writer, time, "NO")
		}
	}
}

func timeSteram(t int) int {
	if t == 12 {
		return 1
	}
	return t + 1
}

func timeReverse(t int) int {
	if t <= 1 {
		return 12
	}
	return t - 1
}

func isReverse(s string) bool {
	return s == "HOURGLASS"
}

func isSync(a, b int) bool {
	return a == b
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
