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
	fib     map[int]*[]int
	MOD     = 1000000007
)

// 메모리: 980KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	fib = make(map[int]*[]int)
	fib[1] = &[]int{1, 1, 1, 0}

	ans := 0

	/*
		# 0번째 부터 n번째 피보나치 수 중에서 홀수 번째 피보나치 수의 합

		1. n이 짝수인 경우: n번째 피보나치 수
		2. n이 홀수인 경우: n+1번째 피보나치 수

		검증:

		n = 8
		fib(8) = fib(7) + fib(6)
		= fib(7) + fib(5) + fib(4)
		= fib(7) + fib(5) + fib(3) + fib(2)
		= fib(7) + fib(5) + fib(3) + fib(1) + fib(0)
		= fib(7) + fib(5) + fib(3) + fib(1)

		따라서, 0번째 부터 8번째 피보나치 수 중에서 홀수 번째 피보나치 수의 합 ans는
		fib(8) = fib(7) + fib(5) + fib(3) + fib(1)이 된다
	*/

	if n%2 != 0 {
		n += 1
	}

	ans = (*rec(n))[1]
	fmt.Fprintln(writer, ans)
}

func rec(x int) *[]int {
	if x == 1 {
		return fib[1]
	}

	f, ok := fib[x]
	if ok {
		return f
	}

	if x%2 == 0 {
		_, ok = fib[x/2]
		if !ok {
			fib[x/2] = rec(x / 2)
		}
		fib[x] = fibMul(fib[x/2], fib[x/2])
		return fib[x]
	}

	_, ok = fib[x-1]
	if !ok {
		fib[x-1] = rec(x - 1)
	}
	fib[x] = fibMul(fib[x-1], fib[1])
	return fib[x]
}

func fibMul(a, b *[]int) *[]int {
	mul := make([]int, 4)
	mul[0] = ((*a)[0]*(*b)[0] + (*a)[1]*(*b)[2]) % MOD
	mul[1] = ((*a)[0]*(*b)[1] + (*a)[1]*(*b)[3]) % MOD
	mul[2] = ((*a)[2]*(*b)[0] + (*a)[3]*(*b)[2]) % MOD
	mul[3] = ((*a)[2]*(*b)[1] + (*a)[3]*(*b)[3]) % MOD
	return &mul
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
