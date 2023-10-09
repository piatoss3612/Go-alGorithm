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
	inp     []SC
	n, m    int
)

type SC struct {
	start, stay int
}

// 메모리: 22440KB
// 시간: 320ms
// 최소 힙, 그리디 알고리즘
// 문제가 이해가 안 돼서 풀이를 참고했는데도 뭔소린지 잘 모르겠다...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	inp = make([]SC, n)
	for i := 0; i < n; i++ {
		inp[i] = SC{scanInt(), scanInt()}
	}

	// 사용 시작 시간으로 오름차순 정렬
	sort.Slice(inp, func(i, j int) bool {
		if inp[i].start == inp[j].start {
			return inp[i].stay < inp[j].stay
		}
		return inp[i].start < inp[j].start
	})

	// 최소 힙: 슈퍼컴퓨터의 사용이 끝나는 시간
	leave := new(MinHeap)
	heap.Init(leave)

	count := 0

	for i := 0; i < n; i++ {
		unlock := false

		for leave.Len() > 0 {
			x := heap.Pop(leave).(int)

			if inp[i].start < x {
				heap.Push(leave, x)
				break
			}

			if x+m < inp[i].start {
				continue
			} else {
				unlock = true
				break
			}
		}

		if unlock {
			count++
		}

		heap.Push(leave, inp[i].start+inp[i].stay)
	}

	fmt.Fprintln(writer, count)
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
