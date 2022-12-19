package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner       = bufio.NewScanner(os.Stdin)
	writer        = bufio.NewWriter(os.Stdout)
	F, S, G, U, D int
	visited       []bool
)

// 난이도: Silver 1
// 메모리: 10276KB
// 시간: 88ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	F, S, G, U, D = scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	visited = make([]bool, F+1)
}

type Move struct {
	floor int
	turn  int
}

func Solve() {
	ans := -1

	q := []Move{}
	q = append(q, Move{S, 0})

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.floor == G {
			ans = curr.turn
			break
		}

		if up := curr.floor + U; up <= F && !visited[up] {
			visited[up] = true
			q = append(q, Move{up, curr.turn + 1})
		}

		if down := curr.floor - D; down >= 1 && !visited[down] {
			visited[down] = true
			q = append(q, Move{down, curr.turn + 1})
		}
	}

	if ans == -1 {
		fmt.Fprintln(writer, "use the stairs")
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
