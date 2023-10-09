package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	// 현재 위치가 i번 선택지이고 j(0<=j<=1)회 선택지를 건너뛴 상황에서 왼쪽(0)또는 오른쪽(1) 선택지를 선택하는 경우의 사람의 최대 수
	dp [100001][2][2]int
)

// 난이도: Gold 4
// 메모리: 10900KB
// 시간: 36ms
// 분류: 다이나믹 프로그래밍
// 진짜 똥게임...
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanLines)
	N = scanInt()
}

func Solve() {
	// 선택지를 건너뛰지 않은 상태에서 1명으로 시작
	dp[0][0][0] = 1
	dp[0][0][1] = 1

	for i := 1; i <= N; i++ {
		options := scanOptions()

		// i번째 선택지에 대해 j가 0일 때는 왼쪽 선택지를, 1일 때는 오른쪽 선택지를 선택
		for j := 0; j <= 1; j++ {
			// 87%에서 틀린 이유는 게임 오버 조건을 체크하지 않았기 때문!
			// 각각의 경우에 대해서 게임이 오버되었는지 제대로 체크해주고 조건을 만족하는 경우에만 값을 갱신해준다

			// 1. 아직 선택지를 건너뛰지 않았고 이전에 왼쪽 선택지를 선택한 결과의 최댓값이 0보다 큰 경우
			if dp[i-1][0][0] > 0 {
				dp[i][0][j] = max(dp[i][0][j], calculate(options[j], dp[i-1][0][0]))
				dp[i][1][j] = max(dp[i][1][j], dp[i-1][0][0])
			}

			// 2. 아직 선택지를 건너뛰지 않았고 이전에 오른쪽 선택지를 선택한 결과의 최댓값이 0보다 큰 경우
			if dp[i-1][0][1] > 0 {
				dp[i][0][j] = max(dp[i][0][j], calculate(options[j], dp[i-1][0][1]))
				dp[i][1][j] = max(dp[i][1][j], dp[i-1][0][1])
			}

			// 3. 선택지를 한 번 건너뛰고 이전에 왼쪽 선택지를 선택한 결과의 최댓값이 0보다 큰 경우
			if dp[i-1][1][0] > 0 {
				dp[i][1][j] = max(dp[i][1][j], calculate(options[j], dp[i-1][1][0]))
			}

			// 4. 선택지를 한 번 건너뛰고 이전에 왼쪽 선택지를 선택한 결과의 최댓값이 0보다 큰 경우
			if dp[i-1][1][1] > 0 {
				dp[i][1][j] = max(dp[i][1][j], calculate(options[j], dp[i-1][1][1]))
			}
		}
	}

	ans := max(max(max(dp[N][0][0], dp[N][0][1]), dp[N][1][0]), dp[N][1][1]) // 마지막 선택지를 선택한 모든 결과의 최댓값

	if ans <= 0 {
		fmt.Fprintln(writer, "ddong game")
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func calculate(option string, value int) int {
	var res int

	switch option[0] {
	case '+':
		res = value + int(option[1]-'0')
	case '-':
		res = value - int(option[1]-'0')
	case '*':
		res = value * int(option[1]-'0')
	case '/':
		res = value / int(option[1]-'0')
	}
	return res
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

func scanOptions() []string {
	scanner.Scan()
	return strings.Fields(scanner.Text())
}
