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

// 우선순위 큐 정의 및 구현
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
// 메모리: 1008KB
// 시간: 4ms
// 분류: 그리디 알고리즘, 우선순위 큐, 구현
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

	// A역에서 출발해야 하는 기차 정보 입력
	for i := 1; i <= NA; i++ {
		depart, arrive := scanTime(), scanTime()
		trains = append(trains, &Train{DepartAt: depart, TravelTime: arrive - depart, To: 'B'})
	}

	// B역에서 출발해야 하는 기차 정보 입력
	for i := 1; i <= NB; i++ {
		depart, arrive := scanTime(), scanTime()
		trains = append(trains, &Train{DepartAt: depart, TravelTime: arrive - depart, To: 'A'})
	}

	// 배차 정보를 출발해야 하는 순서로 오름차순 정렬
	sort.Slice(trains, func(i, j int) bool {
		if trains[i].DepartAt == trains[j].DepartAt {
			return trains[i].TravelTime < trains[j].TravelTime
		}
		return trains[i].DepartAt < trains[j].DepartAt
	})
}

func Solve(turn int) {
	cntA, cntB := 0, 0 // A역에서 출발시켜야 하는 차량 수, B역에서 출발시켜야 하는 차량 수
	fromA := new(Trains)
	fromB := new(Trains)
	heap.Init(fromA)
	heap.Init(fromB)

	for len(trains) > 0 {
		curr := trains[0] // 배차 순서가 된 기차
		trains = trains[1:]

		// 이동해야 하는 역에 따라 분기 처리
		switch curr.To {
		// 1. A역으로 이동해야 하는 경우
		case 'A':
			// B역에서 대기중인 차량이 없는 경우
			if len(*fromB) == 0 {
				// 새로운 차량 출발
				cntB++
				curr.DepartAt += curr.TravelTime + T
				curr.To = 'B'
				heap.Push(fromA, curr)
				continue
			}

			trainAtB := heap.Pop(fromB).(*Train)
			// B역에 대기중인 차량이 아직 준비중이거나 역에 도착하지 않은 경우
			if trainAtB.DepartAt > curr.DepartAt {
				cntB++
				heap.Push(fromB, trainAtB)
			}

			curr.DepartAt += curr.TravelTime + T
			curr.To = 'B'
			heap.Push(fromA, curr)

		// 2. B역으로 이동해야 하는 경우
		case 'B':
			// A역에서 대기중인 차량이 없는 경우
			if len(*fromA) == 0 {
				// 새로운 차량 출발
				cntA++
				curr.DepartAt += curr.TravelTime + T
				curr.To = 'A'
				heap.Push(fromB, curr)
				continue
			}

			trainAtA := heap.Pop(fromA).(*Train)
			// A역에 대기중인 차량이 아직 준비중이거나 역에 도착하지 않은 경우
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
