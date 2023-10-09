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
	T, N    int
	machine *Machine // 캔버스를 색칠하는 머신
)

// 우선순위 큐 정의 및 구현
type Machine []int

func (m Machine) Len() int           { return len(m) }
func (m Machine) Less(i, j int) bool { return m[i] < m[j] }
func (m Machine) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *Machine) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *Machine) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 7888KB
// 시간: 92ms
// 분류: 우선순위 큐, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N = scanInt()
	machine = new(Machine)
	heap.Init(machine)
	for i := 1; i <= N; i++ {
		heap.Push(machine, scanInt())
	}
}

func Solve() {
	// 문제 지문은 이해하기 어려운데 제시된 그림을 보면 직관적으로 이해할 수 있다
	// 그림에서 페인팅 과정을 아래에서부터 위로 역순으로 살펴보면
	// 결국 13975번 문제 파일 합치기 3과 같은 로직으로 최소 잉크를 구할 수 있다는 것을 이해할 수 있다
	ans := 0
	for len(*machine) > 1 {
		n := heap.Pop(machine).(int) + heap.Pop(machine).(int)
		ans += n
		heap.Push(machine, n)
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
