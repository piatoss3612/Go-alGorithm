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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	primes := getPrimes(n)

	if n == 1 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 두 포인터 알고리즘
	l, r := 0, 0
	sum := primes[0]
	ans := 0
	for l <= r && r < len(primes) {
		if sum == n {
			ans += 1
			sum -= primes[l]
			l += 1
		} else if sum > n {
			sum -= primes[l]
			l += 1
		} else {
			r += 1
			if r == len(primes) {
				break
			}
			sum += primes[r]
		}
	}

	fmt.Fprintln(writer, ans)
}

// 에라토스테네스의 체
func getPrimes(n int) []int {
	check := make([]int, n+1)
	primes := []int{}

	for i := 2; i <= n; i++ {
		if check[i] == 0 {
			primes = append(primes, i)
			for j := i; j <= n; j += i {
				check[j] = 1
			}
		}
	}
	return primes
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
