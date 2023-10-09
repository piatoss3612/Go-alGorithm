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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	primes := make([][]int, n+1)
	primes[1] = []int{2, 3, 5, 7}

	/*
		7331 소수, 733 소수, 73 소수, 7 소수
		즉, 길이가 1~7인 소수 * 10을 한 값에 1~9를 더하여 소수인지를 판별하는 것을 반복하면 된다
	*/

	for i := 2; i <= n; i++ {
		for _, prime := range primes[i-1] {
			for j := 1; j <= 9; j++ {
				tmp := prime*10 + j
				if isPrime(tmp) {
					primes[i] = append(primes[i], tmp)
				}
			}
		}
	}
	for _, prime := range primes[n] {
		fmt.Fprintln(writer, prime)
	}
}

func isPrime(n int) bool {
	check := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			check = false
			break
		}
	}
	return check
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
