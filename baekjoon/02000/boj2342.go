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
	inst    []int                                           // 지시 사항
	N       int                                             // 지시 사항의 개수
	adj     = [5][2]int{{}, {2, 4}, {1, 3}, {2, 4}, {1, 3}} // 각 방향과 인접한 방향의 번호
	opp     = [5]int{0, 3, 4, 1, 2}                         // 각 방향의 반대쪽 방향의 번호
	dp      [100000][5][5]int                               // dp[i][j][k]: 왼발이 j, 오른발이 k에 있고 i번 지시를 따라야 할 때, 남은 모든 지시사항을 만족했을 경우에 필요한 힘의 최솟값
)

const INF = 987654321

// 난이도: Gold 3
// 메모리: 58220KB
// 시간: 76ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	for {
		n := scanInt()
		if n == 0 {
			return
		}

		inst = append(inst, n)
	}
}

func Solve() {
	N = len(inst)
	ans := rec(0, 0, 0) // 0번 지시사항, 왼발 0, 오른발 0인 상태에서 시작하여 모든 지시사항을 만족하는데 필요한 힘의 최솟값
	fmt.Fprintln(writer, ans)
}

// idx: idx (0~N-1)번째 지시 사항
// left: 왼발이 놓인 방향의 번호
// right: 오른발이 놓인 방향의 번호
// 재귀 호출을 통해 모든 지시 사항을 만족하는 데 사용되는 최소의 힘 구하기
func rec(idx, left, right int) int {
	// 기저 사례: 모든 지시 사항을 만족한 경우
	if idx == N {
		return 0
	}

	ret := &dp[idx][left][right]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 큰 정숫값으로 초기화

	next := inst[idx] // 수행해야 할 지시 사항

	// # 조건: 두 발이 같은 지점에 있으면 안된다 (시작 시점 제외)

	// 1. 왼발 움직이기

	// 1-1. 중앙에서 시작
	if left == 0 && next != right {
		*ret = min(*ret, rec(idx+1, next, right)+2)
	}

	// 1-2. 동일한 지점
	if next == left {
		*ret = min(*ret, rec(idx+1, next, right)+1)
	}

	// 1-3. 인접한 지점(왼쪽, 오른쪽)
	if left != 0 && next == adj[left][0] && next != right {
		*ret = min(*ret, rec(idx+1, next, right)+3)
	}

	if left != 0 && next == adj[left][1] && next != right {
		*ret = min(*ret, rec(idx+1, next, right)+3)
	}

	// 1-4. 반대편
	if left != 0 && next == opp[left] && next != right {
		*ret = min(*ret, rec(idx+1, next, right)+4)
	}

	// 2.오른발 움직이기

	// 2-1. 중앙에서 시작
	if right == 0 && next != left {
		*ret = min(*ret, rec(idx+1, left, next)+2)
	}

	// 2-2. 동일한 지점
	if next == right {
		*ret = min(*ret, rec(idx+1, left, next)+1)
	}

	// 2-3. 인접한 지점(왼쪽, 오른쪽)
	if right != 0 && next == adj[right][0] && next != left {
		*ret = min(*ret, rec(idx+1, left, next)+3)
	}

	if right != 0 && next == adj[right][1] && next != left {
		*ret = min(*ret, rec(idx+1, left, next)+3)
	}

	// 2-4. 반대편
	if right != 0 && next == opp[right] && next != left {
		*ret = min(*ret, rec(idx+1, left, next)+4)
	}

	return *ret
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
