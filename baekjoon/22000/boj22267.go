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
	h       *MinHeap
)

// 정수 슬라이스 타입의 최소 힙 정의
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 난이도: Gold 5
// 메모리: 6700KB
// 시간: 92ms
// 분류: 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	h = new(MinHeap)
	heap.Init(h)

	// 최소 힙에는 사탕의 개수를 2의 거듭제곱으로 표현했을 때의 차수 x가 들어간다
	for i := 1; i <= N; i++ {
		x := scanInt()
		heap.Push(h, x)
	}
}

// 2의 거듭제곱을 사용해 2의 거듭제곱을 만들려면
// 같은 차수 a를 가진 2개의 수를 병합하여 차수가 a+1인 2의 거듭제곱을 만들면 된다
// 2^0 + 2^0 = 2^1
// 2^1 + 2^1 = 2^2
// 2^2 + 2^2 = 2^3
// ...
func Solve() {
	if N == 1 {
		fmt.Fprintln(writer, "N")
		return
	}

	ex := []int{} // 같은 거듭제곱 차수가 없는 사탕들을 담는 박스

	for len(*h) > 2 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int) // a <= b
		if a == b {
			// a와 b가 같은 경우: 2^a + 2^a = 2^(a+1)이 되므로 a와 b를 제거하고 a+1을 힙에 추가
			heap.Push(h, a+1)
		} else {
			// a와 b가 다른 경우: a < b
			// b와는 같은 수가 있을 수 있으므로 b를 최소 힙으로 되돌리고 a는 ex로 분리
			ex = append(ex, a)
			heap.Push(h, b)

			// 따로 빼둔 박스의 개수가 2개 이상이 되면
			// 따로 빼둔 박스들의 거듭제곱 차수는 절대 같을 수 없기 때문에 병합할 수 없고
			// 밥과 찰리가 동시에 2의 거듭제곱 개의 사탕을 가질 수 있는 경우는 0이 된다
			// 따라서 반복문 종료
			if len(ex) > 1 {
				break
			}
		}
	}

	ans := false

	// 밥과 찰리가 동시에 2의 거듭제곱 개의 사탕을 가질 수 있는 경우 판별
	if len(ex) == 1 && len(*h) == 2 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		if a == b {
			// a와 b가 같으면 하나의 박스로 c로 병합
			// c와 ex 두 개의 박스로 조건 만족
			ans = true
		}
	} else if len(ex) == 0 && len(*h) == 2 {
		// ex에 사탕이 담겨있지 않고
		// 최소 힙에 2개의 사탕 박스가 남아있으면 그 두 개의 사탕박스를 밥과 찰리가 나눠가지면 조건 만족
		ans = true
	}

	if ans {
		fmt.Fprintln(writer, "Y")
	} else {
		fmt.Fprintln(writer, "N")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
