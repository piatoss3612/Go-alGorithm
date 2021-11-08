package bj2581

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	primes := []int{}

	if m == 1 {
		m = 2
	}

	for i := m; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	if len(primes) == 0 {
		fmt.Println(-1)
		return
	}

	fmt.Printf("%d\n%d\n", sum(primes), primes[0])
}

func isPrime(n int) bool {
	p := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			p = false
			break
		}
	}
	return p
}

func sum(n []int) (result int) {
	for _, v := range n {
		result += v
	}
	return
}
