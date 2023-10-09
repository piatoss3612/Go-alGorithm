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
	N, M    int
	wok     []int      // 웍의 크기
	dp      [10001]int // dp[i]: 짜장면 1개를 만들기 위해 필요한 요리 횟수의 최솟값
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 1136KB
// 시간: 124ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	wok = make([]int, M)

	// dp값 INF로 초기화
	// i개의 짜장면을 요리하기 위해 정확히 i개의 그릇을 요리하는 웍을
	// 선택하는 경우의 최솟값 비교를 위해 dp[0]는 0으로 남겨둔다
	for i := 1; i <= N; i++ {
		dp[i] = INF
	}

	// M개의 웍의 크기 입력
	for i := 0; i < M; i++ {
		wok[i] = scanInt()
	}

	// M개의 웍 중에 순서에 상관없이 2개의 웍을 선택했을 경우의 값을 wok 슬라이스에 추가
	for i := 0; i < M-1; i++ {
		for j := i + 1; j < M; j++ {
			// 웍을 조합한 값은 N보다 작거나 같아야 한다
			if wok[i]+wok[j] <= N {
				wok = append(wok, wok[i]+wok[j])
			}
		}
	}

	sort.Ints(wok) // wok 슬라이스를 오름차순으로 정렬

	for i := 1; i <= N; i++ {
		for j := 0; j < len(wok); j++ {
			// j번째 웍의 조합으로 요리할 수 있는 그릇의 수가 i보다 작거나 같은 경우
			if wok[j] <= i {
				// dp[i]의 값은 dp[i]와 (i-wok[j]만큼의 그릇을 요리했을 때의 요리 횟수의 최솟값+1)의 최솟값을 비교한 값
				dp[i] = min(dp[i], dp[i-wok[j]]+1)
			} else {
				break
			}
		}
	}

	if dp[N] == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[N])
	}
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
