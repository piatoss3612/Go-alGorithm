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
)

// dequeue
type DQ []int

// push x to the back
func (dq *DQ) PushBack(x int) {
	*dq = append(*dq, x)
}

// push x to the front
func (dq *DQ) PushFront(x int) {
	*dq = append(DQ{x}, *dq...)
}

// pop from the back
func (dq *DQ) PopBack() int {
	old := *dq
	n := len(old)
	x := old[n-1]
	*dq = old[0 : n-1]
	return x
}

// pop from the front
func (dq *DQ) PopFront() int {
	old := *dq
	x := old[0]
	*dq = old[1:]
	return x
}

// check value of front
func (dq DQ) Front() int {
	return dq[0]
}

// check value of back
func (dq DQ) Back() int {
	n := len(dq)
	return dq[n-1]
}

// check if dequeue is empty
func (dq DQ) Empty() bool {
	return len(dq) == 0
}

// 메모리: 182652KB
// 시간: 2448ms
// 덱(dequeue)을 사용하여 구간의 최솟값 구하기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, L := scanInt(), scanInt()
	inp := make([]int, N) // 입력값 저장

	var dq DQ // 입력값의 인덱스를 저장

	for i := 0; i < N; i++ {
		inp[i] = scanInt()

		// 덱이 비어있지 않고 덱의 가장 앞에 있는 인덱스 값이 구간의 시작(i-L+1)보다 작은 경우
		if !dq.Empty() && dq.Front() <= i-L {
			dq.PopFront() // 해당 인덱스는 필요없으므로 제거
		}

		// 덱이 비어있지 않고 덱의 가장 앞에 있는 인덱스가 가리키는 입력값이 i번째 입력값보다 큰 경우
		for !dq.Empty() && inp[dq.Back()] > inp[i] {
			dq.PopBack() // 해당 인덱스가 가리키는 값은 최솟값이 될 가능성이 없으므로 덱에서 제거
		}

		dq.PushBack(i)                              // i번째 입력값을 덱의 뒤쪽에 추가
		fmt.Fprintf(writer, "%d ", inp[dq.Front()]) // 덱의 가장 앞의 인덱스가 가리키는 값은 구간의 최솟값
	}
	fmt.Fprintln(writer)

	/*
		예제 입력:
		12 3
		1 5 2 3 6 2 3 7 3 5 2 6

		풀이 과정:

		i=0
		inp: [1] dq: [0] 구간: -2~0 최솟값: inp[0]

		i=1
		inp: [1 5] dq: [0 1] 구간: -1~1 최솟값: inp[0]

		i=2
		inp: [1 5 2] dq: [0 2] 구간: 0~2 최솟값: inp[0]

		i=3
		inp: [1 5 2 3] dq: [2 3] 구간: 1~3 최솟값: inp[2]

		i=4
		inp: [1 5 2 3 6] dq: [2 3 4] 구간: 2~4 최솟값: inp[2]

		i=5
		inp: [1 5 2 3 6 2] dq: [5] 구간: 3~5 최솟값: inp[5]

		i=6
		inp: [1 5 2 3 6 2 3] dq: [5 6] 구간: 4~6 최솟값: inp[5]

		i=7
		inp: [1 5 2 3 6 2 3 7] dq: [5 6 7] 구간: 5~7 최솟값: inp[5]

		i=8
		inp: [1 5 2 3 6 2 3 7 3] dq: [6 8] 구간: 6~8 최솟값: inp[6]

		i=9
		inp: [1 5 2 3 6 2 3 7 3 5] dq: [8 9] 구간: 7~9 최솟값: inp[8]

		i=10
		inp: [1 5 2 3 6 2 3 7 3 5 2] dq: [10] 구간: 8~10 최솟값: inp[10]

		i=11
		inp: [1 5 2 3 6 2 3 7 3 5 2 6] dq: [10 11] 구간: 9~11 최솟값: inp[10]
	*/
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

/*===================================================================================*/

// 최소힙을 사용한 방법은 시간 초과가 발생했다
func main2() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, l := scanInt(), scanInt()
	inp := &MinHeap{}
	heap.Init(inp)

	var left int
	for i := 1; i <= n; i++ {
		heap.Push(inp, &Item{scanInt(), i})

		left = i - l + 1

		var temp *Item
		for {
			temp = heap.Pop(inp).(*Item)
			if temp.idx >= left {
				fmt.Fprintf(writer, "%d ", temp.val)
				heap.Push(inp, temp)
				break
			}
		}
	}
	fmt.Fprintln(writer)
}

type Item struct {
	val, idx int
}

type MinHeap []*Item

func (m MinHeap) Len() int { return len(m) }
func (m MinHeap) Less(i, j int) bool {
	return m[i].val < m[j].val
}
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*Item))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
