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

type PQ []int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// 메모리: 630628KB -> 20304KB
// 시간: 2484ms -> 128ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	k, n := scanInt(), scanInt()
	inp := make([]int, k)
	// check := make(map[int]bool)

	for i := 0; i < k; i++ {
		inp[i] = scanInt()
		// check[inp[i]] = true
	}

	pq := make(PQ, k)
	copy(pq, inp)
	heap.Init(&pq)

	for cnt := 0; cnt != n; {
		x := heap.Pop(&pq).(int)
		cnt += 1

		if cnt == n {
			fmt.Fprintln(writer, x)
			return
		}

		// var temp int
		for i := 0; i < k; i++ {
			// temp = x * inp[i]
			// if !check[temp] {
			// 	check[temp] = true
			// 	heap.Push(&pq, temp)
			// }
			heap.Push(&pq, x*inp[i])
			if x%inp[i] == 0 {
				break
			}
		}
	}

	/*
		예제 입력:
		4 19
		2 3 5 7

		풀이 과정:

		x=0 cnt=0
		pq: 2 3 5 7

		x=2 cnt=1
		pq: 3 4 5 7

		x=3 cnt=2
		pq: 4 5 6 7 9

		...

		x=27 cnt=19
		pq: 28 30 32 35 36 40 42 45 48 49 50 63 75 125

		예제 출력:
		27
	*/
	/*
		소수들의 곱을 저장하는데 중복된 값을 피하기 위해 맵을 사용하고
		메모리나 시간 초과날까 전전긍긍하고 있었지만 다행히 첫시도에 통과했다

		그러나 더 빠르고 간단한 중복을 피하는 방법을 찾았으니...

		if x%inp[i] == 0 {
			break
		}

		왜 x를 i번째 소수로 나눈 나머지가 0이면 곱하기를 그만두고 다음 수로 넘어가는 것일까?

		소수 2,3,5를 각각 곱한 값을 테이블로 표현하면 이렇다

		   2  3  5
		2  4  6 10
		3  6  9 15
		5 10 15 25

		우리는 여기서 대각선을 기준으로 같은 값이 대칭되어 있음을 알 수 있다

		   2  3  5
		2     6 10
		3  6    15
		5 10 15

		여기서 x가 i번째 소수로 나누어 떨어지면 곱셈을 그만둠으로써 중복되는 수를 피할 수 있다

		   2  3  5
		2  4
		3  6  9
		5 10 15 25

		pq에 다른 수가 추가되어도 마찬가지이다

		   2  3  5
		2  4
		3  6  9
		4  8
		5 10 15 25
		6 12
		9 18 27
	*/
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
