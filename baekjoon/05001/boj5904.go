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

	N int
	S []int
)

// 난이도: Gold 5
// 메모리: 912KB
// 시간: 4ms
// 분류: 분할 정복, 재귀
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = append(S, 3)
	i := 1
	for {
		next := S[i-1]*2 + i + 3
		S = append(S, next)
		i += 1
		// 문자열의 길이가 10^9를 넘어가는 경우를 마지막으로 추가
		if next >= 1000000000 {
			break
		}
	}

	N = scanInt()
}

func Solve() {
	idx := 0

	// N이 속한 구간을 찾는다.
	for i := 0; i < len(S); i++ {
		if N <= S[i] {
			idx = i
			break
		}
	}

	// N이 속한 구간에서 문자열을 찾는다.
	fmt.Fprintln(writer, divideAndConquer(idx, N))
}

func divideAndConquer(idx, n int) string {
	if idx == 0 {
		if n == 1 {
			return "m"
		} else {
			return "o"
		}
	}

	
	if n <= S[idx-1] { // S[idx-1]: 왼쪽 구간에 속하는 경우
		return divideAndConquer(idx-1, n)
	} else if n <= S[idx-1]+idx+3 { // S[idx-1]+idx+3: 중간 구간에 속하는 경우
		if n == S[idx-1]+1 {
			return "m"
		} else {
			return "o"
		}
	} else { // 오른쪽 구간에 속하는 경우
		return divideAndConquer(idx-1, n-S[idx-1]-idx-3)
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
