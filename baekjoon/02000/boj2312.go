package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	T         int
	minFactor [100001]int
)

// 난이도: Silver 3
// 메모리: 1700KB
// 시간: 4ms
// 분류: 소수 판정, 에라토스테네스의 체, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	T = scanInt()

	for i := 2; i <= 100000; i++ {
		minFactor[i] = i
	}

	for i := 2; i <= 100000; i++ {
		if minFactor[i] == i {
			for j := i * i; j <= 100000; j += i {
				if minFactor[j] == j {
					minFactor[j] = i
				}
			}
		}
	}
}

func Solve() {
	for i := 1; i <= T; i++ {
		n := scanInt()

		factors := make([]int, 0)

		for n > 1 {
			factors = append(factors, minFactor[n])
			n /= minFactor[n]
		}

		sort.Ints(factors)

		prev := 0
		cnt := 0

		for _, factor := range factors {
			if prev == factor {
				cnt++
			} else {
				if prev != 0 {
					fmt.Fprintf(writer, "%d %d\n", prev, cnt)
				}

				prev = factor
				cnt = 1
			}
		}
		fmt.Fprintf(writer, "%d %d\n", prev, cnt)
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
