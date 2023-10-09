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
	N       int
	cows    []*Cow
	pq      *Cows
	ans     int // 얻을 수 있는 우유의 최대량
)

// 소의 정보
type Cow struct {
	amount int
	time   int
}

// 우유의 양이 가장 적은 소를 반환하는 우선순위 큐
type Cows []*Cow

func (c Cows) Len() int { return len(c) }
func (c Cows) Less(i, j int) bool {
	return c[i].amount < c[j].amount
}
func (c Cows) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c *Cows) Push(x interface{}) {
	*c = append(*c, x.(*Cow))
}
func (c *Cows) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 난이도: Gold 3
// 메모리: 1620KB
// 시간: 8ms
// 분류: 그리디 알고리즘, 정렬, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	cows = make([]*Cow, N)
	for i := 0; i < N; i++ {
		cows[i] = &Cow{scanInt(), scanInt()}
	}

	// 우유를 짜야 하는 데드라인이 빠른 소를 기준으로 오름차순 정렬
	sort.Slice(cows, func(i, j int) bool {
		if cows[i].time == cows[j].time {
			return cows[i].amount > cows[j].amount
		}
		return cows[i].time < cows[j].time
	})
}

func Solve() {
	pq = new(Cows)
	heap.Init(pq)

	for len(cows) > 0 {
		cow := cows[0]
		cows = cows[1:]

		// 우유를 짜는데 걸리는 시간은 1
		// pq의 길이는 전체 걸리는 시간과 마찬가지

		if len(*pq) < cow.time {
			// 현재 선택된 소 cow의 데드라인이 pq의 길이보다 큰 경우
			heap.Push(pq, cow) // cow를 pq에 추가
			ans += cow.amount
		} else {
			// cow의 데드라인 안에 작업을 끝내기 어려운 경우

			min := heap.Pop(pq).(*Cow)
			// 우유의 양이 가장 적은 소 min을 pq에서 꺼내와 cow의 우유량과 비교
			if min.amount < cow.amount {
				// cow의 우유량이 더 많은 경우
				heap.Push(pq, cow) // cow를 pq에 추가
				ans += cow.amount - min.amount
			} else {
				heap.Push(pq, min) // min을 pq로 되돌려 놓음
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
