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
)

// 강연 정보
type Speech struct {
	reward, deadline int
	index            int
}

// 우선순위 큐
type PQ []*Speech

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	if pq[i].reward == pq[j].reward {
		return pq[i].deadline < pq[j].deadline
	}
	return pq[i].reward < pq[j].reward
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Speech))
}
func (t *PQ) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

// 메모리: 1704KB
// 시간: 12ms
// 우선순위 큐를 사용하여 가장 돈을 많이 벌 수 있는 최적의 강연들을 선택하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp := make([]*Speech, n)
	for i := 0; i < n; i++ {
		inp[i] = &Speech{scanInt(), scanInt(), i}
	}

	// 데드라인이 짧은 강연을 우선으로 정렬하고
	// 데드라인이 같은 경우, reward가 더 큰 강연이 앞으로 오게 한다
	sort.Slice(inp, func(i, j int) bool {
		if inp[i].deadline == inp[j].deadline {
			return inp[i].reward > inp[j].reward
		}
		return inp[i].deadline < inp[j].deadline
	})

	res := &PQ{}

	// 정렬된 입력값을 앞에서부터 순회하면서
	// 가장 돈을 많이 벌 수 있는 최적의 강연들을 선택하는 것은
	// 보상이 적은 강연은 버리고 데드라인이 가장 긴 강연의 데드라인보다
	// 작거나 같은 수의 가능한한 많은 강연을 선택하는 것이다
	for len(inp) > 0 {
		x := inp[0]
		inp = inp[1:]

		// 선택한 강연들의 수가 x의 데드라인보다 작은 경우
		if len(*res) < x.deadline {
			heap.Push(res, x)
		} else if len(*res) == x.deadline { // 선택한 강연들의 수가 x의 데드라인과 같은 경우
			// 선택한 강연들 중 보상이 가장 적은 강연,
			// 가장 적은 보상이 여러 개인 경우 데드라인이 더 짧은 강연을 선택해 꺼내온다
			y := heap.Pop(res).(*Speech)
			if y.reward < x.reward { // 꺼내온 y의 보상이 x의 보상보다 적다면
				heap.Push(res, x) // x를 선택한다
			} else {
				heap.Push(res, y) // y를 선택한 강연에 되돌려 놓는다
			}
		}
	}

	total := 0

	for len(*res) > 0 {
		total += heap.Pop(res).(*Speech).reward
	}
	fmt.Fprintln(writer, total)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
