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

	N      int
	coords [][2]int
)

// 난이도: Gold 5
// 메모리: 984KB
// 시간: 3164ms
// 분류: 브루트포스 알고리즘, 기하학, 피타고라스의 정리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	coords = make([][2]int, N+1)
	for i := 1; i <= N; i++ {
		coords[i][0] = scanInt()
		coords[i][1] = scanInt()
	}
}

func Solve() {
	cnt := 0

	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			for k := j + 1; k <= N; k++ {
				if isRightTriangle(i, j, k) {
					cnt++
				}
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

func isRightTriangle(i, j, k int) bool {
	x1, y1 := coords[i][0], coords[i][1]
	x2, y2 := coords[j][0], coords[j][1]
	x3, y3 := coords[k][0], coords[k][1]

	a := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	b := (x1-x3)*(x1-x3) + (y1-y3)*(y1-y3)
	c := (x2-x3)*(x2-x3) + (y2-y3)*(y2-y3)

	if a == b+c || b == a+c || c == a+b {
		return true
	}

	return false
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
