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
	level   []int // 악보 난이도
	mistake []int // 실수 누적 합
)

// 난이도: Silver 1
// 메모리: 5864KB
// 시간: 76ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Solve()
}

func Solve() {
	N = scanInt()
	level = make([]int, N+1)
	mistake = make([]int, N+1)
	for i := 1; i <= N; i++ {
		level[i] = scanInt()
		if i != 1 {
			if level[i-1] > level[i] {
				// i-1번 악보가 i번 악보보다 난이도가 높은 경우
				mistake[i] = mistake[i-1] + 1 // 실수 추가
			} else {
				mistake[i] = mistake[i-1]
			}
		}
	}

	Q = scanInt()
	var x, y int
	for i := 1; i <= Q; i++ {
		x, y = scanInt(), scanInt()
		fmt.Fprintln(writer, mistake[y]-mistake[x])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
