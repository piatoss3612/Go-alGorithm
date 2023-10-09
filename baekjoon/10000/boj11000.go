package bj11000

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
)

type Class struct {
	start int
	end   int
}

type PQ []Class

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	if pq[i].start == pq[j].start {
		return pq[i].end < pq[j].end
	}
	return pq[i].start < pq[j].start
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Class))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

type TPQ []Class

func (tpq TPQ) Len() int           { return len(tpq) }
func (tpq TPQ) Less(i, j int) bool { return tpq[i].end < tpq[j].end }
func (tpq TPQ) Swap(i, j int)      { tpq[i], tpq[j] = tpq[j], tpq[i] }

func (tpq *TPQ) Push(x interface{}) {
	*tpq = append(*tpq, x.(Class))
}

func (tpq *TPQ) Pop() interface{} {
	old := *tpq
	n := len(old)
	x := old[n-1]
	*tpq = old[0 : n-1]
	return x
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	time := &PQ{} // 회의 시작 시간 오름차 순으로 정렬된 우선순위 큐
	heap.Init(time)
	for i := 0; i < n; i++ {
		heap.Push(time, Class{scanInt(), scanInt()})
	}

	// 회의가 끝나는 시간이 오름차 순으로 정렬될 우선순위 큐
	class := &TPQ{}
	heap.Init(class)

	for time.Len() > 0 {
		t := heap.Pop(time).(Class)

		if class.Len() == 0 {
			heap.Push(class, t)
			continue
		}

		s := heap.Pop(class).(Class)
		// class 우선순위 큐의 회의 종료 시간이 가장 작은 값과
		// time 우선순위 큐에서 꺼내온 Class 구조체의 회의 시작 시간을 비교
		if t.start >= s.end {
			heap.Push(class, t)
		} else {
			heap.Push(class, s)
			heap.Push(class, t)
		}
	}

	fmt.Fprintln(writer, class.Len())
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
