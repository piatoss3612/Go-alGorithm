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
	N, M         int
	numOfFriends []int                // numOfFriends[i]: i의 친구의 수
	areFriends   map[int]map[int]bool // areFriends[i][j]: i와 j가 친구인지 판별
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 2744KB
// 시간: 340ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	numOfFriends = make([]int, N+1)
	areFriends = make(map[int]map[int]bool)
	for i := 1; i <= N; i++ {
		areFriends[i] = make(map[int]bool)
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		numOfFriends[a]++
		numOfFriends[b]++
		areFriends[a][b] = true
		areFriends[b][a] = true
	}
}

func Solve() {
	ans := INF

	// a, b, c가 모두 친구인 경우 찾기
	for a := 1; a <= N-2; a++ {
		for b := a + 1; b <= N-1; b++ {
			// a와 b가 친구인 것이 확인된 경우에만 c를 찾는다
			if areFriends[a][b] {
				for c := b + 1; c <= N; c++ {
					if areFriends[b][c] && areFriends[c][a] {
						temp := numOfFriends[a] + numOfFriends[b] + numOfFriends[c] - 6 // a, b, c가 모두 친구이므로 친구 관계의 총합에서 항상 6을 빼야한다
						ans = min(ans, temp)
					}
				}
			}

		}
	}

	if ans == INF {
		fmt.Fprintln(writer, -1) // 세 명이 모두 친구인 경우를 찾기 못했을 때
	} else {
		fmt.Fprintln(writer, ans)
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
