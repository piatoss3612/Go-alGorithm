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
	M, N, S int
	d, n    int // 모든 주사위의 기댓값의 합을 기약 분수 형태로 나타냈을 때의 분모와 분자
)

const MOD = 1000000007

// 난이도: Gold 4
// 메모리: 1224KB
// 시간: 8ms
// 분류: 정수론, 모듈로 곱셈 역원
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	M = scanInt()
	for i := 1; i <= M; i++ {
		N, S = scanInt(), scanInt()
		if i == 1 {
			n, d = N, S
		} else {
			// 새로 입력된 주사위의 기댓값을 더한 분모, 분자값 갱신
			// 값이 너무 커질 수 있으므로 모듈로 연산은 필수
			nn, nd := (n*N)%MOD, (d*N+n*S)%MOD
			n, d = nn, nd
		}
	}
}

func Solve() {
	inverse := ExEuclidean(n)             // 분모 n의 모듈로 곱셈 역원 구하기
	fmt.Fprintln(writer, (d*inverse)%MOD) // d*n^-1 MOD 1000000007
}

// b: 모듈로 MOD에 대한 곱셈 역원을 구하고자 하는 대상
// 확장 유클리드 알고리즘으로 모듈로 MOD에 대한 b의 곱셈 역원을 구한다
func ExEuclidean(b int) int {
	a := MOD
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

	return (t1 + MOD) % MOD
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
