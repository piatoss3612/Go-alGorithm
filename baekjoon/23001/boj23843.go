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
	N, M    int
	devices []int
)

type Sockets []int

func (s Sockets) Len() int { return len(s) }
func (s Sockets) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sockets) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *Sockets) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *Sockets) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 메모리: 1140KB
// 시간: 8ms
// 그리디 알고리즘, 최소 힙
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	devices = make([]int, N)
	for i := 0; i < N; i++ {
		devices[i] = scanInt()
	}

	// 그리디 알고리즘:
	// 충전 시간이 가장 오래걸리는 전자기기부터 먼저 충전하면 최적의 해를 찾을 수 있다
	// 따라서 devices 슬라이스를 내림차순으로 정렬
	sort.Slice(devices, func(i, j int) bool {
		return devices[i] > devices[j]
	})

	s := new(Sockets)
	heap.Init(s)

	// 비어있는 콘센트에 전자기기 연결
	for len(devices) > 0 && s.Len() < M {
		heap.Push(s, devices[0])
		devices = devices[1:]
	}

	// 충전해야 할 기기가 남아있는 경우
	for len(devices) > 0 {
		x := heap.Pop(s).(int) // x는 특정 콘센트에 연결된 전자기기가 충전될 때까지의 누적 시간
		x += devices[0]        // 대기중인 전자기기를 충전하는 시간만큼 누적 시간 증가
		devices = devices[1:]
		heap.Push(s, x) // 최소 힙에 누적 시간 push
	}

	ans := 0
	// 모든 전자기기를 충전하는데 걸리는 최소 시간은
	// 최소 힙의 최댓값, 즉 다른 모든 값을 pop하고 남은 최소 힙의 마지막 값과 같다
	for s.Len() > 0 {
		ans = heap.Pop(s).(int)
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
