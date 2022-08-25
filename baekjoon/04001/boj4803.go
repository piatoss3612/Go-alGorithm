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
	n, m    int
	parent  [501]int
)

// 분리 집합을 구하기 위한 parent 슬라이스 초기화
func init() {
	for i := 1; i <= 500; i++ {
		parent[i] = i
	}
}

// 메모리: 2044KB
// 시간: 44ms
// 분리 집합, 메모리 초과 발생 2회
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	count := 1

	for {
		n, m = scanInt(), scanInt()
		if n == 0 && m == 0 {
			return
		}

		var a, b int
		for i := 1; i <= m; i++ {
			a, b = scanInt(), scanInt()

			union(a, b) // a와 b 정점 연결
		}

		tree := 0
		for i := 1; i <= n; i++ {
			// i번째 정점의 루트 정점이 i인 경우
			// 해당 정점은 트리를 구성하고 있음
			if find(i) == i {
				tree++
			}
		}

		if tree > 1 {
			fmt.Fprintf(writer, "Case %d: A forest of %d trees.\n", count, tree)
		} else if tree == 1 {
			fmt.Fprintf(writer, "Case %d: There is one tree.\n", count)
		} else {
			fmt.Fprintf(writer, "Case %d: No trees.\n", count)
		}

		for i := 1; i <= n; i++ {
			parent[i] = i
		}
		count++
	}
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		// 메모리초과 해결:
		// x가 0이거나 y가 0인 경우는 어느 한쪽 정점이 사이클에 포함되어 있는 상태이므로
		// 다른 정점도 사이클에 포함된 것으로 분류를 하여야 한다
		if x == 0 || y == 0 {
			parent[x] = 0
			parent[y] = 0
		} else {
			parent[y] = x
		}
		// 부모 정점이 같은 경우는 사이클이 형성되어 트리를 구성할 수 없다
	} else if x == y {
		parent[x] = 0
		parent[y] = 0
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
