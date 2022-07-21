package main

import (
	"bufio"
	"fmt"
	"os"
	_ "sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	children []child // i번째 아이의 친구의 친구의 친구의 총 친구 수, 가지고 있는 캔디 수를 저장
	leader   []int   // 분리 집합을 적용하기 위해 친구가 가장 많은 아이가 임시 리더를 맡게 된다
	n, m, k  int
)

type child struct {
	friend, candy int
}

// 메모리: 3012KB
// 시간: 116ms
// 이게 골드3?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, k = scanInt(), scanInt(), scanInt()
	leader = make([]int, n+1)
	children = make([]child, n+1)
	for i := 1; i <= n; i++ {
		leader[i] = i
		children[i].candy = scanInt()
		children[i].friend = 1
	}

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		union(a, b) // a와 b가 속한 그룹의 친구 수, 캔디 수를 합친다
	}

	dp := make([]int, k)

	for i := 1; i <= n; i++ {
		// i가 속한 그룹의 통이 자기자신인 경우에만 탐색
		if leader[i] != i {
			continue
		}

		// 통의 정보를 가져온다
		leaderChild := children[i]
		// 여기서 leaderChild.friend부터 시작하면 값이 누적되서 이상하게 큰 값이 나오게 된다
		// 반드시 뒤에서부터 탐색해 준다
		for j := k - 1; j >= leaderChild.friend; j-- {
			dp[j] = max(dp[j], dp[j-leaderChild.friend]+leaderChild.candy)
		}
	}

	fmt.Fprintln(writer, dp[k-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func find(x int) int {
	if leader[x] == x {
		return x
	}
	leader[x] = find(leader[x])
	return leader[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		if children[x].friend >= children[y].friend {
			leader[y] = x
		} else {
			leader[x] = y
		}
		children[x].friend += children[y].friend
		children[y].friend = children[x].friend

		children[x].candy += children[y].candy
		children[y].candy = children[x].candy
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
