package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, V, E int      // 기사단의 팀원 수, 정점의 개수, 간선의 개수
	A, B    int      // KIST, CRFood의 위치
	knights []int    // knights[i] = j (i번째 기사단원의 위치 j)
	paths   [][]Node // paths[i] = Node{j, c} (i에서 j로 가는 비용이 c인 경로)
)

type Node struct {
	number int
	dist   int
}

const INF = 987654321

// 난이도: Gold 4
// 메모리: 1992KB
// 시간: 8ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, V, E = scanInt(), scanInt(), scanInt()
	A, B = scanInt(), scanInt()

	knights = make([]int, N+1)
	for i := 1; i <= N; i++ {
		knights[i] = scanInt()
	}

	paths = make([][]Node, V+1)
	for i := 1; i <= E; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		paths[a] = append(paths[a], Node{b, c})
		paths[b] = append(paths[b], Node{a, c})
	}
}

func Solve() {
	// KIST나 씨알푸드로 갈 수 없는 경우 -1로 처리하라는 의미가
	// -1로 처리된 최단거리를 포함하여 누적합을 구하라는 것인지
	// 결과로 -1을 출력하라는 것인지 모르겠어서
	// 두 경우를 모두 실행해보았고 두 경우 모두 정답으로 처리되었다

	ok, kist := Dijkstra(A) // KIST에서 각 기사단원까지의 최단거리가 모두 존재하는지 여부와 그 합
	if !ok {
		fmt.Fprintln(writer, -1)
		return
	}

	ok, crfood := Dijkstra(B) // CRFood에서 각 기사단원까지의 최단거리가 모두 존재하는지 여부와 그 합
	if !ok {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, kist+crfood) // KIST에서 각 기사단원까지의 최단거리 + CRFood에서 각 기사단원까지의 최단거리
}

type PQ []*Node

func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool {
	return q[i].dist < q[j].dist
}
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(*Node))
}
func (q *PQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

// 경로가 존재하지 않는 경우 false를 반환하는 방법
func Dijkstra(start int) (bool, int) {
	// dist[i] = j (start에서 i까지의 최단거리가 j)
	dist := make([]int, V+1)

	// 모든 정점까지의 최단거리를 INF로 초기화
	for i := 1; i <= V; i++ {
		dist[i] = INF
	}

	// 시작 위치의 최단거리를 0으로 초기화
	dist[start] = 0

	// 우선순위 큐 초기화
	q := new(PQ)
	heap.Init(q)

	// 시작 위치를 우선순위 큐에 삽입
	heap.Push(q, &Node{start, 0})

	// 우선순위 큐가 빌 때까지 반복
	for len(*q) > 0 {
		root := heap.Pop(q).(*Node)

		here := root.number
		acc := root.dist

		// 최단거리가 아닌 경우 스킵
		if acc > dist[here] {
			continue
		}

		// 인접한 정점들을 검사
		for _, conn := range paths[here] {
			next := conn.number
			extra := conn.dist

			// 최단거리 갱신
			if dist[next] > acc+extra {
				dist[next] = acc + extra
				heap.Push(q, &Node{next, dist[next]})
			}
		}
	}

	sum := 0

	// 모든 기사단원까지의 최단거리가 존재하는지 여부와 최단거리의 누적합을 구한다
	for i := 1; i <= N; i++ {
		// 최단거리가 존재하지 않는다면 false를 반환
		if dist[knights[i]] == INF {
			return false, 0
		}
		sum += dist[knights[i]]
	}

	return true, sum
}

// 경로가 존재하지 않는 경우 -1로 처리하여 누적합을 구하는 방법
func Dijkstra2(start int) int {
	dist := make([]int, V+1)
	for i := 1; i <= V; i++ {
		dist[i] = INF
	}

	dist[start] = 0

	q := new(PQ)
	heap.Init(q)
	heap.Push(q, &Node{start, 0})

	for len(*q) > 0 {
		root := heap.Pop(q).(*Node)

		here := root.number
		acc := root.dist

		if acc > dist[here] {
			continue
		}

		for _, conn := range paths[here] {
			next := conn.number
			extra := conn.dist

			if dist[next] > acc+extra {
				dist[next] = acc + extra
				heap.Push(q, &Node{next, dist[next]})
			}
		}
	}

	sum := 0

	for i := 1; i <= N; i++ {
		if dist[knights[i]] == INF {
			sum += -1
		} else {
			sum += dist[knights[i]]
		}
	}

	return sum
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
