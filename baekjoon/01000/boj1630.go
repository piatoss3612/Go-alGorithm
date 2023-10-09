package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	n         int
	minFactor [1000001]int
)

const MOD = 987654321

// 난이도: Gold 4
// 메모리: 32084KB
// 시간: 400ms
// 분류: 에라토스테네스의 체
// 1부터 n까지 모든 수의 최소 공배수 구하기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	n = scanInt()
}

func Solve() {
	for i := 1; i <= n; i++ {
		minFactor[i] = i
	}

	for i := 2; i*i <= n; i++ {
		if minFactor[i] == i {
			for j := i * i; j <= n; j += i {
				if minFactor[j] == j {
					minFactor[j] = i
				}
			}
		}
	}

	lcmFactors := map[int]int{}

	for i := 2; i <= n; i++ {
		if minFactor[i] == i {
			lcmFactors[i]++
			continue
		}

		primeFactors := map[int]int{}

		temp := i
		for temp > 1 {
			primeFactors[minFactor[temp]]++
			temp /= minFactor[temp]
		}

		for k, v := range primeFactors {
			if lcmFactors[k] < v {
				lcmFactors[k] = v
			}
		}
	}

	ans := 1

	for k, v := range lcmFactors {
		for i := v; i > 0; i-- {
			ans *= k
			ans %= MOD
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
