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
	MAXBUF  = 500001
)

type Stack []byte

func (s *Stack) Push(n byte) {
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

func (s *Stack) Top() byte {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, MAXBUF), MAXBUF) // 입력값의 길이가 50만이므로 버퍼의 용량을 늘려줘야 한다
	n, k := scanInt(), scanInt()
	s := scanBytes()

	var ans Stack

	for i := 0; i < n; i++ {
		// k개의 숫자를 지워 가장 큰 수를 얻기 위해서
		// 가장 높은 자릿수에 있는 수가 커야 한다

		// 따라서 스택의 마지막 숫자보다 큰 수가 들어온다면
		// k가 0보다 클 동안, 그보다 작은 수들을 스택에서 제거한다
		// 그리고 k를 1만큼 줄인다
		for k > 0 && !ans.Empty() && ans.Top() < s[i] {
			ans.Pop()
			k -= 1
		}
		ans.Push(s[i])
	}

	// 혹시 k가 남아있을 경우
	for k > 0 {
		ans.Pop()
		k -= 1
	}

	for _, v := range ans {
		fmt.Fprint(writer, string(v))
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
