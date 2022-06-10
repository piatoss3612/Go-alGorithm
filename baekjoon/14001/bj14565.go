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

// 메모리: 896KB
// 시간: 4ms
// 확장 유클리드 호제법을 사용하여 모듈러 곱셈 역원을 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, A := scanInt(), scanInt()
	addInverse := N - A // 모듈러 덧셈 역원

	// N과 A가 서로소가 아니라면
	// 모듈러 N에 대한 A의 곱셈의 역원을 구할 수 없다
	if GCD(N, A) != 1 {
		fmt.Fprintln(writer, addInverse, -1)
		return
	}

	mulInverse := ExEuclidean(N, A) // 모듈러 곱셈 역원
	fmt.Fprintln(writer, addInverse, mulInverse)
}

// 유클리드 호제법
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// 확장 유클리드 호제법
// ax + by = 1을 만족하는 y를 구한다
func ExEuclidean(a, b int) int {
	n := a
	var q, r int
	var t1, t2, T int = 0, 1, 0

	for b > 0 {
		q = a / b
		r = a % b
		a = b
		b = r
		T = t1 - t2*q
		t1 = t2
		t2 = T
	}
	return (t1 + n) % n // c는 양의 정수여야 하므로, 음수인 경우 모듈러 덧셈 역원을 반환
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
