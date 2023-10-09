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
	lo      *PQ
	hi      *PQ
	waiting map[int]int // 대기중인 클라이언트 아이디와 우선순위 매핑
)

// 클라이언트 정보
type Client struct {
	ID       int
	Priority int
}

// 우선순위 큐 정의 및 구현
type PQ []*Client

func (m PQ) Len() int           { return len(m) }
func (m PQ) Less(i, j int) bool { return m[i].Priority > m[j].Priority }
func (m PQ) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *PQ) Push(x interface{}) {
	*m = append(*m, x.(*Client))
}
func (m *PQ) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 7224KB
// 시간: 116ms
// 분류: 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Query()
}

func Setup() {
	lo = new(PQ)
	hi = new(PQ)
	heap.Init(lo)
	heap.Init(hi)
	waiting = make(map[int]int) // 요청 사항을 처리하고 나면 동일한 아이디의 클라이언트가 이전과 다른 우선순위로 대기열에 추가될 수 있다
}

func Query() {
	for {
		q := scanInt()
		if q == 0 {
			return
		}

		switch q {
		case 1:
			addClient()
		case 2:
			highestPriority()
		case 3:
			lowestPriority()
		}
	}
}

func addClient() {
	id, pr := scanInt(), scanInt()
	waiting[id] = pr
	heap.Push(hi, &Client{id, pr})
	heap.Push(lo, &Client{id, -pr}) // 우선순위를 음수로 변형하여 집어넣으면 우선순위가 낮은 클라이언트를 먼저 꺼내온다
}

func highestPriority() {
	for len(*hi) > 0 {
		client := heap.Pop(hi).(*Client)
		// 대기열에 있는 클라이언트의 우선순위와 우선순위 큐에서 꺼내온 클라이언트의 우선순위가 일치하는 경우
		if waiting[client.ID] == client.Priority {
			fmt.Fprintln(writer, client.ID)
			delete(waiting, client.ID)
			return
		}
		// 우선순위가 일치할 때까지 반복
	}

	if len(*hi) == 0 {
		fmt.Fprintln(writer, 0)
		return
	}
}

func lowestPriority() {
	for len(*lo) > 0 {
		client := heap.Pop(lo).(*Client)
		// 대기열에 있는 클라이언트의 우선순위와 우선순위 큐에서 꺼내온 클라이언트의 우선순위가 일치하는 경우
		if waiting[client.ID] == -client.Priority {
			fmt.Fprintln(writer, client.ID)
			delete(waiting, client.ID)
			return
		}
		// 우선순위가 일치할 때까지 반복
	}

	if len(*lo) == 0 {
		fmt.Fprintln(writer, 0)
		return
	}
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
