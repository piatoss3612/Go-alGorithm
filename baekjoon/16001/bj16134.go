package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	factorial [1000001]int
)

const P = 1000000007

// 팩토리얼 수 전처리
func init() {
	factorial[0], factorial[1] = 1, 1
	for i := 2; i <= 1000000; i++ {
		factorial[i] = (factorial[i-1] * i) % P
	}
}

// 메모리: 8728KB
// 시간: 16ms
// 페르마의 소정리, 모듈러 곱셈 역원(확장된 유클리드 알고리즘)을 사용해 이항 계수를 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, r := scanInt(), scanInt()

	// nCr^P ≡ nCr (mod P) (P는 소수이며 nCr과 P는 서로소)
	// (n!/(n-r)!r!)^P ≡ nCr (mod P)
	// (n-r)!r!의 모듈러 곱셈 역원 = ((n-r)!r!)^P-2
	// n!*(n(n-r)!r!)^P-2 ≡ nCr (mod P)
	ans := factorial[n] % P
	ans *= ExEuclidean(P, factorial[n-r])
	ans %= P
	ans *= ExEuclidean(P, factorial[r])
	ans %= P
	fmt.Fprintln(writer, ans)
}

// 확장된 유클리드 알고리즘을 사용해 모듈러 곱셈 역원을 구하는 함수
// a와 b가 서로소(GCD(a,b)=1)이고 a>=b인 경우
// a*s + b*t = 1을 만족하는 s와 t를 찾을 수 있는데
// 여기서 t는 b에 대한 mod a 연산의 곱셈 역원
/*
	a = 26, b = 11인 경우

	1. a = 26, b = 11, q = 2, r = 4, t1 = 0, t2 = 1, t = -2
	2. a = 11, b = 4, q = 2, r = 3, t1 = 1, t2 = -2, t = 5
	3. a = 4, b = 3, q = 1, r = 1, t1 = -2, t2 = 5, t = -7
	4. a = 3, b = 1, q = 3, r = 0, t1 = 5, t2 = -7, t = 26

	11*t ≡ 1 (mod 26)이 되는 11의 모듈러 곱셈 역원은 t2 = -7
	7의 모듈러 덧셈 역원 19 또한 11의 모듈러 곱셈 역원

*/
func ExEuclidean(a, b int) int {
	var q, r int
	var t1, t2, t int = 0, 0, 1
	for b > 0 {
		q = a / b
		r = a % b
		a = b
		b = r
		t1 = t2
		t2 = t
		t = t1 - t2*q
	}
	return (P + t2) % P // t2가 음수인 경우를 고려하여 t2 mod P의 덧셈의 역원을 구해준다
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
