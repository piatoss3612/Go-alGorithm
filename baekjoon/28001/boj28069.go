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
	visited [MAX + 1][3][3]bool
)

const MAX = 1000000

// 난이도: Gold 5
// 메모리: 92836KB
// 시간: 356ms
// 분류: 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
}

func Solve() {
	q := [][3]int{{0, 0, 0}}
	visited[0][0][0] = true

	// 중복해서 여러 번 이동할 수 있는 경우는 0번째 계단에서 제자리걸음하는 경우뿐이다

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		stair, cnt, prev := front[0], front[1], front[2]

		// N번째 계단에 도착했고, K번 이하로 계단을 오른 경우
		// K번 미만으로 오른 경우는 0번째 계단에서 제자리걸음하는 횟수를 늘려주면 되므로
		// K번 만에 N번째 계단에 도착했다고 볼 수 있다
		if stair == N && cnt <= K {
			fmt.Fprintln(writer, "minigimbob")
			return
		}

		// K번 계단을 올랐지만 N번째 계단에 도착하지 못한 경우
		if cnt == K {
			continue
		}

		// 계단을 한 칸 올라가는 경우
		next := stair + 1
		if next > N || visited[next][prev][1] {
			continue
		} else if !visited[next][prev][1] {
			visited[next][prev][1] = true
			q = append(q, [3]int{next, cnt + 1, 1})
		}

		// 오른 계단의 개수 / 2 만큼 올라가는 경우 (내림 처리)
		next = stair + stair/2
		if next > N || visited[next][prev][2] {
			continue
		} else if !visited[next][prev][2] {
			visited[next][prev][2] = true
			q = append(q, [3]int{next, cnt + 1, 2})
		}
	}

	fmt.Fprintln(writer, "water")
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
