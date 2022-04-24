package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// 최솟값 우선 순위 큐
type MinPQ []int

func (h MinPQ) Len() int { return len(h) }
func (h MinPQ) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinPQ) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinPQ) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinPQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 최댓값 우선 순위 큐
type MaxPQ []int

func (h MaxPQ) Len() int { return len(h) }
func (h MaxPQ) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxPQ) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxPQ) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxPQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 메모리: 257676KB
// 시간: 1684ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	var m1 MinPQ
	var m2 MaxPQ

	var size int               // 전체 입력 결과로 남는 배열의 크기
	check := make(map[int]int) // 전체 입력 결과로 남아있는 입력값들을 체크

	heap.Init(&m1)
	heap.Init(&m2)

	k := scanInt()
	for i := 1; i <= k; i++ {
		op, d := scanString(), scanInt()
		if op == "I" {
			// 삽입 연산, 두 개의 큐에 모두 삽입해준다
			heap.Push(&m1, d)
			heap.Push(&m2, d)
			check[d] += 1 // d가 여러 번 중복으로 입력될 수 도 있으므로 1씩 늘려준다
			size += 1     // 결과 배열의 사이즈를 1 늘려준다
		} else {
			// 삭제 연산
			// 결과 배열의 크기가 0이면 무시
			if size == 0 {
				continue
			}
			// 최댓값 삭제
			if d == 1 {
				for m2.Len() > 0 {
					tmp := heap.Pop(&m2).(int)
					// check[tmp]가 0이 아니라는 것은
					// tmp가 결과 배열에 남아있는 수라는 뜻이므로
					// check[tmp] 값을 1만큼 줄이고 루프를 종료한다
					// 0인 경우는 이미 존재하지 않는 수이므로 해당 값을
					// 제거하고 다음 최댓값이 결과 배열에 존재하는지 체크한다
					if check[tmp] != 0 {
						check[tmp] -= 1
						break
					}
				}
			} else {
				for m1.Len() > 0 {
					tmp := heap.Pop(&m1).(int)
					// 최댓값을 탐색하는 경우와 마찬가지로 최솟값을 탐색한다
					if check[tmp] != 0 {
						check[tmp] -= 1
						break
					}
				}
			}
			size -= 1
		}
	}

	if size == 0 {
		fmt.Fprintln(writer, "EMPTY")
	} else if size == 1 {
		max := heap.Pop(&m2).(int)
		for m2.Len() > 0 {
			if check[max] != 0 {
				check[max] -= 1
				break
			} else {
				max = heap.Pop(&m2).(int)
			}
		}
		fmt.Fprintf(writer, "%d %d\n", max, max)
	} else {
		max := heap.Pop(&m2).(int)
		// max 값이 결과 배열에 남아있는 값인지 체크
		for m2.Len() > 0 {
			if check[max] != 0 {
				check[max] -= 1
				break
			} else {
				max = heap.Pop(&m2).(int)
			}
		}
		min := heap.Pop(&m1).(int)
		// min 값이 결과 배열에 남아있는 값인지 체크
		for m1.Len() > 0 {
			if check[min] != 0 {
				check[min] -= 1
				break
			} else {
				min = heap.Pop(&m1).(int)
			}
		}
		fmt.Fprintf(writer, "%d %d\n", max, min)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
