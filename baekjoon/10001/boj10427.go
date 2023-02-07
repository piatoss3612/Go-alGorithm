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
	T, N    int
	A       []int
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 1208KB
// 시간: 20ms
// 분류: 정렬, 부분합, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N = scanInt()
	A = make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = scanInt()
	}

	// 그리디 알고리즘:
	// N개의 수 중에서 i(1~N)개를 골라 (i개의 수 중 가장 큰 수 * i) = x, (i개의 수를 합한 값) = y라고 했을 때
	// x-y를 최소화하려면 i개의 수는 연속해서 증가하는 수 또는 연속해서 감소하는 수여야 한다
	// 따라서 입력받은 값을 오름차순으로 정렬한다
	sort.Ints(A)
}

func Solve() {
	pSumA := make([]int, N+1) // 부분합
	S := make([]int, N+1)     // N번 중 M번을 선택하여 추가적으로 갚아야 하는 돈의 최솟값

	for i := 1; i <= N; i++ {
		pSumA[i] = pSumA[i-1] + A[i] // 부분합 구하기
		S[i] = INF                   // 최솟값 비교를 위해 INF로 초기화
	}

	for i := 1; i <= N; i++ {
		for j := i; j <= N; j++ {
			// i에서 j까지 선택하여 빚을 갚을 경우의 추가 비용의 최솟값
			// A[j]: i에서 j까지 선택했을 때 가장 많은 돈을 빌린 경우의 금액
			// j-i+1: 돈을 갚는 횟수
			// pSumA[j] - pSumA[i-1]: i부터 j까지의 빌린 금액의 부분합
			extra := A[j]*(j-i+1) - (pSumA[j] - pSumA[i-1])
			S[j-i+1] = min(S[j-i+1], extra) // 최솟값 갱신
		}
	}

	// 1부터 N까지 S의 원소의 합 구하기
	ans := 0
	for i := 1; i <= N; i++ {
		ans += S[i]
	}

	fmt.Fprintln(writer, ans)
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
