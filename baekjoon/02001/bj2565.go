package bj2565

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
)

type Rope struct {
	l, r int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]Rope, n+1)
	for i := 1; i <= n; i++ {
		input[i] = Rope{scanInt(), scanInt()}
	}

	// 전깃줄의 왼쪽 위치의 오름차순으로 정렬
	sort.Slice(input, func(i, j int) bool {
		return input[i].l < input[j].l
	})

	// dp: 교차하지 않는 상태의 전깃줄이 최대가 되는 경우
	dp := make([]int, n+1)
	dp[1] = 1

	/*
		dp:
		1. {1 8} - 1
		2. {2 2} - 1
		3. {1 8} {3 9} - 2
		4. {4 1} - 1
		5. {4 1} {6 4} - 2
		6. {4 1} {6 4} {7 6} - 3
		7. {4 1} {6 4} {7 6} {9 7} - 4
		8. {4 1} {6 4} {7 6} {9 7} {10 10} - 5
	*/

	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 1; j-- {
			// 전깃줄의 오른쪽이 더 작은 값을 찾아서 dp 갱신
			if input[j].r < input[i].r {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
	}

	sort.Ints(dp)
	// 전체 전깃줄 n에서 교차하지 않는 경우의 최댓값을 빼면 결과를 얻을 수 있다
	fmt.Fprintln(writer, n-dp[n])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
