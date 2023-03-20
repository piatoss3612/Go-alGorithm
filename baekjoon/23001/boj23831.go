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
	N, A, B      int
	satisfaction [101][5]int
	dp           [101][101][101][5]int // dp[day][medical][study][prev], day: 날짜, medical: 요양실 이용 횟수, study: 자습 횟수, prev: 직전에 이용한 공간
)

const INF = -987654321

// 난이도: Gold 3
// 메모리: 40096KB
// 시간: 40ms
// 분류: 다이나믹 프로그래밍
// ※dp 초기값을 설정하지 않는 실수로 1회 틀림
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	A, B = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		satisfaction[i] = [5]int{0, scanInt(), scanInt(), scanInt(), scanInt()}
	}
}

func Solve() {
	ans := rec(0, 0, 0, 0)
	fmt.Fprintln(writer, ans)
}

func rec(day, medical, study, prev int) int {
	// 기저사례
	if day == N {
		// 정독실이나 소학습실에서 자습을 B회 미만으로 한 경우, 퇴사 처리
		if study < B {
			return INF
		}
		return 0
	}

	// 메모이제이션 (이미 계산한 값이면 바로 반환)
	ret := &dp[day][medical][study][prev]
	if *ret != INF {
		return *ret
	}

	*ret = INF // 최댓값을 구하는 문제이므로 초기값을 아주 작은 값으로 설정, 이부분을 0으로 설정하여 1회 틀림

	*ret = max(*ret, rec(day+1, medical, study+1, 1)+satisfaction[day+1][1]) // 정독실 자습
	*ret = max(*ret, rec(day+1, medical, study+1, 2)+satisfaction[day+1][2]) // 소학습실 자습

	// 휴게실 자습은 연속으로 할 수 없다
	if prev != 3 {
		*ret = max(*ret, rec(day+1, medical, study, 3)+satisfaction[day+1][3]) // 휴게실 자습, 자습 횟수는 증가하지 않음
	}

	// 요양 신청은 A회 이하로 한다
	if medical < A {
		*ret = max(*ret, rec(day+1, medical+1, study, 4)+satisfaction[day+1][4]) // 요양
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
