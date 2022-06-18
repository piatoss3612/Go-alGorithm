package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	counts  [1000001]int // i의 등장 횟수를 저장하는 슬라이스
)

// 스택 구현
type Stack []int

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func (s *Stack) Top() int {
	n := len(*s)
	return (*s)[n-1]
}

// 메모리: 	40096KB -> 	25544KB
// 시간: 	392ms -> 	432ms
// 스택 구현과 inp 슬라이스를 재사용하여 오등큰수를 저장함으로써 메모리는 줄었는데 시간은 조금 늘었다...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	inp := make([]int, n)
	for i := 0; i < n; i++ {
		x := scanInt()
		counts[x] += 1 // x의 등장 횟수 1증가
		inp[i] = x
	}

	var s Stack // 스택 사용

	// 입력값을 뒤에서부터 순회
	for i := n - 1; i >= 0; i-- {
		temp := inp[i] // inp를 재사용하여 오등큰수를 저장하기 위해 inp[i]를 캐싱

		// 스택이 비어있지 않다면
		if !s.Empty() {
			for !s.Empty() {
				// 스택의 top에 있는 수의 등장횟수가 temp의 등장 횟수보다 작거나 같다면
				// 이 과정을 통해 temp 오른쪽에 있는 수들 중에서
				// temp의 등장 횟수보다 크면서 가장 왼쪽에 있는 수를 찾을 수 있다
				if counts[s.Top()] <= counts[temp] {
					s.Pop() // 스택에서 제거
				} else {
					break // 아닌 경우는 반복문 탈출
				}
			}
		}

		// 스택이 비어있지 않다면
		if !s.Empty() {
			inp[i] = s.Top() // temp의 오등큰수는 스택의 top에 있는 수
		} else {
			inp[i] = -1 // 스택이 비어있다면 temp는 등장 횟수가 가장 큰 수
		}

		s.Push(temp) // temp를 스택에 push
	}

	// 결과 출력
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", inp[i])
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
