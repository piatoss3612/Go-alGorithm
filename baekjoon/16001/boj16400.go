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
	MOD     = 123456789
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	primes := getPrimes()

	dp := make([]int, n+1)
	dp[0] = 1

	for _, prime := range primes {
		if prime > n {
			break
		}
		for i := prime; i <= n; i++ {
			// i는 i-prime을 만들 수 있는 조합에 prime을 더한 것이므로
			// i-prime을 만들 수 있는 조합의 수를 누적해서 더해준다
			dp[i] = (dp[i] + dp[i-prime]) % MOD
		}
	}
	fmt.Fprintln(writer, dp[n])
}

// 에라토스테네스의 체
func getPrimes() []int {
	primes := []int{}
	nums := make([]int, 40001)
	for i := 2; i <= 40000; i++ {
		if nums[i] == 0 {
			primes = append(primes, i)
			for j := i; j <= 40000; j += i {
				nums[j] = 1
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
