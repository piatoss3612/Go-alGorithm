package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	fib     map[int][]int
	MOD     = 1000000007
)

// 피보나치 행렬 분할 / 거듭제곱 문제

// 메모리:
// 1044KB -> 1004KB
// map을 초기화 할 때, 전체 연산에 필요한 log2N만큼의 크기로 초기화

// 시간:
// 8ms -> 4ms
// _, ok := fib[x]으로 2번 확인하는 대신 바로 fib[x] 값이 nil인지 확인

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	if n == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	if n <= 2 {
		fmt.Fprintln(writer, 1)
		return
	}

	fib = make(map[int][]int, int(math.Log2(float64(n))))
	fib[1] = []int{1, 1, 1, 0}

	ans := fibonacci(n)
	fmt.Fprintln(writer, ans[1])
}

func fibonacci(x int) []int {
	if fib[x] != nil {
		return fib[x]
	}

	a := x / 2
	b := x - a

	if fib[a] == nil {
		fib[a] = fibonacci(a)
	}

	if fib[b] == nil {
		fib[b] = fibonacci(b)
	}

	fib[x] = mulMatrix(fib[a], fib[b])

	return fib[x]
}

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
