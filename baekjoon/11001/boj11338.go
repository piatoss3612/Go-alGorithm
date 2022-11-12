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
	T       int
	Q, K    int
)

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
	*h = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 6720KB
// 시간: 136ms
// 분류: 우선순위 큐, 자료 구조

// 문제:
// print 명령어가 입력되었을 때 insert된 정수 중에 가장 큰 수부터 최대 K개의 정수를 xor 연산한 값을 출력해야 한다
// -> print가 입력될 때마다 K개의 정수를 일일히 찾아서 xor 연산하는 것은 엄청난 시간 낭비
// 해결 방안:
// 1. xor 연산을 정수가 새롭게 입력될 때마다 누적
// 2. 1을 실행하기 위해서 입력된 정수들 중에 값이 큰 정수를 우선으로 하여 최대 K개의 정수를 우선순위 큐에 저장
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		TestCase()
	}
}

func TestCase() {
	Q, K = scanInt(), scanInt()
	h := new(Heap) // 우선순위 큐 초기화
	heap.Init(h)
	value := 0 // xor 연산을 누적한 값

	for i := 1; i <= Q; i++ {
		cmd := scanCmd()
		switch cmd {
		case "insert":
			n := scanInt()
			// 우선순위 큐의 길이가 K보다 작은 경우
			if len(*h) < K {
				heap.Push(h, n) // 입력받은 정수 n을 우선순위 큐에 추가
				value ^= n      // n을 누적하여 xor연산값 갱신
				continue
			}

			// 우선순위 큐의 길이가 K인 경우
			min := heap.Pop(h).(int) // 우선 순위 큐에서 가장 작은 값을 pop

			// 가장 작은 값 min보다 입력받은 정수 n이 큰 경우
			if n > min {
				heap.Push(h, n) // 입력받은 정수 n을 우선순위 큐에 추가
				value ^= min    // min을 누적된 xor값에서 제거
				value ^= n      // n을 누적하여 xor연산값 갱신
			} else {
				heap.Push(h, min) // n을 버리고 min을 우선순위 큐에 되돌려 놓는다
			}
		case "print":
			fmt.Fprintln(writer, value) // 누적된 xor연산값 출력
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanCmd() string {
	scanner.Scan()
	return scanner.Text()
}
