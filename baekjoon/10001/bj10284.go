package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n, d, c int
	rels    [][]infection // 감염 관계
	INF     = 987654321
)

type infection struct {
	target, time int // 감영 대상과 걸리는 시간
}

// 메모리: 10644KB
// 시간: 196ms
// 다익스트라 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		TestCase()
	}
}

func TestCase() {
	n, d, c = scanInt(), scanInt(), scanInt()
	rels = make([][]infection, n+1)

	for i := 1; i <= d; i++ {
		a, b, s := scanInt(), scanInt(), scanInt()
		rels[b] = append(rels[b], infection{a, s}) // b가 감염되면 s초 후에 a가 감염되는 관계
	}

	Dijkstra(c, 1) // c 최초 감염
}

// 최소 힙: 감염에 걸리는 시간이 적은 대상부터 힙에서 꺼내온다
type Heap []*infection

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].time < h[j].time
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*infection))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 다익스트라
func Dijkstra(v, cnt int) {
	seconds := make([]int, n+1)  // 마지막 컴퓨터가 감염되기까지 걸리는 시간
	checked := make([]bool, n+1) // 감염 여부
	checked[v] = true
	for i := 1; i <= n; i++ {
		seconds[i] = INF // i -> j로 가는 최솟값을 찾기 위해 seconds 슬라이스를 INF로 초기화
	}

	h := &Heap{{v, 0}} // 최소 힙 초기화, v가 최초로 감염되고 아직 다른 컴퓨터로 이동하지 않은 상황
	heap.Init(h)

	for h.Len() > 0 {
		current := heap.Pop(h).(*infection)
		infected := current.target
		taken := current.time

		// infected가 감염되는 시간이 꺼내온 taken보다 이미 작다면
		if seconds[infected] < taken {
			continue
		}

		// infected의 감염으로 인한 바이러스 전파 과정
		for _, next := range rels[infected] {
			// 다음 컴퓨터가 감염되는 최소 시간을 갱신
			if seconds[next.target] > taken+next.time {
				seconds[next.target] = taken + next.time
				heap.Push(h, &infection{next.target, seconds[next.target]})
				// 아직 감염된 적 없는 컴퓨터라면
				if !checked[next.target] {
					checked[next.target] = true
					cnt++ // 감염된 컴퓨터의 수 증가
				}
			}
		}
	}

	sort.Ints(seconds) // 오름차순 정렬

	maxSeconds := 0

	// 모든 컴퓨터가 감염되는데 걸리는 시간은 INF를 제외한 최댓값이 되어야 한다
	for i := 1; i <= n; i++ {
		if seconds[i] != INF {
			maxSeconds = seconds[i]
		} else {
			break
		}
	}

	fmt.Fprintln(writer, cnt, maxSeconds)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
