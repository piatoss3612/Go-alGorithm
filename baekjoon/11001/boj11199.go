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
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, T       int
	customers  []*Customer // 고객들
	candidates *Candidates // 고객들의 예금액 중 가장 큰 금액을 반환하는 우선순위 큐(최대 힙)
)

type Customer struct {
	amount int // 예금할 금액
	ttl    int // 떠나는 시간
}

// 우선순위 큐 정의 및 구현
type Candidates []int

func (c Candidates) Len() int { return len(c) }
func (c Candidates) Less(i, j int) bool {
	return c[i] > c[j]
}
func (c Candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Candidates) Push(x interface{}) {
	*c = append(*c, x.(int))
}
func (c *Candidates) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 1724KB
// 시간: 8ms
// 분류: 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, T = scanInt(), scanInt()
	customers = make([]*Customer, N)

	for i := 0; i < N; i++ {
		customers[i] = &Customer{scanInt(), scanInt()}
	}

	// 나중에 떠나는 고객 ~ 먼저 떠나는 고객 순으로 오름차순 정렬
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].ttl > customers[j].ttl
	})
}

func Solve() {
	candidates = new(Candidates)
	heap.Init(candidates)

	// 예외 case: 나중에 떠나는 고객들을 먼저 처리해야 예금액의 최댓값을 구할 수 있는 경우
	// 예를 들어,
	// N=4, T=4,
	// 고객1 - amount: 1000, ttl: 0
	// 고객2 - amount: 1500, ttl: 1
	// 고객3 - amount: 2000, ttl: 2
	// 고객4 - amount: 3000, ttl: 2
	// 인 경우
	// 처음으로 고객 2, 1분 뒤에 고객3, 2분 뒤에 고객4, 3분 뒤에는 처리할 수 있는 고객 없음
	// 위의 순서로 요청을 처리하면 예금액이 6500으로 최댓값이 된다
	// 3번 고객과 4번 고객의 순서를 바꿔도 가능하다

	totalAmount, time := 0, T-1

	// 그리디 알고리즘:
	// time값을 마감하기 전 마지막으로 고객을 받는 시간부터 고객을 최초로 받기 시작하는 시간으로 거를러 올라가면서 (T-1 ~ 0)
	// time 시간에 떠나지 않고 대기중인 고객들을 후보로 색출, 후보들 중에 예금액이 가장 큰 고객을 time 시간에 맞춰 처리하면 최적해를 구할 수 있다
	for time >= 0 {
		for len(customers) > 0 {
			if customers[0].ttl == time {
				heap.Push(candidates, customers[0].amount)
				customers = customers[1:]
			} else {
				break
			}
		}

		// time 시간에 처리할 수 있는 고객 요청이 있는 경우에만
		if len(*candidates) > 0 {
			picked := heap.Pop(candidates).(int)
			totalAmount += picked
		}

		time--
	}

	fmt.Fprintln(writer, totalAmount)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
