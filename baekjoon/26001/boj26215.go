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
	N       int
	snow    *Snow
)

// 최대 힙 정의 및 구현
type Snow []int

func (s Snow) Len() int { return len(s) }
func (s Snow) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s Snow) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *Snow) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *Snow) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 난이도: Silver 3
// 메모리: 932KB
// 시간: 4ms
// 분류: 그리디 알고리즘, 정렬, 구현, (최대 힙)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	snow = new(Snow)
	heap.Init(snow)
	for i := 0; i < N; i++ {
		heap.Push(snow, scanInt())
	}

	ans := 0
	for len(*snow) > 0 {
		n := heap.Pop(snow).(int)

		if len(*snow) > 0 {
			m := heap.Pop(snow).(int)
			ans += m
			n -= m
			if n > 0 {
				heap.Push(snow, n)
			}
		} else {
			ans += n
		}
	}

	if ans > 1440 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
