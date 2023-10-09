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
	a, b := scanInt(), scanInt()

	// 1천만 이상 1억 이하의 팰린드롬이면서 소수인 수는 없으므로 범위를 좁힌다
	if b > 10000000 {
		b = 10000000
	}

	for i := a; i <= b; i++ {
		if isPalindromeAndPrime(i) {
			fmt.Fprintln(writer, i)
		}
	}
	fmt.Fprintln(writer, -1)
}

func isPalindromeAndPrime(n int) bool {
	// 팰린드롬 체크
	s := strconv.Itoa(n)
	mid := 0
	if len(s)%2 == 0 {
		mid = len(s) / 2
	} else {
		mid = (len(s) + 1) / 2
	}
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	// 소수 체크
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
