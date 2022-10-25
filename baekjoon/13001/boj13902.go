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
	wok     []int
	dp      [10001]int
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 2204KB
// 시간: 164ms
// 분류: 다이나믹 프로그래밍
// 회고: 13910번 풀어놓은 거 그대로 변수 범위만 바꿔서 냈다가 시간 초과
func main() {
	defer writer.Flush()
	Input()
	Preprocess()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	wok = make([]int, M)
	for i := 0; i < M; i++ {
		wok[i] = scanInt()
	}
}

func Preprocess() {
	for i := 1; i <= N; i++ {
		dp[i] = INF
	}

	// 웍의 개수가 최대 100개일 때는 2개를 조합하는 경우의 수가 100C2로 그렇게 많지 않아서
	// 궂이 중복된 값을 제거할 필요가 없었지만 웍의 개수의 최댓값이 1000이 되어버리면
	// 조금 말이 달라지므로 시간 초과가 나는 것을 피하기 위해 중복된 값을 제거해줍니다
	// go언어에도 파이썬 set 자료형 같은 게 필요하다고 생각합니다
	set := make(map[int]bool)

	for i := 0; i < M-1; i++ {
		for j := i + 1; j < M; j++ {
			if temp := wok[i] + wok[j]; temp <= N && !set[temp] {
				set[temp] = true
			}
		}
	}

	for k, _ := range set {
		wok = append(wok, k)
	}

	sort.Ints(wok)
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 0; j < len(wok); j++ {
			if wok[j] <= i {
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
