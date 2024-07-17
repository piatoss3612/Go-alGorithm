package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	n, m, c0, d0 int
	mandus       []Mandu
)

type Mandu struct {
	a, b, c, d int
}

// 14855번: 만두 가게 사장 박승원
// hhttps://www.acmicpc.net/problem/14855
// 난이도: 골드 4
// 메모리: 932 KB
// 시간: 12 ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	// n: 현재 가지고 있는 밀가루의 양
	// m: 만두의 종류의 수
	// c0: 스페셜 만두를 만들기 위해 필요한 밀가루의 양
	// d0: 스페셜 만두의 가격
	n, m, c0, d0 = scanInt(), scanInt(), scanInt(), scanInt()
	mandus = make([]Mandu, m+1)
	for i := 1; i <= m; i++ {
		// a: i번째 만두의 만두 속의 남은 양
		// b: i번째 만두 하나를 만들기 위해 필요한 만두 속의 양
		// c: i번째 만두를 만들기 위해 필요한 밀가루의 양
		// d: i번째 만두의 가격
		a, b, c, d := scanInt(), scanInt(), scanInt(), scanInt()
		mandus[i] = Mandu{a, b, c, d}
	}
}

func Solve() {
	// dp[i][j]: i번째 만두를 몇 개 만들었고, j만큼의 밀가루를 사용했을 때의 만두 금액의 최대값
	// dp[i][j] = max(dp[i-1][j], dp[i-1][j-k*c]+k*d) (1 <= k <= j/c, a >= k*b)
	// 배낭 문제처럼 만두를 하나씩 만들어가면서 최대값을 갱신해나간다.
	dp := [12][1001]int{}

	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// i번째 만두를 만들지 않는 경우
			dp[i][j] = max(dp[i][j], dp[i-1][j])

			// i번째 만두를 만드는 경우
			a, b, c, d := mandus[i].a, mandus[i].b, mandus[i].c, mandus[i].d
			// k: i번째 만두를 몇 개 만들 수 있는지
			for k := 1; k*c <= j; k++ {
				// i번째 만두를 k개 만들기 위해 필요한 만두 속의 양이 a보다 작거나 같아야 한다.
				if a >= k*b {
					// i-1번째 만두를 만든 후, j-k*c만큼의 밀가루를 사용했을 때의 금액에
					// i번째 만두를 k개 만들었을 때의 금액을 더한 값을 dp[i][j]와 비교하여 최대값을 저장
					dp[i][j] = max(dp[i][j], dp[i-1][j-k*c]+k*d)
				}
			}
		}
	}

	// 스페셜 만두 만들기
	i := m + 1
	for j := 0; j <= n; j++ {
		// 스페셜 만두를 만들지 않는 경우
		dp[i][j] = dp[i-1][j]

		// 스페셜 만두를 만드는 경우
		for k := 1; k*c0 <= j; k++ {
			dp[i][j] = max(dp[i][j], dp[i-1][j-k*c0]+k*d0)
		}
	}

	ans := 0

	for i := 0; i <= n; i++ {
		ans = max(ans, dp[m+1][i])
	}

	fmt.Fprintln(writer, ans)
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
