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

// 메모리: 6704KB
// 시간: 60ms
// 입력된 값들의 중앙값을 찾는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	smaller := &MaxHeap{} // 현재 중앙값보다 작은 값들을 저장, 중앙값 조정을 위해 최대 힙을 사용
	bigger := &MinHeap{}  // 현재 중앙값보다 큰 값들을 저장, 중앙값 조정을 위해 최소 힙을 사용
	heap.Init(smaller)
	heap.Init(bigger)

	median := scanInt() // 첫 번째 입력값을 중앙값으로 설정
	fmt.Fprintln(writer, median)

	for i := 2; i <= n; i++ {
		x := scanInt()
		if x >= median {
			heap.Push(bigger, x) // 입력값이 현재 중앙값보다 크거나 같은 경우, bigger 힙에 추가
		} else {
			heap.Push(smaller, x) // 입력값이 현재 중앙값보다 작은 경우, smaller 힙에 추가
		}

		sLen := smaller.Len() // 현재 smaller 길이
		bLen := bigger.Len()  // 현재 bigger 길이

		// 입력된 값의 개수가 짝수개인 경우
		if i%2 == 0 {
			// 중앙값은 중간에 있는 두 수 중에 작은 수가 되어야 한다
			// 따라서 smaller의 길이가 bigger의 길이보다 큰 경우
			if sLen > bLen {
				heap.Push(bigger, median)        // 중앙값을 bigger로 이동
				median = heap.Pop(smaller).(int) // smaller에서 Pop한 값을 median으로 설정
			}
			// bigger의 길이가 smaller의 길이보다 큰 경우는 고려하지 않아도 된다
		} else {
			// 입력된 값의 개수가 홀수개인 경우
			// 이 경우는 중앙값을 기준으로 smaller의 길이와 bigger의 길이가 같아야만 한다
			if sLen > bLen {
				heap.Push(bigger, median)
				median = heap.Pop(smaller).(int)
			} else if sLen < bLen {
				heap.Push(smaller, median)
				median = heap.Pop(bigger).(int)
			}
		}
		fmt.Fprintln(writer, median) // 현재 중앙값 출력
	}
}

// 최대 힙
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

// 최소 힙
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
