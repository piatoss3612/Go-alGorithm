package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	names    []string
	inDegree []int
	graph    [][]int
	visited  []bool
)

// 31423번: 신촌 통폐합 계획
// hhttps://www.acmicpc.net/problem/31423
// 난이도: 골드 5
// 메모리: 115904 KB
// 시간: 520 ms
// 분류: 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	names = make([]string, N+1)
	for i := 1; i <= N; i++ {
		names[i] = scanString()
	}

	inDegree = make([]int, N+1)
	graph = make([][]int, N+1)

	for i := 1; i < N; i++ {
		a, b := scanInt(), scanInt()
		graph[a] = append(graph[a], b)
		inDegree[b]++
	}

	visited = make([]bool, N+1)
}

func Solve() {
	idxChan := make(chan int, N)

	// 문제의 맥락상 트리의 루트는 하나이므로 in degree가 0인 정점을 찾아서 DFS를 시작한다.

	go func() {
		defer close(idxChan)
		for i := 1; i <= N; i++ {
			if inDegree[i] == 0 {
				DFS(i, idxChan)
				break
			}
		}
	}()

	builder := strings.Builder{}
	builder.Grow(1000000)

	for idx := range idxChan {
		builder.WriteString(names[idx])
	}

	fmt.Fprintln(writer, builder.String())
}

func DFS(cur int, idxChan chan<- int) {
	visited[cur] = true
	idxChan <- cur

	for _, next := range graph[cur] {
		if !visited[next] {
			DFS(next, idxChan)
		}
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
