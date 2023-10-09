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
	inp     []*Lecture
	alloc   []int // 강의 번호에 부여되는 강의실 번호
	N       int
)

type Lecture struct {
	number     int
	start, end int
}

type Lectures []*Lecture

func (l Lectures) Len() int { return len(l) }
func (l Lectures) Less(i, j int) bool {
	return l[i].end < l[j].end
}
func (l Lectures) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l *Lectures) Push(x interface{}) {
	*l = append(*l, x.(*Lecture))
}
func (l *Lectures) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

type MinHeap []int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

// 메모리: 11812KB
// 시간: 216ms
// 우선 순위 큐, 정렬, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]*Lecture, N)
	alloc = make([]int, N+1)
	for i := 0; i < N; i++ {
		inp[i] = &Lecture{
			number: scanInt(),
			start:  scanInt(),
			end:    scanInt(),
		}
	}

	// 강의 시작 시간 오름차순으로 정렬
	sort.Slice(inp, func(i, j int) bool {
		if inp[i].start == inp[j].start {
			return inp[i].end < inp[j].end
		}
		return inp[i].start < inp[j].start
	})

	count := 0
	pq := new(Lectures)
	room := new(MinHeap)
	heap.Init(pq)
	heap.Init(room)

	for i := 0; i < N; i++ {
		// 사용하고 있던 강의실 비우기
		for pq.Len() > 0 {
			l := heap.Pop(pq).(*Lecture)
			if l.end <= inp[i].start {
				heap.Push(room, alloc[l.number])
			} else {
				heap.Push(pq, l)
				break
			}
		}

		// 사용 가능한 강의실이 있다면
		if room.Len() > 0 {
			roomNo := heap.Pop(room).(int)
			alloc[inp[i].number] = roomNo
			heap.Push(pq, inp[i])
		} else {
			// 없다면 강의실 개수 증가
			count++
			alloc[inp[i].number] = count
			heap.Push(pq, inp[i])
		}
	}

	fmt.Fprintln(writer, count)
	for i := 1; i <= N; i++ {
		fmt.Fprintln(writer, alloc[i])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
