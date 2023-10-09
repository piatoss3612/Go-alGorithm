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
	inp     [][]Employee // 화장실을 사용하려는 사원들
	N, M, K int
)

type Employee struct {
	index        int
	row          int
	jjam, urgent int
}

// 메모리: 14992KB
// 시간: 120ms
// 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()
	inp = make([][]Employee, M) // M개의 줄을 먼저 만든다
	for i := 0; i < N; i++ {
		index := i % M // 각 사원이 서야하는 줄
		inp[index] = append(inp[index], Employee{
			index:  i,
			row:    index,
			jjam:   scanInt(),
			urgent: scanInt(),
		})
	}

	toilet := &Toilet{}
	heap.Init(toilet)

	count := 0

	// 각 줄의 맨 앞에 있는 사원들을 우선순위 큐에 먼저 집어넣는다
	for i := 0; i < M; i++ {
		if len(inp[i]) > 0 {
			heap.Push(toilet, &inp[i][0])
			inp[i] = inp[i][1:]
		}
	}

	for toilet.Len() > 0 {
		cur := heap.Pop(toilet).(*Employee)

		// 우선순위 큐에서 꺼내온 사원이 K번째 사원인 경우
		// count를 출력하고 종료
		if cur.index == K {
			fmt.Fprintln(writer, count)
			return
		}

		count++ // count 증가
		// 우선순위 큐를 빠져나온 사원과 같은 줄에 사람이 남아있다면
		if len(inp[cur.row]) > 0 {
			heap.Push(toilet, &inp[cur.row][0])
			inp[cur.row] = inp[cur.row][1:]
		}
	}
}

type Toilet []*Employee

func (t Toilet) Len() int { return len(t) }
func (t Toilet) Less(i, j int) bool {
	// 근무 일수가 같다면
	if t[i].jjam == t[j].jjam {
		// 화장실이 급한 정도가 같다면
		if t[i].urgent == t[j].urgent {
			// 줄 번호가 적은 순서
			return t[i].row < t[j].row
		}
		// 화장실이 급한 순서
		return t[i].urgent > t[j].urgent
	}
	// 근무 일수가 많은 순서
	return t[i].jjam > t[j].jjam
}
func (t Toilet) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t *Toilet) Push(x interface{}) {
	*t = append(*t, x.(*Employee))
}
func (t *Toilet) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
