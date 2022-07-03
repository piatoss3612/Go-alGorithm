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
	meal    [301][301]int // 기내식의 맛...
	dp      [301][301]int // i번 도시를 j번째에 들렀을 때 기내식 점수를 메모이제이션
	N, M, K int
)

// 전처리: 현재 접근할 수 없는 도시를 -1로 표시
func init() {
	for i := 0; i <= 300; i++ {
		for j := 0; j <= 300; j++ {
			dp[i][j] = -1
		}
	}
	dp[1][1] = 0 // 도시 1은 반드시 1번째로 들려야 하는 도시
}

// 메모리: 4124KB
// 시간: 44ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()

	for i := 1; i <= K; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		meal[a][b] = max(meal[a][b], c) // a->b 경로가 여러 개가 있을 수 있으므로 기내식 점수의 최댓값을 갱신
	}

	// 여행은 반드시 1번 도시에서 시작
	for i := 1; i <= N; i++ {
		// 반드시 현재 머물고 있는 도시 번호보다 큰 번호를 가진 도시로 이동
		for j := i + 1; j <= N; j++ {
			// i에서 j로 가는 경로가 있다면
			if meal[i][j] != 0 {
				for k := 1; k <= M-1; k++ {
					// i도시까지 오는데 k번 이동했다면
					if dp[i][k] != -1 {
						dp[j][k+1] = max(dp[j][k+1], dp[i][k]+meal[i][j])
					}
				}
			}
		}
	}

	ans := 0

	// 기내식 점수의 최댓값 찾기
	for i := 1; i <= M; i++ {
		ans = max(ans, dp[N][i])
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
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
