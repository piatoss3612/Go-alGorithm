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
	MOD     = 10000
)

/*
문제에서 힌트를 보여줘서 너무 즐겁게 풀은 문제.

피보나치 수열을 2 by 2 행렬로 표현하면 분할/제곱을 사용하여 빠른 시간 안에 풀 수 있다.

피보나치 수열의 n번째 항을 계산하기 위한 행렬을 슬라이스로 풀어 정리하였다.

fib[1] = []int{1, 1, 1, 0}

...

fib[n] = []int{n+1번째 항, n번째 항, n번째 항, n-1번째 항} // fib[1]의 n제곱

메모리: 2252KB
시간: 8ms
*/

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	fib = make(map[int][]int) // 모든 테스트 케이스에서 행렬값을 공유할 것이므로 가장 앞에서 초기화
	fib[1] = []int{1, 1, 1, 0}

	var n int

	for {
		n = scanInt()
		if n == -1 {
			break
		}

		if n == 0 {
			fmt.Fprintln(writer, 0)
			continue
		}

		if n <= 2 {
			fmt.Fprintln(writer, 1)
			continue
		}

		ans := fibonacci(n)
		fmt.Fprintln(writer, ans[1])
	}
}

func fibonacci(x int) []int {
	_, ok := fib[x]
	if ok {
		return fib[x]
	}

	a := x / 2
	b := x - a

	_, ok = fib[a]
	if !ok {
		fib[a] = fibonacci(a)
	}

	_, ok = fib[b]
	if !ok {
		fib[b] = fibonacci(b)
	}

	fib[x] = mulMatrix(fib[a], fib[b])

	return fib[x]
}

// 행렬 a와 b를 곱한 결과 반환
func mulMatrix(a, b []int) []int {
	res := make([]int, 4)
	res[0] = (a[0]*b[0] + a[1]*b[2]) % MOD
	res[1] = (a[0]*b[1] + a[1]*b[3]) % MOD
	res[2] = (a[2]*b[0] + a[3]*b[2]) % MOD
	res[3] = (a[2]*b[1] + a[3]*b[3]) % MOD
	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
