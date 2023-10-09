package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	n, m     int
	inDegree []int   // 진입차수
	graph    [][]int // 연결 상태
)

// 메모리: 7456KB
// 시간: 56ms
// 위상 정렬 + 최소 힙(우선순위 큐)로 푸는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	inDegree = make([]int, n+1)
	graph = make([][]int, n+1)

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		graph[a] = append(graph[a], b) // 그래프는 a -> b 형태
		inDegree[b] += 1               // b로 진입차수 1증가
	}

	topologicalSort() // 위상 정렬
}

// 최소 힙 정의
type Queue []int

func (q Queue) Len() int { return len(q) }
func (q Queue) Less(i, j int) bool {
	return q[i] < q[j]
}
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func topologicalSort() {
	var res []int
	q := &Queue{}
	heap.Init(q) // 최소 힙 초기화

	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			heap.Push(q, i)
		}
	}

	for len(*q) > 0 {
		x := heap.Pop(q).(int)
		res = append(res, x)

		for i := 0; i < len(graph[x]); i++ {
			inDegree[graph[x][i]] -= 1
			if inDegree[graph[x][i]] == 0 {
				heap.Push(q, graph[x][i])
			}
		}
	}

	// 항상 문제를 모두 풀 수 있는 경우만 입력으로 주어지므로 res를 그대로 출력
	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
