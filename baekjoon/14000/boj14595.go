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
	parent  []int
	N, M    int
)

// 메모리: 9864KB
// 시간: 68ms
// 분리 집합: 빅-종빈빌런은 매우 허당이기 때문에 동일한 행동을 여러 번 할 수 있음에 유의하라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	var a, b int
	for i := 1; i <= M; i++ {
		a, b = scanInt(), scanInt()

		// 동일한 행동을 반복할 수 있으므로
		// 이미 부순 벽은 건너뛰고 경계부분부터 Union
		for i := a; i <= b; {
			x, y := Find(i), Find(b)
			if x == y {
				break
			}

			next := x + 1
			parent[x] = y
			i = next
		}
	}

	cnt := 0
	for i := 1; i <= N; i++ {
		// Find 연산으로 확실하게 처리
		if Find(i) == i {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}

func Find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
