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

type Stack []rune

func (s *Stack) Push(n rune) {
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

func (s *Stack) Top() rune {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	S := Stack((scanString()))

	ans := recursive(0, &S)
	fmt.Fprintln(writer, ans)
}

/*
7번을 틀리고 https://moons-memo.tistory.com/171를 참고하여 풀었습니다

입력: 3(4(5)4)

5 -> 1
4(5) -> 4

4(5)4 -> 5
3(4(5)4) -> 15

출력: 15

입력: 5(5(52(3)33(2)))

2 -> 1
3(2) -> 3
33(2) -> 4

3 -> 1
2(3) -> 2
52(3) -> 3

52(3)33(2) -> 7
5(52(3)33(2)) -> 35
5(5(52(3)33(2))) -> 175

출력: 175
*/

func recursive(temp int, s *Stack) int {
	for s.Len() > 0 {
		top := s.Top()
		s.Pop()
		// ')'인 경우 재귀 호출
		if top == ')' {
			temp += recursive(0, s)
			// '('인 경우 K(Q)를 풀어낸 값을 반환
		} else if top == '(' {
			n, _ := strconv.Atoi(string(s.Top()))
			temp *= n
			s.Pop()
			return temp
			// 정수인 경우 길이에 +1
		} else {
			temp += 1
		}
	}
	return temp
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
