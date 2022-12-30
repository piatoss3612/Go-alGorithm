package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner       = bufio.NewScanner(os.Stdin)
	writer        = bufio.NewWriter(os.Stdout)
	humblePrimes  = []int{2, 3, 5, 7} // 험블수의 소인수인 소수들
	humbleNumbers [5843]int           // 1번부터 5842번 험블수를 저장하는 배열
)

// 우선순위 큐 정의 및 구현
type HumbleNumber []int

func (h HumbleNumber) Len() int { return len(h) }
func (h HumbleNumber) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h HumbleNumber) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *HumbleNumber) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *HumbleNumber) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 난이도: Gold 3
// 메모리: 1760KB
// 시간: 20ms
// 분류: 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	initHumbleNumbers()
	Scan()
}

// 전처리로 5842개의 험블수 구하기
func initHumbleNumbers() {
	h := new(HumbleNumber)
	heap.Init(h)
	heap.Push(h, 1) // 첫번째 험블수는 1

	for i := 1; i <= 5842; i++ {
		hnum := heap.Pop(h).(int)
		humbleNumbers[i] = hnum // i번째 험블수

		// i번째 험블수와 값이 동일한 수(중복)를 우선순위 큐에서 제거
		for len(*h) > 0 {
			next := heap.Pop(h).(int)

			if next != hnum {
				heap.Push(h, next)
				break
			}
		}

		// i번째 험블수와 험블수의 소인수가 되는 소수들을 각각 곱하여
		// 새로운 험블수 생성 및 우선순위 큐에 삽입
		for _, hp := range humblePrimes {
			heap.Push(h, hnum*hp)
		}
	}

	h = nil
}

func Scan() {
	for {
		n := scanInt()
		if n == 0 {
			return
		}

		// 수의 순서 표현
		var ordinalExp string

		switch n % 10 {
		case 1:
			ordinalExp = "st"
		case 2:
			ordinalExp = "nd"
		case 3:
			ordinalExp = "rd"
		default:
			ordinalExp = "th"
		}

		// 11, 12, 13으로 끝나는 수는 예외 처리
		if n%100 == 11 || n%100 == 12 || n%100 == 13 {
			ordinalExp = "th"
		}

		fmt.Fprintf(writer, "The %d%s humble number is %d.\n", n, ordinalExp, humbleNumbers[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
