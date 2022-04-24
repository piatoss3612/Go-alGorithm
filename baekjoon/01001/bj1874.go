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

type Stack []int // 스택 타입을 정수 슬라이스로 선언

func (s *Stack) push(n int) { // 스택에 push
	*s = append(*s, n)
}

func (s *Stack) pop() { // 스택에서 pop
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) empty() bool { // 스택이 비어있는지 확인
	return len(*s) == 0
}

func (s *Stack) back() int { // 스택 입구에서 정수 가져오기
	return (*s)[len(*s)-1]
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	target := make([]int, n)
	idx := 0
	result := make([]string, 0)
	var s Stack
	for i := 0; i < n; i++ {
		target[i] = scanInt()
	}
	/*
		입력:
		8
		4
		3
		6
		8
		7
		5
		2
		1

		과정:
		s = []Stack{1} push
		s = []Stack{1, 2} push
		s = []Stack{1, 2, 3} push
		s = []Stack{1, 2, 3, 4} push
		s = []Stack{1, 2, 3} pop
		s = []Stack{1, 2} pop
		s = []Stack{1, 2, 5} push
		s = []Stack{1, 2, 5, 6} push
		s = []Stack{1, 2, 5} pop
		s = []Stack{1, 2, 5, 7} push
		s = []Stack{1, 2, 5, 7, 8} push
		s = []Stack{1, 2, 5, 7} pop
		s = []Stack{1, 2, 5} pop
		s = []Stack{1, 2} pop
		s = []Stack{1} pop
		s = []Stack{} pop
	*/
	for i := 1; i <= n; i++ {
		s.push(i)
		result = append(result, "+")
		for !s.empty() && s.back() == target[idx] {
			s.pop()
			result = append(result, "-")
			idx += 1
		}
	}

	if !s.empty() {
		fmt.Fprintln(writer, "NO")
	} else {
		for _, v := range result {
			fmt.Fprintln(writer, v)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
