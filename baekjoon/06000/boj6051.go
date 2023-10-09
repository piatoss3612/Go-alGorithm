package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
func (s *Stack) Pop() int {
	if s.Empty() {
		return 0
	}
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
func (s *Stack) Top() int {
	if s.Empty() {
		return -1
	}
	n := len(*s)
	return (*s)[n-1]
}
func (s *Stack) Empty() bool {
	n := len(*s)
	if n == 0 {
		return true
	}
	return false
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	backup  map[int]Stack // 1~N번째 쿼리의 결괏값을 저장할 해시맵
)

// 메모리: 23368KB
// 시간: 108ms
// 스택, 슬라이스값 복사
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	backup = make(map[int]Stack)
	backup[0] = Stack{} // 혹시 모르니까 확실하게 해둡니다

	for i := 1; i <= N; i++ {
		op := scanQuery() // 쿼리 읽기
		// 슬라이스는 참조값으로 사용되므로
		// 반복문 내에서 새로운 스택 슬라이스를 생성하고
		// 이전 쿼리의 결과 슬라이스를 완전히 복사하는 식으로 문제풀이 진행
		var s Stack

		// 쿼리에 따른 분기 처리
		switch op {
		// 쿼리가 'a'인 경우
		case 'a':
			// s를 이전에 처리한 쿼리의 결괏값만큼의 길이로 초기화
			s = make(Stack, len(backup[i-1]))
			copy(s, backup[i-1]) // 슬라이스의 값을 복사
			problem := scanInt()
			s.Push(problem) // 새로운 문제 번호 추가
		// 쿼리가 's'인 경우
		case 's':
			// s를 이전에 처리한 쿼리의 결괏값만큼의 길이로 초기화
			s = make(Stack, len(backup[i-1]))
			copy(s, backup[i-1]) // 슬라이스의 값을 복사
			s.Pop()              // 마지막으로 푼 문제 번호 제거

		// 쿼리가 't'인 경우
		case 't':
			queryTo := scanInt()                    // 되돌아갈 쿼리 번호 읽기
			s = make(Stack, len(backup[queryTo-1])) // 되돌아갈 쿼리 번호 이전의 결괏값의 슬라이스 길이만큼 s 초기화
			copy(s, backup[queryTo-1])              // 슬라이의 값을 복사
		}

		backup[i] = s                 // 쿼리 결괏값 슬라이스 저장
		fmt.Fprintln(writer, s.Top()) // i번째 쿼리 결괏값의 마지막 문제 번호 출력
	}
}

func scanQuery() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
