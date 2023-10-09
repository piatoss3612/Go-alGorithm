package bj11502

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
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		primes := getPrimes(n)
		hasAnswer := false
	Loop:
		for i := 0; i < len(primes); i++ {
			for j := 0; j < len(primes); j++ {
				for k := 0; k < len(primes); k++ {
					tmp := primes[i] + primes[j] + primes[k]
					if tmp == n {
						hasAnswer = true
						fmt.Fprintln(writer, primes[i], primes[j], primes[k])
						break Loop
					}
				}
			}
		}
		if !hasAnswer {
			fmt.Fprintln(writer, 0)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getPrimes(n int) []int {
	primes := []int{}
	checked := make([]int, n)
	for i := 2; i < n; i++ {
		if checked[i] == 0 {
			primes = append(primes, i)
			checked[i] = 1
		}
		for j := i + i; j < n; j += i {
			if checked[j] == 0 {
				checked[j] = 1
			}
		}
	}
	return primes
}
