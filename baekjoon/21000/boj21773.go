package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T, n    int
)

// 프로세스 정보
type process struct {
	id       int
	time     int
	priority int
}

type PQ []*process // 우선순위 큐 타입 선언

// 우선순위 큐 인터페이스 구현
func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool {
	if q[i].priority == q[j].priority {
		return q[i].id < q[j].id
	}
	return q[i].priority > q[j].priority
}
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(*process))
}
func (q *PQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

// 메모리: 12496KB
// 시간: 768ms
// 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	T, n = scanLine()
	pq := new(PQ)
	heap.Init(pq) // 우선순위 큐 초기화

	for i := 1; i <= n; i++ {
		p := scanProcess()
		heap.Push(pq, p)
	}

	// T초 반복
	for i := 1; i <= T; i++ {
		// 우선순위 큐에서 프로세스를 하나 꺼내오고
		// 꺼내온 프로세스의 아이디 출력
		x := heap.Pop(pq).(*process)
		fmt.Fprintln(writer, x.id)

		// 나머지 프로세스들의 우선순위가 1상승하는 것은 곧
		// 프로세스 x의 우선순위가 1감소하는 것과 마찬가지
		x.priority--
		// 실행시간 1감소
		x.time--

		// 프로세스 x의 실행시간이 남아있다면
		if x.time != 0 {
			heap.Push(pq, x)
		}
	}
}

func scanLine() (int, int) {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	return x, y
}

func scanProcess() *process {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	z, _ := strconv.Atoi(fields[2])
	return &process{x, y, z}
}
