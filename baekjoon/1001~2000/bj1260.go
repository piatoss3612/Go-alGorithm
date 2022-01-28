package bj1260

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, v := scanInt(), scanInt(), scanInt()
	graph = make([][]int, n+1) // 작은 수를 우선 탐색하기 위해 (n + 1)*(n + 1) 크기로 슬라이스 초기화
	visited = make([]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		graph[a][b] = 1
		graph[b][a] = 1
	}
	DFS(v) // 깊이 우선 탐색
	fmt.Fprintln(writer)
	visited = make([]bool, n+1) // 방문 여부 초기화
	BFS(v)                      // 너비 우선 탐색
	fmt.Fprintln(writer)
}

func DFS(v int) {
	if visited[v] == true {
		return
	}
	visited[v] = true
	fmt.Fprint(writer, v, " ")
	for idx, next := range graph[v] {
		if next == 1 {
			DFS(idx)
		}
	}
}

func BFS(v int) {
	visited[v] = true
	queue := []int{v}
	for len(queue) > 0 {
		front := queue[0]
		fmt.Fprint(writer, front, " ")
		queue = queue[1:]
		for i := 0; i < len(graph[front]); i++ {
			if graph[front][i] == 1 && !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
