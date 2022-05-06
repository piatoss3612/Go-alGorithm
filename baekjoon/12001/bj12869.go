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
	msk     [4]int          // 뮤탈리스크 1, 2, 3의 체력
	dp      [61][61][61]int // 뮤탈리스크 1,2,3의 체력 변화를 메모이제이션하는 배열
)

// 메모리: 4168KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for i := 1; i <= n; i++ {
		msk[i] = scanInt()
	}
	fmt.Fprintln(writer, rec(msk[1], msk[2], msk[3]))
}

func rec(a, b, c int) int {
	// 기저 사례: 모든 뮤탈리스크가 파괴된 경우
	if a <= 0 && b <= 0 && c <= 0 {
		return 0
	}

	// 뮤탈리스크의 체력이 0보다 작으면 인덱스로 인한 런타임에러가 발생하므로 0으로 되돌려 준다
	if a < 0 {
		a = 0
	}

	if b < 0 {
		b = 0
	}

	if c < 0 {
		c = 0
	}

	ret := &dp[a][b][c]

	if *ret != 0 {
		return *ret
	}

	// 최솟값 비교를 위해 최댓값처럼 보이는 수를 할당
	*ret = 1000

	// 뮤탈리스크가 모두 파괴되는 최소 공격 횟수를 찾는다
	if a > 0 {
		*ret = min(*ret, min(rec(a-9, b-3, c-1), rec(a-9, b-1, c-3))+1)
	}

	if b > 0 {
		*ret = min(*ret, min(rec(a-3, b-9, c-1), rec(a-1, b-9, c-3))+1)
	}

	if c > 0 {
		*ret = min(*ret, min(rec(a-3, b-1, c-9), rec(a-1, b-3, c-9))+1)
	}

	return *ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
