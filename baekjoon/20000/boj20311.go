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
	N, K    int
)

// 문제 조건: 모든 이웃한 시험관 쌍에 대해, 두 시험관에 들어 있는 시약의 색깔이 서로 다르다

type Test struct {
	index int // 시험관의 종류
	count int // 시험관 종류에 해당하는 개수
}

// 우선순위 큐 정의 및 인터페이스 구현
type Tests []*Test

func (c Tests) Len() int { return len(c) }

// 그리디 알고리즘: 조건을 만족하는 최적해는 시험관의 개수가 가장 큰 것을 기준으로 내림차순 정렬함으로써 찾을 수 있다
func (c Tests) Less(i, j int) bool {
	return c[i].count > c[j].count
}
func (c Tests) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Tests) Push(x interface{}) {
	*c = append(*c, x.(*Test))
}
func (c *Tests) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 메모리: 18216KB
// 시간: 256ms
// 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	tests := &Tests{}
	heap.Init(tests)

	// 시험관 정보 입력
	for i := 1; i <= K; i++ {
		test := Test{
			index: i,
			count: scanInt(),
		}
		heap.Push(tests, &test)
	}

	ans := make([]int, 0, N) // 문제 조건에 맞는 해답
	var prev int             // 이전에 선택한 시험관 정보

	for len(*tests) > 0 {
		x := heap.Pop(tests).(*Test)

		// 1. ans의 길이가 0이거나 우선순위 큐에서 꺼내온 시험관의 종류가 이전에 선택한 것과 다른 경우
		if len(ans) == 0 || x.index != prev {
			ans = append(ans, x.index)
			prev = x.index
			x.count--
			// 시험관 x의 개수가 0이 아니라면 우선순위 큐로 되돌린다
			if x.count != 0 {
				heap.Push(tests, x)
			}
			continue
		}

		// 2. 우선순위 큐에서 꺼내온 시험관의 종류가 이전에 선택한 것과 동일한 경우
		if len(*tests) > 0 {
			// 2-1. 우선순위 큐에 다른 시험관이 남아있는 경우
			// x 대신 우선순위 큐에서 다른 시험관 y를 꺼내와 배열하면 문제 조건을 만족시킬 수 있다
			y := heap.Pop(tests).(*Test)
			ans = append(ans, y.index)
			prev = y.index
			y.count--
			// 시험관 y의 개수가 0이 아니라면 우선순위 큐로 되돌린다
			if y.count != 0 {
				heap.Push(tests, y)
			}
			heap.Push(tests, x) // 시험관 x는 사용하지 않았으므로 우선순위 큐로 되돌린다
		} else {
			// 2-2. 우선순위 큐에 다른 시험관이 남아있지 않은 경우
			// 문제 조건에 해당하는 해답을 찾을 수 없으므로 -1을 출력하고 프로그램 종료
			fmt.Fprintln(writer, -1)
			return
		}
	}

	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
