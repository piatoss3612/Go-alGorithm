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

// 최소 힙 구현
type FileCombiner []int

func (f FileCombiner) Len() int { return len(f) }
func (f FileCombiner) Less(i, j int) bool {
	return f[i] < f[j]

}
func (f FileCombiner) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f *FileCombiner) Push(x interface{}) {
	*f = append(*f, x.(int))
}
func (f *FileCombiner) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

// 메모리: 66496KB -> 25764KB
// 시간: 1928ms -> 1712ms
// 최소 힙, 그리디 알고리즘 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		TestCase()
	}
}

// 테스트 케이스
func TestCase() {
	k := scanInt()

	// 시간, 메모리 절약 요소: 길이는 0, 캡션을 k으로 지정하여 k만큼만 입력값을 받도록 슬라이스의 크기를 제한
	// 왜? 파일 합치는 과정에서 결국 슬라이스의 길이는 k보다 작아지면 작아졌지 늘어나지는 않을 것이므로
	fc := make(FileCombiner, 0, k)
	for i := 0; i < k; i++ {
		heap.Push(&fc, scanInt())
	}

	res := 0

	// 그리디:
	// 모든 파일을 합치는데 필요한 최소비용을 구하는 최적의 방법은
	// 가장 작은 파일을 2개씩 꺼내와 합치고 다시 최소 힙에 집어넣는 과정을 반복하는 것
	// *이 문제에서 파일을 꺼내오는데는 순서 제약이 없다*

	// 반복문: 최소 힙의 길이가 1보다 큰 경우
	for len(fc) > 1 {
		x := heap.Pop(&fc).(int)
		y := heap.Pop(&fc).(int)
		res += x + y // 파일을 합친 비용 누적
		heap.Push(&fc, x+y)
	}

	fmt.Fprintln(writer, res)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
