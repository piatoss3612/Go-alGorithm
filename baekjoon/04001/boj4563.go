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
)

// 난이도: Gold 5
// 메모리: 908KB
// 시간: 120ms
// 분류: 수학, 정수론, 피타고라스 정리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {

}

func Solve() {
	for {
		N = scanInt()
		if N == 0 {
			return
		}

		fmt.Fprintln(writer, getTheNumberOfHypotenuse(N))
	}
}

func getTheNumberOfHypotenuse(a int) (cnt int) {
	s := a * a

	// a^2 = (c+b)(c-b) 이므로 c+b와 c-b가 각각 a^2의 약수임을 이용하여 자연수 c의 개수를 구한다.
	// b > a 임을 주의한다.
	for i := 1; i*i <= s; i++ {
		// i가 s의 약수인 경우
		if s%i == 0 {
			b := (i - s/i) / 2
			c := (i + s/i) / 2

			if s+b*b == c*c && b > a {
				cnt++
			}

			// i가 s의 약수이므로 s/i도 s의 약수이다.
			r := s / i

			b = (r - s/r) / 2
			c = (r + s/r) / 2

			if s+b*b == c*c && b > a {
				cnt++
			}
		}
	}

	return
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
