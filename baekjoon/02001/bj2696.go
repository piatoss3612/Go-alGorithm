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

// 메모리: 1780KB
// 시간: 8ms
// 홀수 번째 입력 순서일 때의 중앙값 구하기
// 1655번 문제와 동일한 해법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		testCase()
	}
}

// 테스트 케이스
func testCase() {
	n := scanInt()

	smaller := &MaxHeap{}
	bigger := &MinHeap{}
	heap.Init(smaller)
	heap.Init(bigger)

	ans := make([]int, 0)

	median := scanInt()
	ans = append(ans, median)
	for i := 2; i <= n; i++ {
		x := scanInt()
		if x >= median {
			heap.Push(bigger, x)
		} else {
			heap.Push(smaller, x)
		}

		sLen := smaller.Len()
		bLen := bigger.Len()

		if i%2 == 0 {
			if sLen > bLen {
				heap.Push(bigger, median)
				median = heap.Pop(smaller).(int)
			}
		} else {
			if sLen > bLen {
				heap.Push(bigger, median)
				median = heap.Pop(smaller).(int)
			} else if sLen < bLen {
				heap.Push(smaller, median)
				median = heap.Pop(bigger).(int)
			}

			ans = append(ans, median) // 홀수 번째 입력일 때의 중앙값을 ans 슬라이스에 추가
		}
	}

	fmt.Fprintln(writer, len(ans))

	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(writer, "%d ", ans[i])
		if i%10 == 9 { // 한 줄에 10개의 숫자만 출력
			fmt.Fprintln(writer)
		}
	}
	fmt.Fprintln(writer)
}

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

type MinHeap []int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
