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

const P = 1000000007

// 분할 제곱 과정을 저장할 맵 초기화
func init() {
	fib = make(map[int][]int)
	fib[1] = []int{1, 1, 1, 0}
}

// 메모리: 972KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	/*
		gcd(n번째 피보나치 수, m번째 피보나치 수) = gcd(n, m)번째 피보나치 수
		참고: https://www.cut-the-knot.org/arithmetic/algebra/FibonacciGCD.shtml
	*/
	fmt.Fprintln(writer, rec(gcd(n, m))[1])
}

// 최대 공약수
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
	// gcd(n, m)번째 피보나치 수라는 것을 알고 있으므로 모듈러 연산을 사용할 수 있다
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
