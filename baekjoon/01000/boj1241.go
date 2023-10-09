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

	N       int
	numbers [100001]int
	cnt     [1000001]int
)

// 난이도: Gold 5
// 메모리: 10616KB
// 시간: 384ms
// 분류: 수학, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		numbers[i] = scanInt()
		cnt[numbers[i]]++
	}
}

func Solve() {
	// O(N*sqrt(N))
	for i := 1; i <= N; i++ {
		total := 0
		for j := 1; j*j <= numbers[i]; j++ {
			// numbers[i]가 j로 나누어 떨어질 경우
			if numbers[i]%j == 0 {
				total += cnt[j]
				// numbers[i]가 j*j로 표현할 수 있는 제곱수가 아닐 경우
				if j*j != numbers[i] {
					total += cnt[numbers[i]/j]
				}
			}
		}
		fmt.Fprintln(writer, total-1) // 자기 자신 제외
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
