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
	T, N, M int
	friends [][]Friend // 친구 관계
)

const INF = 987654321

type Friend struct {
	y, z int
}

// 데이크스트라 알고리즘에서 우선순위 큐를 사용하기 위한 구현
type PQ []*Friend

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].z < pq[j].z // 친밀도가 높은 것(1: 최측근, 4: 지인)이 우선순위가 높음
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Friend))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve(i)
	}
}

func Setup() {
	N, M = scanInt(), scanInt()
	friends = make([][]Friend, M)
	for i := 1; i <= N; i++ {
		x, y, z := scanInt(), scanInt(), scanInt()
		// #주의# 친구는 양방향 관계이므로 양쪽에 모두 추가해줘야 한다
		friends[x] = append(friends[x], Friend{y, z})
		friends[y] = append(friends[y], Friend{x, z})
	}
}

func Solve(caseNo int) {
	pq := new(PQ)
	heap.Init(pq)

	dist := make([]int, M) // 친밀도
	prev := make([]int, M) // 이전 친구

	// 최솟값을 찾기 위해 초기값을 INF로 설정
	for i := 0; i < M; i++ {
		dist[i] = INF
	}
	dist[0] = 0 // 0번 한신이로부터 시작

	heap.Push(pq, &Friend{0, 0})

	for len(*pq) > 0 {
		head := heap.Pop(pq).(*Friend)

		x := head.y
		z := head.z

		// 이미 최솟값으로 갱신된 노드는 무시
		if dist[x] < z {
			continue
		}

		for _, f := range friends[x] {
			if dist[f.y] > z+f.z {
				dist[f.y] = z + f.z // 친밀도 갱신
				prev[f.y] = x       // 이전 친구 갱신
				heap.Push(pq, &Friend{f.y, dist[f.y]})
			}
		}
	}

	// 최고위원을 만날 수 없는 경우
	if dist[M-1] == INF {
		fmt.Fprintf(writer, "Case #%d: %d\n", caseNo, -1)
		return
	}

	fmt.Fprintf(writer, "Case #%d: ", caseNo)

	// 최고위원을 만날 수 있는 경우
	curr := M - 1
	track := []int{}
	track = append(track, curr)

	// 최고위원을 만날 수 있는 경로를 역추적
	for curr != 0 {
		curr = prev[curr]
		track = append(track, curr)
	}

	// 경로를 출력
	for i := len(track) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", track[i])
	}

	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
