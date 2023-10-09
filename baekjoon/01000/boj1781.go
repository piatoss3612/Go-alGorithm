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

// 문제 정보
type Question struct {
	deadline int
	reward   int
	index    int
}

// 우선순위 큐
type Result []*Question

func (res Result) Len() int { return len(res) }

// 우선순위 큐 정렬
// 1순위: 컵라면 보상이 적은 문제
// 2순위: 데드라인이 짧은 문제
func (res Result) Less(i, j int) bool {
	if res[i].reward == res[j].reward {
		return res[i].deadline < res[j].deadline
	}
	return res[i].reward < res[j].reward
}
func (res Result) Swap(i, j int) {
	res[i], res[j] = res[j], res[i]
	res[i].index = i
	res[j].index = j
}
func (res *Result) Push(x interface{}) {
	*res = append(*res, x.(*Question))
}
func (res *Result) Pop() interface{} {
	old := *res
	n := len(old)
	x := old[n-1]
	*res = old[:n-1]
	return x
}

// 메모리: 	19448 -> 17600KB
// 시간: 372 -> 224ms
// 우선순위 큐를 사용하여 받을 수 있는 컵라면의 수의 최댓값을 구하는 그리디 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp := make([]*Question, n)
	for i := 0; i < n; i++ {
		inp[i] = &Question{scanInt(), scanInt(), i}
	}

	// 정렬 1순위: 데드라인 오름차순, 2순위: 컵라면 보상 내림차순
	sort.Slice(inp, func(i, j int) bool {
		if inp[i].deadline == inp[j].deadline {
			return inp[i].reward > inp[j].reward
		}
		return inp[i].deadline < inp[j].deadline
	})

	res := &Result{}

	for len(inp) > 0 {
		resSize := len(*res)
		x := inp[0]
		inp = inp[1:]

		// 지금까지 선택한 문제의 수가 문제 x의 데드라인보다 작다면
		if resSize < x.deadline {
			heap.Push(res, x) // 문제 x를 선택한다
		} else if resSize == x.deadline {
			// 지금까지 선택한 문제의 수가 문제 x의 데드라인과 같다면
			// y는 res에서 가장 보상이 적은 문제, 또는
			// 가장 보상이 적은 문제가 여러 개라면 데드라인이 더 짧은 문제
			y := heap.Pop(res).(*Question)
			// y의 보상이 x의 보상보다 적다면
			if y.reward < x.reward {
				heap.Push(res, x) // x를 res에 추가한다
			} else {
				heap.Push(res, y) // y를 res로 되돌린다
			}
		}
	}

	sum := 0

	for len(*res) > 0 {
		sum += heap.Pop(res).(*Question).reward
	}

	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
