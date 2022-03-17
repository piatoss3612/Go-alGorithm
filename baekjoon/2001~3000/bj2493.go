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

type Stack []int

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	var s Stack
	ans := []int{}

	/*
		레이저는 오른쪽에서 왼쪽으로 높이가 같거나 높은 탑으로만 송신된다
		스택에 인덱스를 저장하고 해당 인덱스의 탑의 높이가 i번째 탑의 높이보다 작다면
		스택에서 pop하는 것을 반복
		제거하는 과정이 끝나면 i를 스택에 push하고
		스택의 길이가 1인 경우는 0을
		아닌 경우는 스택의 길이 - 2번째로 스택에 저장된 인덱스 값을 반환한다

		예제 입력:
		5
		6 9 5 7 4

		프로세스:
		[1]: 6(인덱스 1)은 어떤 탑으로도 레이저 송신을 할 수 없다
		[2]: 9(인덱스 2)는 어떤 탑으로도 레이저 송신을 할 수 없다
		[2 3]: 5(인덱스 3)는 9(인덱스 2)로 레이저를 송신할 수 있다
		[2 4]: 7(인덱스 4)는 9(인덱스 2)로 레이저를 송신할 수 있다
		[2 4 5]: 4(인덱스 5)는 7(인덱스 4)로 레이저를 송신할 수 있다

		예제 출력:
		0 0 2 2 4
	*/

	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
		for !s.Empty() && input[s.Top()] < input[i] {
			s.Pop()
		}
		s.Push(i)
		if s.Len() == 1 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, s[s.Len()-2])
		}
	}
	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
