package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	min, max *MaxD        // 가장 쉬운 문제순, 가장 어려운 문제순으로 문제를 추천해주는 우선순위 큐
	toSolve  map[int]bool // 풀어야 하는 문제인지 판단
)

// 우선순위 큐 정의
type MaxD []*Problem

func (d MaxD) Len() int { return len(d) }
func (d MaxD) Less(i, j int) bool {
	if d[i].difficulty == d[j].difficulty {
		return d[i].number > d[j].number
	}
	return d[i].difficulty > d[j].difficulty
}
func (d MaxD) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d *MaxD) Push(x interface{}) {
	*d = append(*d, x.(*Problem))
}
func (d *MaxD) Pop() interface{} {
	old := *d
	n := len(old)
	x := old[n-1]
	*d = old[:n-1]
	return x
}

type Problem struct {
	number     int
	difficulty int
}

// 난이도: Gold 4
// 메모리: 17608KB
// 시간: 112ms
// 분류: 우선순위 큐, 트리를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Query()
}

func Input() {
	N = scanInt()
	min, max = new(MaxD), new(MaxD)
	heap.Init(min)
	heap.Init(max)
	toSolve = make(map[int]bool)

	for i := 1; i <= N; i++ {
		addProblem()
	}
}

func Query() {
	M = scanInt()
	for i := 1; i <= M; i++ {
		q := scanQuery()
		switch q {
		case "add":
			// 문제를 추가할 때도 solved 처리된 문제들을 각각의 큐에서 미리 제거해줘야 한다
			for len(*min) > 0 {
				p := heap.Pop(min).(*Problem)
				if toSolve[-p.number] {
					heap.Push(min, p)
					break
				}
			}

			for len(*max) > 0 {
				p := heap.Pop(max).(*Problem)
				if toSolve[p.number] {
					heap.Push(max, p)
					break
				}
			}

			addProblem()

		case "recommend":
			x := scanInt()
			if x == 1 {
				// 가장 어려운 문제 추천
				for len(*max) > 0 {
					p := heap.Pop(max).(*Problem)
					if toSolve[p.number] {
						fmt.Fprintln(writer, p.number)
						heap.Push(max, p)
						break
					}
				}
			} else {
				// 가장 쉬운 문제 추천
				for len(*min) > 0 {
					p := heap.Pop(min).(*Problem)
					if toSolve[-p.number] {
						fmt.Fprintln(writer, -p.number)
						heap.Push(min, p)
						break
					}
				}
			}
		case "solved":
			toSolve[scanInt()] = false
		}
	}
}

func addProblem() {
	p, l := scanInt(), scanInt()
	heap.Push(min, &Problem{-p, -l})
	heap.Push(max, &Problem{p, l})
	toSolve[p] = true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanQuery() string {
	scanner.Scan()
	return scanner.Text()
}
