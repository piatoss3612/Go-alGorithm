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
	test := make([]int, n+1)
	for i := 1; i <= n; i++ {
		test[i] = i
	}

	// 에라토스테네스의 체
	for i := 2; i <= n; i++ {
		if test[i] == 0 {
			continue
		}
		for j := i * 2; j <= n; j += i {
			test[j] = 0
		}
	}
	primes := []int{}

	for i := 2; i <= n; i++ {
		if test[i] != 0 {
			primes = append(primes, i)
		}
	}

	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primes); j++ {
			for k := 0; k < len(primes); k++ {
				for l := 0; l < len(primes); l++ {
					tmp := primes[i] + primes[j] + primes[k] + primes[l]
					if tmp > n {
						break
					}
					if tmp == n {
						fmt.Fprintf(writer, "%d %d %d %d\n", primes[i], primes[j], primes[k], primes[l])
						return
					}
				}
			}
		}
	}
	fmt.Fprintln(writer, -1)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
