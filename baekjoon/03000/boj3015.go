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
)

// 첫시도 시간 초과
func FirstTry() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp := make([]int, 0, n)
	ans := 0

	var x int
	for i := 0; i < n; i++ {
		x = scanInt()

		for len(inp) > 0 && inp[len(inp)-1] < x {
			inp = inp[:len(inp)-1]
			ans += 1
		}

		// 정답은 맞는 것 같지만 이부분에서 시간 초과가 발생하는 것으로 보인다
		for j := len(inp) - 1; j >= 0; j-- {
			if inp[j] != x {
				ans += 1
				break
			}
			ans += 1
		}

		inp = append(inp, x)
	}
	fmt.Fprintln(writer, ans)
}

type Item struct {
	height int // 신장
	pairs  int // 같은 신장이면서 마주볼 수 있는 사람의 수
}

// 스택 구현
type Stack []*Item

func (s *Stack) Push(x *Item) {
	*s = append(*s, x)
}

func (s *Stack) Pop() *Item {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func (s *Stack) Top() *Item {
	n := len(*s)
	return (*s)[n-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

// 메모리: 17364KB
// 시간: 144ms
// 오아시스 재결합 기원합니다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp := make(Stack, 0, n) // 스택 초기화
	ans := 0

	var x int
	for i := 0; i < n; i++ {
		x = scanInt()
		cnt := 0 // 같은 신장이면서 마주볼 수 있는 사람의 수

		// inp가 비어있지 않고 inp의 마지막에 있는 사람의 신장이 i번째 사람의 신장 x보다 작은 경우
		for !inp.Empty() && inp.Top().height < x {
			// i번째 사람과 마주볼 수는 있지만 앞으로 쓸모가 없으므로 스택에서 제거
			inp.Pop()
			ans += 1
		}

		// inp가 비어있지 않고 inp의 마지막에 있는 사람의 신장이 i번째 사람의 신장 x와 같은 경우
		if !inp.Empty() && inp.Top().height == x {
			cnt = inp.Top().pairs + 1 // i번째 입력값의 pairs 갱신
			ans += cnt                // cnt만큼 마주볼 수 있는 사람의 수 추가
		}

		// 신장이 같은 사람의 수보다 inp의 길이가 크면 반드시 x보다 큰 사람과 한 번 마주볼 수 있다
		if len(inp) > cnt {
			ans += 1
		}

		inp.Push(&Item{x, cnt}) // i번째 입력값 스택에 추가
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
