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
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N         int
	processes []Process
	pq        *PQ // 실행 요청받은 프로세스 대기열
)

type Process struct {
	i, t, p, b int // i: 번호, t: 실행 요청 시점, p: 초기 우선 순위, b: 실행 시간
}

type PQ []*Process

func (p PQ) Len() int { return len(p) }
func (p PQ) Less(i, j int) bool {
	// 우선순위가 같고 실행 시간이 같은 프로세스가 여러 개라면 부여된 번호가 작은 프로세스가 먼저 실행
	if p[i].p == p[j].p && p[i].b == p[j].b {
		return p[i].i < p[j].i
	}
	// 우선순위가 같은 프로세스가 여러 개라면 실행 시간이 짧은 프로세스가 먼저 실행
	if p[i].p == p[j].p {
		return p[i].b < p[j].b
	}
	// 우선순위가 가장 높은 프로세스 실행
	return p[i].p > p[j].p
}
func (p PQ) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*Process))
}
func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

// 난이도: Gold 2
// 메모리: 35448KB
// 시간: 356ms
// 분류: 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	processes = make([]Process, N)
	for i := 0; i < N; i++ {
		processes[i] = Process{i + 1, scanInt(), scanInt(), scanInt()}
	}
	// 요청 시간 오름차순으로 정렬
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].t < processes[j].t
	})

	pq = new(PQ)
	heap.Init(pq)
}

func Solve() {
	ans := make([]int, 0, N)
	timeConsumed := 0 // 스케줄러의 시간

	for len(processes) > 0 {
		// 요청 큐가 비어있는 상태인 경우
		if len(*pq) == 0 {
			x := processes[0]
			processes = processes[1:]
			heap.Push(pq, &Process{x.i, x.t, x.p - x.t, x.b})

			// *중요* 스케줄러 요청을 처리하고 다음 요청을 받기까지 공백이 존재할 수 있다
			if x.t > timeConsumed {
				timeConsumed = x.t
			}

			// 프로세스 x와 실행 요청 시간이 동일한 프로세스를 선별하여 큐에 추가
			for len(processes) > 0 {
				y := processes[0]
				if x.t == y.t {
					processes = processes[1:]
					heap.Push(pq, &Process{y.i, y.t, y.p - y.t, y.b})
				} else {
					break
				}
			}
		}

		// 요청 큐에서 조건에 맞는 프로세스를 하나 꺼내와 실행
		next := heap.Pop(pq).(*Process)
		ans = append(ans, next.i)

		timeConsumed += next.b // 프로세스 next를 실행한만큼 스케줄러의 시간 변화

		for len(processes) > 0 {
			x := processes[0]
			// 프로세스 next가 종료되는 시점을 포함하여 그전에 실행 요청을 받은 프로세스를 큐에 추가
			if x.t <= timeConsumed {
				processes = processes[1:]
				heap.Push(pq, &Process{x.i, x.t, x.p - x.t, x.b})
				// 우선순위에서 요청 시간을 뺌으로써 큐에 들어있는 다른 모든 프로세스의 우선순위를
				// 큐에 들어가서 대기한 단위 시간만큼 증가시키는 것과 같은 효과를 볼 수 있다
			} else {
				break
			}
		}
	}

	// 요청 큐에 남은 프로세스를 차례대로 실행
	for len(*pq) > 0 {
		next := heap.Pop(pq).(*Process)
		ans = append(ans, next.i)
	}

	// 결과 형식에 맞춰 프로세스의 실행 순서 출력
	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
