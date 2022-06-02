package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n       int
)

// 메모리: 892KB
// 시간: 4ms
// 오일러 피 함수: n보다 작으면서 n과 서로소인 양의 정수의 수를 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		n = scanInt()
		// n이 0이면 프로그램 종료
		if n == 0 {
			return
		}

		// n이 1이면 n보다 작은 양의 정수가 존재하지 않으므로 0을 출력
		if n == 1 {
			fmt.Fprintln(writer, 0)
			continue
		}

		// 오일러 피 함수 실행
		EulerPhi(n)
	}
}

// 오일러 피 함수
// 양의 정수 n보다 작으면서 n과 서로소인 수의 개수: n * (1-1/p1) * (1-1/p2) * ... * (1-1/pn-1) * (1-1/pn)
// p1, p2, ..., pn-1, pn은 n의 소인수
func EulerPhi(x int) {
	res := x

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			res = res * (i - 1) / i
			for x%i == 0 {
				x /= i
			}
		}
	}

	//
	if x != 1 {
		res = res * (x - 1) / x
	}

	fmt.Fprintln(writer, res)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
