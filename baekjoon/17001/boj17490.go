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
	N, M, K      int
	rocks        []int        // i번 강의동과 와우도를 연결하는데 필요한 돌의 개수
	parent       []int        // 연결된 루트 요소 (루트: 와우도와 연결하는데 필요한 돌이 가장 적은 강의동)
	disconnected [][]int      // 공사로 인해 i번 강의동과 연결이 끊어진 강의동의 번호
	sets         map[int]bool // 공사로 인해 분리된 강의동 집합의 루트 요소 판별
	totalRocks   int          // 모든 강의동을 연결하는데 필요한 돌의 개수
)

// 난이도: Gold 4
// 메모리: 197476KB
// 시간: 860ms
// 분류: 분리 집합, 최소 스패닝 트리
// 분리 집합으로 풀었다
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()
	rocks = make([]int, N+1)
	parent = make([]int, N+1)
	disconnected = make([][]int, N+1)
	// i번 강의동과 와우도를 연결하는데 필요한 돌의 개수 입력
	for i := 1; i <= N; i++ {
		rocks[i] = scanInt()
		parent[i] = i
	}
	// 연결이 끊어진 강의동의 관계 입력
	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		disconnected[a] = append(disconnected[a], b)
		disconnected[b] = append(disconnected[b], a)
	}
}

func Solve() {
	// 공사로 인해 끊어진 길의 개수가 1개 이하인 경우
	// 돌을 놓지 않아도 모든 강의동은 이미 연결되어 있다
	if M <= 1 {
		fmt.Fprintln(writer, "YES")
		return
	}

	for i := 1; i <= N; i++ {
		left, right := neighbors(i) // i번 강의동의 왼쪽, 오른쪽에 있는 강의동 번호 확인

		// i번 강의동과 연결이 끊어진 강의동의 개수에 따라 분기 처리
		switch len(disconnected[i]) {
		// 끊어진 강의동이 없는 경우
		case 0:
			union(i, left)
			union(i, right)
		// 끊어진 강의동이 1개인 경우: 왼쪽이 끊어졌으면 오른쪽과 union, 오른쪽이 끊어졌으면 왼쪽과 union
		case 1:
			if disconnected[i][0] == left {
				union(i, right)
			} else {
				union(i, left)
			}
		}
	}

	sets = make(map[int]bool)
	for i := 1; i <= N; i++ {
		x := find(i) // i번 강의동과 연결된 루트 요소 확인

		// 아직 x를 체크하지 않은 경우
		if !sets[x] {
			sets[x] = true
			totalRocks += rocks[x]
		}
	}

	// 모든 강의동을 연결하는데 필요한 돌의 개수가 K개 이하인 경우
	if totalRocks <= K {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		if rocks[x] > rocks[y] {
			parent[x] = y
		} else {
			parent[y] = x
		}
	}
}

func neighbors(x int) (int, int) {
	left := x - 1
	if left == 0 {
		left = N
	}
	right := x + 1
	if right > N {
		right = 1
	}
	return left, right
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
