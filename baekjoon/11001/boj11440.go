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
	fibo    map[int][]int
)

const MOD = 1000000007

func init() {
	fibo = make(map[int][]int)
	fibo[1] = []int{1, 1, 1, 0}
}

// 메모리: 988KB
// 시간: 4ms
// 피보나치 행렬 분할 제곱 문제

// 손으로 계산해보면서 f(2n+1) = f(n)^2 + f(n-1)^2 임을 찾았지만,
// 이 문제와는 상관이 없는 공식이었다... 헛발질했지만 나중에 어딘가에서 쓰이겠지

// 이 문제에서 사용하는 피보나치 수열의 특성은 f(n)^2 = f(n)*f(n) = f(n)*(f(n+1)-f(n-1)) = f(n)f(n+1) - f(n)f(n-1) 이다
// f(0)의 제곱부터 f(4)의 제곱의 합을 구해보면...

// f(0)^2 = 0
// f(1)^2 = f(1)f(2) - f(1)f(0) = f(1)f(2)
// f(2)^2 = f(2)f(3) - f(1)f(2)
// f(3)^2 = f(3)f(4) - f(2)f(3)
// f(4)^2 = f(4)f(5) - f(3)f(4)

// f(0)^2 + f(1)^2 + f(2)^2 + f(3)^2 + f(4)^2 = f(4)f(5)

// 즉, f(0)^2 + f(1)^2 + ... + f(n-1)^2 + f(n)^2 = f(n)*(fn+1) 이라는 것을 알 수 있다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	_ = rec(n)

	// 피보나치 행렬의 n거듭 제곱의 형태
	// f(n+1) f(n)
	// f(n)  f(n-1)
	fmt.Fprintln(writer, fibo[n][0]*fibo[n][1]%MOD)
}

func rec(x int) []int {
	f, ok := fibo[x]
	if ok {
		return f
	}

	if x%2 == 0 {
		_, ok = fibo[x/2]
		if !ok {
			fibo[x/2] = rec(x / 2)
		}
		fiboMul(x, x/2, x/2)
		return fibo[x]
	}

	_, ok = fibo[x-1]
	if !ok {
		fibo[x-1] = rec(x - 1)
	}
	fiboMul(x, x-1, 1)
	return fibo[x]
}

// 행렬 곱셈
func fiboMul(x, a, b int) []int {
	fibo[x] = make([]int, 4)
	fibo[x][0] = (fibo[a][0]*fibo[b][0] + fibo[a][1]*fibo[b][2]) % MOD
	fibo[x][1] = (fibo[a][0]*fibo[b][1] + fibo[a][1]*fibo[b][3]) % MOD
	fibo[x][2] = (fibo[a][2]*fibo[b][0] + fibo[a][3]*fibo[b][2]) % MOD
	fibo[x][3] = (fibo[a][2]*fibo[b][1] + fibo[a][3]*fibo[b][3]) % MOD
	return fibo[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
