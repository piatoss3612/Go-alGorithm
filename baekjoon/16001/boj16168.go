package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	graph    [][]int
	visited  []bool
	inDegree []int
	V, E     int
)

// 메모리: 1396KB
// 시간: 4ms
// 깊이 우선 탐색, 오일러 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	V, E = scanInt(), scanInt()

	// 1. 그래프 정보 및 진입 차수 정보 갱신
	graph = make([][]int, V+1)
	inDegree = make([]int, V+1)
	for i := 1; i <= E; i++ {
		a, b := scanInt(), scanInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		inDegree[a]++
		inDegree[b]++
	}

	// 2. 모든 정점이 연결되어 있는지 dfs 실행 횟수를 통해 확인
	dfsCount := 0
	visited = make([]bool, V+1)
	for i := 1; i <= V; i++ {
		if !visited[i] {
			dfs(i)
			dfsCount++
		}
	}

	// dfs 실행 횟수가 1보다 크다면 한붓그리기 불가능
	if dfsCount > 1 {
		fmt.Fprintln(writer, "NO")
		return
	}

	// 3. 오일러 경로의 필요충분 조건 확인

	// 출발, 도착 정점의 진입 차수가 홀수이며 나머지 정점의 진입 차수는 짝수인 경우
	// 모든 정점의 진입 차수가 짝수인 경우
	// 즉, 진입 차수가 홀수인 정점의 개수가 2개이거나 0개인 경우를 찾는다
	odd := 0
	for i := 1; i <= V; i++ {
		if inDegree[i]%2 != 0 {
			odd++
		}
	}

	if odd == 0 || odd == 2 {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func dfs(v int) {
	visited[v] = true
	for i := 0; i < len(graph[v]); i++ {
		if !visited[graph[v][i]] {
			dfs(graph[v][i])
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
