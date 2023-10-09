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
	k, n    int
	dp      [64][200]int // i초에 물벼룩의 위치 j, j는 물벼룩의 위치는 n보다 충분히 큰 수여야 한다
)

// 메모리: 1140KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	k, n = scanInt(), scanInt()

	ans := solve(0, k)
	fmt.Fprintln(writer, ans)
}

func solve(sec, pos int) int {
	// 기저 사례1: 물벼룩이 잡아먹히는 경우
	if pos == 0 {
		return 0
	}

	// 기저 사례2: 물벼룩이 생존한 경우
	if sec == n {
		return 1
	}

	ret := &dp[sec][pos]
	if *ret != 0 {
		return *ret
	}

	// 수면 위로 1cm 이동했을 때 생존할 경우의 수 + 수면 아래로 1cm 이동했을 때 생존할 경우의 수
	*ret += solve(sec+1, pos-1) + solve(sec+1, pos+1)

	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
