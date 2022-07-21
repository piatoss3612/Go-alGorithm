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
	var sum int

	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}

	/*
		예제 입력:
		6
		10
		3
		7
		4
		12
		2

		프로세스:
		해당 건물 옥상에서 볼 수 있는 옥상의 수 보다
		해당 옥상을 볼 수 있는 옥상의 수를 스택으로 파악해야 한다

		10을 볼 수 있는 건물:
		3을 볼 수 있는 건물: 10
		7을 볼 수 있는 건물: 10
		4를 볼 수 있는 건물: 10, 7
		12를 볼 수 있는 건물:
		2를 볼 수 있는 건물: 12
		다해서 5라는 결과가 나와야 한다


	*/

	for i := 0; i < n; i++ {
		// input[i] 옥상을 볼 수 없는 옥상은 스택에서 pop
		for !s.Empty() && s.Top() <= input[i] {
			s.Pop()
		}
		s.Push(input[i])
		// input[i] 옥상을 볼 수 있는 건물들과 input[i]만 남게 되므로
		// 결과에 (스택의 길이 - 1)을 더해준다
		sum += s.Len() - 1
	}
	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
