package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
이 거지같은 문제를 나는 왜 7번이나 틀렸는가?
그것은 바로 입력으로 들어오는 배열의 수의 개수만 최대 10만 개이기 때문이다...
그럼 콤마랑 대괄호랑 합치면 20만은 훌쩍 넘을 수 있다는 소리?
따라서 bufio.Scanner의 버퍼 크기를 늘려주었다...

Go 언어는 참고할만한 풀이가 적어서 머리가 아프다 정말...

참고 - https://www.acmicpc.net/board/view/69300
*/

const MaxBuf int = 400000

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		testCase()
	}
}

func testCase() {
	ops, n := scanString(), scanInt()
	numbers := scanNumbers()
	isReversed := false // 뒤집으라고 진짜 뒤집으면 시간 초과난다. 뒤집혀 있는지 아닌지만 체크.

	for _, op := range ops {
		switch op {
		case 'R':
			{
				isReversed = !isReversed
			}
		case 'D':
			{
				if n == 0 {
					fmt.Fprintln(writer, "error")
					return
				} else {
					if isReversed {
						numbers = numbers[:n-1] // 뒤집혀있을 경우 마지막이 곧 첫 번째 수
					} else {
						numbers = numbers[1:]
					}
				}
				n -= 1 // 버리기 연산이 끝나면 n을 1감소시킨다
			}
		}
	}

	if isReversed {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	fmt.Fprintln(writer, "["+strings.Join(numbers, ",")+"]")
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanNumbers() []string {
	numbers := strings.Split(strings.Trim(scanString(), "[]"), ",")
	return numbers
}
