package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	children [1000001]int // children[i]: i번째 어린이의 번호
	dp       [1000001]int // dp[i]: 번호가 i인 어린이를 마지막으로 포함하는 1씩 증가하는 최장 증가 부분 수열의 길이
)

// 난이도: Gold 3
// 메모리: 21844KB
// 시간: 140ms
// 분류: 다이나믹 프로그래밍, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		children[i] = scanInt()
	}
}

func Solve() {
	/*
		# 초기 시도

		처음으로 시도한 풀이 방법은 반복문을 사용하여 LIS(최장 증가 부분 수열)를 찾아 정렬되지 않은 어린이들의 수 구하기 -> 시간 초과
		두 번째로 시도한 방법은 이분 탐색을 사용하여 LIS를 찾아 정렬되지 않은 어린이들의 수 구하기 -> 틀렸습니다

		# 틀린 이유

		문제에서는 어린이들을 제일 앞이나 제일 뒤로 이동시켜야 한다
		그런데 LIS를 구해 LIS에 포함되지 않은 어린이들을 정렬하는 방식으로는 문제에서 제시한 이동 방법을 적용한 해답을 항상 구할 수 없다

		ex)
		6
		1 2 3 5 6 4

		LIS를 구하면 1 2 3 5 6으로 4번 어린이만 1회 이동하면 될 것 같지만
		4번 어린이를 1 2 3과 5 6 사이에 끼워 넣는 방법을 사용하는 경우에만 1회 이동으로 해결할 수 있고
		4번 어린이를 제일 앞이나 제일 뒤로 이동시켜서 모든 어린이를 번호순으로 정렬시킬 수 있는 방법은 없다!

		이 경우, 어린이들을 번호순으로 정렬시키려면 5번 어린이를 제일 뒤로 보내고 6번 어린이를 제일 뒤로 보내는 것으로 총 2회 이동시켜야 한다


		# 해결

		제일 앞이나 제일 뒤로 보내는 어린이의 수를 최소화하는 것은
		이동시키지 않을 어린이들(빈자리를 메꾸는 경우 제외)의 수를 최대화하는 것과 마찬가지 이므로
		LIS는 문제 해결을 위해 논리적으로 유효한 방법이다

		그렇지만 문제에서 제시한 이동 방법을 항상 만족시킬 수 있는 해를 찾기 위해 LIS 탐색에 조건이 추가되어야 한다

		*LIS 추가 조건*
		증가 폭은 1로써 연속되는 번호여야 한다

		즉, 증가폭이 1인 최장 증가 부분 수열 S를 고정시키고 S의 최솟값보다 작은 값들은 번호순에 맞춰 S의 왼쪽으로,
		S의 최댓값보다 큰 값들은 번호순에 맞춰 S의 오른쪽으로 이동시킴으로써 이동시키는 어린이의 수를 최소화시킬 수 있다

		[참고] https://www.acmicpc.net/board/view/14700
	*/

	temp := 0 // 증가폭이 1인 최장 증가 부분 수열의 길이의 최댓값
	for i := 1; i <= N; i++ {
		dp[children[i]] = 1                                         // i번째 어린이의 번호 children[i]로 끝나는 최장 증가 부분 수열의 길이를 1로 초기화
		dp[children[i]] = max(dp[children[i]], dp[children[i]-1]+1) // children[i]-1번인 어린이가 포함된 최장 증가 부분 수열이 존재하는 경우, children[i]번 어린이를 포함시켜 최댓값 갱신
		temp = max(temp, dp[children[i]])
	}
	fmt.Fprintln(writer, N-temp)
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
