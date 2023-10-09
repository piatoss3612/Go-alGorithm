package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	notPrime [5000000]bool // 소수 판별
	checked  [5000000]bool // 이미 등장한 소수 판별
)

// 에라토스테네스의 체를 사용해 소수 판별
func init() {
	eratosthenes()
}

// 메모리: 20416KB
// 시간: 88ms
// 에라토스테네스의 체, 최소 힙
// 규칙 설명 중에 2번째 규칙 중에 '상대방은 지금까지 상대방이 말한 소수중에서'
// 라는 표현이 해석에 혼란을 불러와 조금 헤맨 문제
// 이 부분은 규성이가 소수가 아닌 수를 부를 경우,
// 대웅이는 대웅이가 부른 소수들 중에 3번째로 큰 수만큼의 점수를 얻는다는 의미이다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	dscore, gscore := 0, 0
	// 최소 힙의 길이를 3으로 제한하면 최솟값을 꺼내올 경우, 3번째로 큰 수를 가져올 수 있다
	dheap, gheap := &Heap{}, &Heap{}
	heap.Init(dheap)
	heap.Init(gheap)

	for i := 1; i <= N; i++ {
		d, g := scanInt(), scanInt()

		// 1. 대웅이부터 시작
		// 1-1. 대웅이가 부른 값이 소수가 아닌 경우
		if notPrime[d] {
			// 규성이가 말한 소수의 개수가 3개 미만이라면
			if len(*gheap) < 3 {
				gscore += 1000
			} else {
				// 규성이가 말한 소수 중에 3번째로 큰 수를 점수에 추가
				gmin := heap.Pop(gheap).(int)
				gscore += gmin
				heap.Push(gheap, gmin)
			}
			// 1-2. 대웅이가 부른 값이 이미 등장한 소수인 경우
		} else if checked[d] {
			dscore -= 1000
			// 1-3. 대웅이가 부른 값이 아직 등장하지 않은 소수인 경우
		} else {
			if len(*dheap) < 3 {
				heap.Push(dheap, d)
			} else {
				// 최소 힙의 길이가 3 이상이라면
				// 최소 힙의 최솟값 갱신
				dmin := heap.Pop(dheap).(int)
				if dmin >= d {
					heap.Push(dheap, dmin)
				} else {
					heap.Push(dheap, d)
				}
			}

			checked[d] = true // 처음 등장한 소수 d를 체크
		}

		// 2. 규성이의 차례
		// 2-1. 규성이가 부른 값이 소수가 아닌 경우
		if notPrime[g] {
			if len(*dheap) < 3 {
				dscore += 1000
			} else {
				dmin := heap.Pop(dheap).(int)
				dscore += dmin
				heap.Push(dheap, dmin)
			}
			// 2-2. 규성이가 부른 값이 이미 등장한 소수인 경우
		} else if checked[g] {
			gscore -= 1000
			// 2-3. 규성이가 부른 값이 아직 등장하지 않은 소수인 경우
		} else {
			if len(*gheap) < 3 {
				heap.Push(gheap, g)
			} else {
				gmin := heap.Pop(gheap).(int)
				if gmin >= g {
					heap.Push(gheap, gmin)
				} else {
					heap.Push(gheap, g)
				}
			}

			checked[g] = true
		}
	}

	// 대웅이와 규성이의 점수 비교
	if dscore > gscore {
		fmt.Fprintln(writer, "소수의 신 갓대웅")
	} else if dscore < gscore {
		fmt.Fprintln(writer, "소수 마스터 갓규성")
	} else {
		fmt.Fprintln(writer, "우열을 가릴 수 없음")
	}
}

func eratosthenes() {
	notPrime[0] = true
	notPrime[1] = true
	for i := 2; i*i < 5000000; i++ {
		if !notPrime[i] {
			for j := i * i; j < 5000000; j += i {
				notPrime[j] = true
			}
		}
	}
}

// 최소 힙
type Heap []int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
