package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	customers []Customer
	N, K      int
)

// 고객 정보
type Customer struct {
	id int // 고객 아이디
	w  int // 구매하는 물건 수 = 계산 시간
}

type Casher struct {
	DeskID   int       // 계산대 번호
	Customer *Customer // 고객 정보
}

// 사용가능한 계산대
// 사용가능한 계산대는 계산대 번호가 작은 것부터 고객을 안내한다
type Avail []*Casher

func (c Avail) Len() int { return len(c) }
func (c Avail) Less(i, j int) bool {
	if c[i].Customer.w == c[j].Customer.w {
		return c[i].DeskID < c[j].DeskID
	}
	return c[i].Customer.w < c[j].Customer.w
}
func (c Avail) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Avail) Push(x interface{}) {
	*c = append(*c, x.(*Casher))
}
func (c *Avail) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 사용하고 있는 계산대
// 사용이 끝난 계산대는 계산대 번호가 큰 것부터 고객을 내보낸다
type Taken []*Casher

func (c Taken) Len() int { return len(c) }
func (c Taken) Less(i, j int) bool {
	if c[i].Customer.w == c[j].Customer.w {
		return c[i].DeskID > c[j].DeskID
	}
	return c[i].Customer.w < c[j].Customer.w
}
func (c Taken) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Taken) Push(x interface{}) {
	*c = append(*c, x.(*Casher))
}
func (c *Taken) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 메모리: 9576KB
// 시간: 120ms
// 우선순위 큐, 계산대로 들어올 때, 나갈 때 따로 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	customers = make([]Customer, N)
	for i := 0; i < N; i++ {
		customers[i] = Customer{scanInt(), scanInt()}
	}
	taken := new(Taken)
	avail := new(Avail)
	heap.Init(taken)
	heap.Init(avail)

	// 계산대를 사용한 순서
	order := make([]int, 0, N)

	// 비어있는 계산대 먼저 채워넣기
	for i := 1; i <= K; i++ {
		if len(customers) > 0 {
			heap.Push(taken, &Casher{
				DeskID:   i,
				Customer: &customers[0],
			})
			customers = customers[1:]
		}

	}

	for len(customers) > 0 {
		// 계산이 끝난 고객 내보내기
		if taken.Len() > 0 {
			t := heap.Pop(taken).(*Casher)
			order = append(order, t.Customer.id)
			time := t.Customer.w // 현재 고객이 계산이 끝난 시간
			heap.Push(avail, t)

			// 동일한 시간에 계산이 끝난 고객 내보내기
			for taken.Len() > 0 {
				t = heap.Pop(taken).(*Casher)
				if t.Customer.w == time {
					order = append(order, t.Customer.id)
					heap.Push(avail, t)
				} else {
					heap.Push(taken, t)
					break
				}
			}
		}

		// 사용가능한 계산대로 고객 안내하기
		for avail.Len() > 0 && len(customers) > 0 {
			a := heap.Pop(avail).(*Casher)
			customers[0].w += a.Customer.w
			heap.Push(taken, &Casher{
				DeskID:   a.DeskID,
				Customer: &customers[0],
			})
			customers = customers[1:]
		}
	}

	// 아직 계산대를 사용중인 고객들 내보내기
	for taken.Len() > 0 {
		t := heap.Pop(taken).(*Casher)
		order = append(order, t.Customer.id)
	}

	sum := 0
	for i := 0; i < N; i++ {
		sum += order[i] * (i + 1)
	}
	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
