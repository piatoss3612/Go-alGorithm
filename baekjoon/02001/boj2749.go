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
	MOD     = 1000000
)

// 7677번과 동일
// 메모리: 1004KB
// 시간: 4ms
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

	fib = make(map[int][]int)
	fib[1] = []int{1, 1, 1, 0}

	ans := fibonacci(n)
	fmt.Fprintln(writer, ans[1])
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
