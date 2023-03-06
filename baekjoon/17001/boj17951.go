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
	N, K    int
	paper   []int
)

// 난이도: Gold 4
// 메모리: 1828KB
// 시간: 20ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	paper = make([]int, N+1)
	for i := 1; i <= N; i++ {
		paper[i] = scanInt()
	}
}

func Solve() {
	// 문제 조건:
	// 시험지를 현재 순서 그대로 K개의 그룹으로 나눈 뒤 -> 그룹에 속한 원소들은 연속해야 한다
	// 각각의 그룹에서 맞은 문제 개수의 합을 구하여 그 중 최솟값을 시험 점수로 하기로 하였다 -> 그룹의 점수합이 최솟값보다 작아서는 안된다
	// 현수가 이번 시험에서 받을 수 있는 최대 점수를 계산하는 프로그램을 작성하자 -> 최솟값의 최댓값을 구한다 (이분 탐색 upper bound)
	l, r := 0, 20*100000
	for l <= r {
		m := (l + r) / 2
		if MeetScore(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Fprintln(writer, r) // 최솟값의 최댓값 (upper bound) 출력
}

// 각 그룹별 점수의 합의 최솟값을 expected라고 가정했을 때
// 만들어지는 그룹의 개수가 K개 이상인지 확인한다
func MeetScore(expected int) bool {
	sum := 0
	cnt := 0

	for i := 1; i <= N; i++ {
		sum += paper[i]
		// 탐색중인 그룹의 점수합이 expected 이상인 경우
		if sum >= expected {
			cnt++   // 그룹의 개수를 증가시키고 다음 그룹 탐색
			sum = 0 // 합 초기화
		}
	}

	// 만들어지는 그룹의 개수가 K개를 초과한 경우
	// 임의로 그룹을 K개의 그룹으로 재구성할 수 있으므로 OK
	return cnt >= K
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
