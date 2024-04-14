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
	N, M, K int
	X, Y    int
	parent  []int
)

// 23324번: 어려운 모든 정점 쌍 최단 거리
// hhttps://www.acmicpc.net/problem/23324
// 난이도: 골드 4
// 메모리: 4820 KB
// 시간: 68 ms
// 분류: 그래프 이론, 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}
}

func Solve() {
	// 크루스칼 알고리즘을 돌리면 N이 최대 100,000이므로 반드시 시간초과가 난다.
	// X-Y 간선만 가중치가 1이기 때문에 X와 Y를 연결하는 간선을 우선 제외해보자.
	// 그리고 X와 연결된 노드의 개수를 cntX, Y와 연결된 노드의 개수를 cntY라고 하자.
	// X와 Y를 제외한 나머지 노드들은
	// 1. X와 Y 중 하나와 연결되어 있거나
	// 2. 둘 다와 연결되어 있거나
	// 3. 둘 다와 연결되어 있지 않다.
	// 만약 X와 Y 둘 다 연결되어 있는 노드가 하나라도 있다면
	// 해당 노드를 통하여 최단 거리를 구하면 모든 정점 쌍의 최단 거리 합은 항상 0이 된다.
	// 반면 X와 연결된 노드들은 반드시 X-Y를 거쳐 Y와 연결된 노드들로 가야하고
	// Y와 연결된 노드들은 반드시 X-Y를 거쳐 X와 연결된 노드들로 가야한다.
	// 따라서 X와 연결된 노드의 개수를 cntX, Y와 연결된 노드의 개수를 cntY라고 하면
	// X와 Y를 각 그룹에 포함하여 나머지 노드들을 연결하는 경우의 수는 (cntX+1)*(cntY+1)이 된다.

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		if i == K {
			X, Y = a, b
			continue
		}

		union(a, b)
	}

	parentX := findParent(X)
	parentY := findParent(Y)

	if parentX == parentY {
		fmt.Fprintln(writer, 0)
		return
	}

	cntX, cntY := 0, 0

	for i := 1; i <= N; i++ {
		if i == X || i == Y {
			continue
		}

		parentI := findParent(i)

		if parentI == parentX {
			cntX++
		} else if parentI == parentY {
			cntY++
		}
	}

	fmt.Fprintln(writer, (cntX+1)*(cntY+1))
}

func findParent(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = findParent(parent[x])
	return parent[x]
}

func union(a, b int) {
	a, b = findParent(a), findParent(b)
	if a != b {
		parent[b] = a
	}
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
