package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	A, B, C int
	visited [201][201][201]bool
)

// 난이도: Gold 5
// 메모리: 2652KB
// 시간: 8ms
// 분류: 너비 우선 탐색, 그래프 이론, 그래프 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	A, B, C = scanInt(), scanInt(), scanInt()
}

func Solve() {
	q := [][3]int{{0, 0, C}}
	visited[0][0][C] = true

	candidates := map[int]bool{} // A 물통이 비어있을 때 C 물통에 담길 수 있는 물의 양의 후보들

	for len(q) > 0 {
		a, b, c := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if a == 0 {
			candidates[c] = true
		}

		// a -> b (a가 비어있지 않고 b가 가득 차지 않았을 때)
		if a > 0 && b < B {
			na, nb := 0, a+b
			if nb > B {
				na, nb = nb-B, B
			}
			if !visited[na][nb][c] {
				visited[na][nb][c] = true
				q = append(q, [3]int{na, nb, c})
			}
		}

		// a -> c (a가 비어있지 않고 c가 가득 차지 않았을 때)
		if a > 0 && c < C {
			na, nc := 0, a+c
			if nc > C {
				na, nc = nc-C, C
			}
			if !visited[na][b][nc] {
				visited[na][b][nc] = true
				q = append(q, [3]int{na, b, nc})
			}
		}

		// b -> a (b가 비어있지 않고 a가 가득 차지 않았을 때)
		if b > 0 && a < A {
			na, nb := a+b, 0
			if na > A {
				na, nb = A, na-A
			}
			if !visited[na][nb][c] {
				visited[na][nb][c] = true
				q = append(q, [3]int{na, nb, c})
			}
		}

		// b -> c (b가 비어있지 않고 c가 가득 차지 않았을 때)
		if b > 0 && c < C {
			nb, nc := 0, b+c
			if nc > C {
				nb, nc = nc-C, C
			}
			if !visited[a][nb][nc] {
				visited[a][nb][nc] = true
				q = append(q, [3]int{a, nb, nc})
			}
		}

		// c -> a (c가 비어있지 않고 a가 가득 차지 않았을 때)
		if c > 0 && a < A {
			na, nc := a+c, 0
			if na > A {
				na, nc = A, na-A
			}
			if !visited[na][b][nc] {
				visited[na][b][nc] = true
				q = append(q, [3]int{na, b, nc})
			}
		}

		// c -> b (c가 비어있지 않고 b가 가득 차지 않았을 때)
		if c > 0 && b < B {
			nb, nc := b+c, 0
			if nb > B {
				nb, nc = B, nb-B
			}
			if !visited[a][nb][nc] {
				visited[a][nb][nc] = true
				q = append(q, [3]int{a, nb, nc})
			}
		}
	}

	ans := []int{}

	for k := range candidates {
		ans = append(ans, k)
	}

	sort.Ints(ans)

	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
