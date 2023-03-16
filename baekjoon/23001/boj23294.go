package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, Q, C    int
	CAP        []int
	back       *DQ
	forward    *DQ
	cacheInUse int
	current    int
)

// 덱 구현
type DQ []int

func New() *DQ {
	dq := make(DQ, 0, 4000)
	return &dq
}

func (q DQ) Len() int    { return len(q) }
func (q DQ) Empty() bool { return len(q) == 0 }
func (q DQ) Front() int {
	if q.Empty() {
		return 0
	}
	return q[0]
}
func (q DQ) Back() int {
	if q.Empty() {
		return 0
	}
	n := len(q)
	return q[n-1]
}
func (q *DQ) PushFront(x int) {
	temp := make(DQ, 0, q.Len()+1)
	temp = append(temp, x)
	*q = append(temp, *q...)
}
func (q *DQ) PushBack(x int) {
	*q = append(*q, x)
}
func (q *DQ) PopFront() int {
	if q.Empty() {
		return 0
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *DQ) PopBack() int {
	if q.Empty() {
		return 0
	}
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 72196KB
// 시간: 168ms
// 분류: 구현, 덱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Result()
}

func Input() {
	N, Q, C = scanInt(), scanInt(), scanInt()
	CAP = make([]int, N+1)
	for i := 1; i <= N; i++ {
		CAP[i] = scanInt()
	}
	back = New()
	forward = New()

	for i := 1; i <= Q; i++ {
		Query()
	}
}

func Query() {
	cmd := scanCmd()
	switch cmd {
	case 'B':
		Back()
	case 'F':
		Forward()
	case 'A':
		Visit()
	case 'C':
		Compress()
	}
}

func Back() {
	if !back.Empty() {
		forward.PushFront(current)
		current = back.PopBack()
	}
}

func Forward() {
	if !forward.Empty() {
		back.PushBack(current)
		current = forward.PopFront()
	}
}

func Visit() {
	visit := scanInt()

	for !forward.Empty() {
		cacheInUse -= CAP[forward.PopFront()]
	}

	if current != 0 {
		back.PushBack(current)
	}
	current = visit
	cacheInUse += CAP[current]

	for cacheInUse > C {
		cacheInUse -= CAP[back.PopFront()]
	}
}

func Compress() {
	if back.Empty() {
		return
	}

	x := back.PopBack()
	temp := New()
	temp.PushBack(x)

	for !back.Empty() {
		y := back.PopBack()
		if x != y {
			x = y
			temp.PushFront(x)
		} else {
			cacheInUse -= CAP[y]
		}
	}

	back = temp
}

func Result() {
	fmt.Fprintln(writer, current)

	if back.Empty() {
		fmt.Fprintln(writer, -1)
	} else {
		for !back.Empty() {
			fmt.Fprintf(writer, "%d ", back.PopBack())
		}
		fmt.Fprintln(writer)
	}

	if forward.Empty() {
		fmt.Fprintln(writer, -1)
	} else {
		for !forward.Empty() {
			fmt.Fprintf(writer, "%d ", forward.PopFront())
		}
		fmt.Fprintln(writer)
	}
}

func scanCmd() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
