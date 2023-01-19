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
	N, M    int
	chart   [101]int       // N개의  월봉의 절댓값
	dp      [101][1024]int // dp[i][j]: 중복을 포함하여 i개의 월봉을 골랐을 때, 절댓값들을 xor 연산하여 절댓값 j를 만들 수 있으면 1, 없으면 0
	// j의 범위가 1023까지인 이유는 월봉의 절댓값의 최대가 1023으로 xor연산을 실행할 경우 이보다 큰 값을 만들 수 없기 때문
)

// 난이도: Gold 3
// 메모리: 1672KB
// 시간: 20ms
// 분류: 다이나믹 프로그래밍, 배낭 문제

// (다음 월봉의 절댓값) = 이전 N개의 월봉 중 중복을 허용해 M개를 골라 절댓값들을 bitwise xor 한 것 중 최대
// '최대'라는 표현에 혹해서 xor 연산의 최댓값을 구해야 하는 구나! 하면 문제가 어려워 진다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		chart[i] = scanInt()
		if chart[i] < 0 {
			chart[i] = -chart[i] // 절댓값으로 저장
		}
	}
}

func Solve() {
	dp[0][0] = 1 // 0개의 월봉으로 0을 만드는 경우는 항상 1

	// 배낭 문제: i번째 월봉을 선택할 차례
	for i := 1; i <= M; i++ {
		// i번째 월봉을 선택하기 이전에 i-1개의 월봉을 선택하여 구한 절댓값을 탐색
		for j := 1023; j >= 0; j-- {
			// i-1개의 월봉을 선택하여 절댓값 j를 구할 수 있는 경우
			if dp[i-1][j] == 1 {
				// 어떤 월봉을 선택해야 M개의 월봉을 선택했을 때의 최댓값을 구할 수 있을지 모르기 때문에
				// i번째 월봉으로 특정한 월봉을 선택하는 것이 아니라,
				// 모든 월봉의 절댓값과 절댓값 j를 xor연산을 실행하여 새로운 절댓값 또는 중복된 절댓값을 구한다
				for k := 1; k <= N; k++ {
					dp[i][j^chart[k]] = 1
				}
			}
		}
	}

	ans := 0
	// M개의 월봉을 선택하여 구할 수 있는 가장 큰 절댓값 찾기
	for i := 0; i <= 1023; i++ {
		if dp[M][i] == 1 {
			ans = i
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
