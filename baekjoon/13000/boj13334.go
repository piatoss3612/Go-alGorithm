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
	n, d    int
	people  []Person
	pq      *PQ
)

type Person struct {
	h, o int
}

type PQ []*Person

func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool {
	// 집의 위치를 기준으로 가장 왼쪽에 있는 사람부터 꺼내온다
	return q[i].h < q[j].h
}
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(*Person))
}
func (q *PQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

// 난이도: Gold 2
// 메모리: 8136KB
// 시간: 108ms
// 분류: 우선순위 큐, 스위핑, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	n = scanInt()
	people = make([]Person, 0, n)
	for i := 0; i < n; i++ {
		h, o := swap(scanInt(), scanInt()) // 연산과정에서 대소 비교를 줄이기 위해 항상 h가 o보다 작은 값이 되도록 스왑
		people = append(people, Person{h, o})
	}

	pq = new(PQ)
	heap.Init(pq)

	// ** 사무실 위치를 기준으로 오름차순 정렬 **
	// 집 위치를 기준으로 시도해보고 안돼서 사무실 위치를 기준으로 시도 ->
	// 왼쪽(집)을 기준으로 철로의 왼쪽이나 오른쪽에 있는 사람들을 걸러내는 것보다
	// 오른쪽(사무실)을 기준으로 철로의 왼쪽에 있는 사람들을 걸러내는 것이 더 간단하고 직관적
	sort.Slice(people, func(i, j int) bool {
		return people[i].o < people[j].o
	})

	d = scanInt()
}

func Solve() {
	ans := 0
	l, r := -100000001, -100000001

	for len(people) > 0 {
		if len(*pq) == 0 { // 철로에 포함된 사람이 없는 경우
			for len(people) > 0 {
				t := people[0]
				people = people[1:]

				// 사람 t의 집에서 회사까지의 거리가 d보다 작거나 같은 경우에만
				if t.o-t.h <= d {
					l, r = t.h, t.o
					heap.Push(pq, &t)
					break
				}
			}
		} else {
			t := people[0]
			people = people[1:]

			// 사람 t의 집에서 회사까지의 거리가 d보다 큰 경우: 제외
			if t.o-t.h > d {
				continue
			}

			if l <= t.h && t.o <= r { // 사람 t의 집에서 회사까지의 거리가 l과 r 사이(r-l <= d)에 있는 경우
				heap.Push(pq, &t)
			} else { // 사람 t의 집에서 회사까지의 거리가 l과 r 범위를 벗어난 경우
				l, r = min(l, t.h), max(r, t.o) // l, r을 갱신

				// r-l <= d를 만족할 때까지 큐에서 사람 제거
				for r-l > d && len(*pq) > 0 {
					y := heap.Pop(pq).(*Person)
					l = max(l, y.h) // 오른쪽은 r로 고정되어 있으므로 l만 갱신
					if r-l <= d {
						heap.Push(pq, y)
						break
					}
				}
				heap.Push(pq, &t)
			}
		}
		ans = max(ans, len(*pq)) // 철로의 포함된 사람의 최대 수 갱신
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func swap(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
