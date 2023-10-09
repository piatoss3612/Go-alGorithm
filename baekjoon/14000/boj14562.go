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

	C    int
	S, T int
)

// 난이도: Silver 1
// 메모리: 7844KB
// 시간: 32ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	C = scanInt()
	for i := 1; i <= C; i++ {
		Input()
		Solve()
	}
}

func Input() {
	S, T = scanInt(), scanInt()
}

func Solve() {
	q := [][3]int{}
	q = append(q, [3]int{S, T, 0})

	ans := 987654321

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		player, enemy, try := front[0], front[1], front[2]

		if player == enemy {
			ans = min(ans, try)
			continue
		}

		if player+1 <= enemy {
			q = append(q, [3]int{player + 1, enemy, try + 1})
		}

		if player*2 <= enemy+3 {
			q = append(q, [3]int{player * 2, enemy + 3, try + 1})
		}
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
