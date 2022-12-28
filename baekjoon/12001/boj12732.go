package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N         int
	T, NA, NB int
	trains    []*Train
)

type Train struct {
	DepartAt   int
	TravelTime int
	To         rune
}

type Trains []*Train

func (t Trains) Len() int { return len(t) }
func (t Trains) Less(i, j int) bool {
	if t[i].DepartAt == t[j].DepartAt {
		return t[i].TravelTime < t[j].TravelTime
	}
	return t[i].DepartAt < t[j].DepartAt
}
func (t Trains) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t *Trains) Push(x interface{}) {
	*t = append(*t, x.(*Train))
}
func (t *Trains) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[:n-1]
	return x
}

// 난이도: Gold 5
// 메모리: 1896KB
// 시간: 12ms
// 분류: 그리디 알고리즘, 우선순위 큐, 구현
// 12731번 재탕
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		Input()
		Solve(i)
	}
}

func Input() {
	T, NA, NB = scanInt(), scanInt(), scanInt()
	trains = make([]*Train, 0, NA+NB)

	for i := 1; i <= NA; i++ {
		depart, arrive := scanTime(), scanTime()
		trains = append(trains, &Train{DepartAt: depart, TravelTime: arrive - depart, To: 'B'})
	}

	for i := 1; i <= NB; i++ {
		depart, arrive := scanTime(), scanTime()
		trains = append(trains, &Train{DepartAt: depart, TravelTime: arrive - depart, To: 'A'})
	}

	sort.Slice(trains, func(i, j int) bool {
		if trains[i].DepartAt == trains[j].DepartAt {
			return trains[i].TravelTime < trains[j].TravelTime
		}
		return trains[i].DepartAt < trains[j].DepartAt
	})
}

func Solve(turn int) {
	cntA, cntB := 0, 0
	fromA := new(Trains)
	fromB := new(Trains)
	heap.Init(fromA)
	heap.Init(fromB)

	for len(trains) > 0 {
		curr := trains[0]
		trains = trains[1:]

		switch curr.To {
		case 'A':
			if len(*fromB) == 0 {
				cntB++
				curr.DepartAt += curr.TravelTime + T
				curr.To = 'B'
				heap.Push(fromA, curr)
				continue
			}

			trainAtB := heap.Pop(fromB).(*Train)
			if trainAtB.DepartAt > curr.DepartAt {
				cntB++
				heap.Push(fromB, trainAtB)
			}

			curr.DepartAt += curr.TravelTime + T
			curr.To = 'B'
			heap.Push(fromA, curr)
		case 'B':
			if len(*fromA) == 0 {
				cntA++
				curr.DepartAt += curr.TravelTime + T
				curr.To = 'A'
				heap.Push(fromB, curr)
				continue
			}

			trainAtA := heap.Pop(fromA).(*Train)
			if trainAtA.DepartAt > curr.DepartAt {
				cntA++
				heap.Push(fromA, trainAtA)
			}

			curr.DepartAt += curr.TravelTime + T
			curr.To = 'A'
			heap.Push(fromB, curr)
		}
	}

	fmt.Fprintf(writer, "Case #%d: %d %d\n", turn, cntA, cntB)
}

func scanTime() int {
	scanner.Scan()
	fields := strings.Split(scanner.Text(), ":")
	h, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	return h*60 + m
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
