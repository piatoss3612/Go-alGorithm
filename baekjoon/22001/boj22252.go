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
	N, X    int
	inp     []int
)

type Lines []int

func (l Lines) Len() int { return len(l) }
func (l Lines) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l Lines) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l *Lines) Push(x interface{}) {
	*l = append(*l, x.(int))
}
func (l *Lines) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[:n-1]
	return x
}

// 메모리: 25856KB
// 시간: 240ms
// 이분 탐색, 최소 힙
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, X = scanInt(), scanInt()
	inp = make([]int, N)
	for i := 0; i < N; i++ {
		inp[i] = scanInt()
	}

	// 1. 필요한 공정 라인의 개수를 이분 탐색으로 구한다
	l, r := 1, N
	for l <= r {
		mid := (l + r) / 2

		// 2. mid개의 공정 라인 개수로 모든 선물을 X시간 안에 제작할 수 있는지 여부
		flag := check(mid)

		// 3. 제작할 수 있다면
		if flag {
			// 공정 라인 개수를 줄이기 위해 r을 mid-1로 갱신
			r = mid - 1
		} else {
			// 제작할 수 없다면
			// 공정 라인 개수를 늘리기 위해 l을 mid+1로 갱신
			l = mid + 1
		}
	}

	// 문제의 해는 공정 라인 개수를 줄일 수 있는 상한선
	// 즉, 마지막으로 r을 갱신할 때 사용한 mid값인 r+1이 되어야 한다
	fmt.Fprintln(writer, r+1)
}

// 4. 최소 힙을 사용해 mid개의 공정 라인 개수로 X시간 내에 모든 선물을 제작할 수 있는지 판별
func check(mid int) bool {
	idx := 0

	lines := &Lines{}
	heap.Init(lines)

	// 5. mid개의 공정 라인을 먼저 가동시킨다
	for i := 1; i <= mid; i++ {
		heap.Push(lines, inp[idx])
		idx++
	}

	// 6. 가장 먼저 완료된 라인의 가동 시간 갱신
	for idx < N {
		next := heap.Pop(lines).(int)
		// 7. idx번째 선물을 제작하기 위해 필요한 시간이 X를 초과한 경우
		if next+inp[idx] > X {
			lines = nil
			// 선물 제작을 완료할 수 없으므로 false 반환
			return false
		} else {
			heap.Push(lines, next+inp[idx])
			idx++
		}
	}
	lines = nil
	// 모든 선물을 제작할 수 있으므로 true 반환
	return true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
