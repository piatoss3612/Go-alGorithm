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
	fib     map[int][]int
)

const P = 1000000000

func init() {
	fib = make(map[int][]int)
	fib[1] = []int{1, 1, 1, 0}
}

// 메모리: 1024KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	a, b := scanInt(), scanInt()

	/*
		# 피보나치 수 a번째 항부터 b번째 항까지의 합

			예제 입력: 4 10

		1. 피보나치 수 1번째 항부터 10번째 항까지의 합 X

			fib(10) + fib(9) + fib(8) + fib(7) + fib(6) + fib(5) + fib(4) + fib(3) + fib(2) + fib(1)
			= fib(10) + fib(10) + fib(8) + fib(6) + fib(4) + fib(2)
			= fib(10) + fib(11) - fib(1)
			= fib(12) - fib(1)


		2. 피보나치 수 1번째 항부터 3번째 항까지의 합 Y

			fib(3) + fib(2) + fib(1) = fib(3) + fib(4) - fib(1) = fib(5) - fib(1)


		3. 피보나치 수 4번째 항부터 10번째 항까지의 합

			X - Y = fib(12) - fib(1) - (fib(5) - fib(1)) = fib(12) - fib(5) = 139
	*/

	// 연산 과정에서 나머지 연산으로 인해 rec(b + 2)[1] - rec(a + 1)[1]가 음수가 될 수 있다
	sum := (rec(b + 2)[1] - rec(a + 1)[1] + P) % P
	fmt.Fprintln(writer, sum)
}

// 피보나치 행렬 분할 제곱
func rec(x int) []int {
	_, ok := fib[x]
	if ok {
		return fib[x]
	}

	if x%2 != 0 {
		_, ok = fib[x-1]
		if !ok {
			fib[x-1] = rec(x - 1)
		}
		fib[x] = fibMul(fib[x-1], fib[1])
		return fib[x]
	}

	_, ok = fib[x/2]
	if !ok {
		fib[x/2] = rec(x / 2)
	}
	fib[x] = fibMul(fib[x/2], fib[x/2])
	return fib[x]
}

// 피보나치 행렬 곱셈
func fibMul(a, b []int) []int {
	res := make([]int, 4)
	res[0] = (a[0]*b[0] + a[1]*b[2]) % P
	res[1] = (a[0]*b[1] + a[1]*b[3]) % P
	res[2] = (a[2]*b[0] + a[3]*b[2]) % P
	res[3] = (a[2]*b[1] + a[3]*b[3]) % P
	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
