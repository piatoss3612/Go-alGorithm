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
	n, m, k int
	parent  []int
)

// 난이도: Gold 3
// 메모리: 5904KB
// 시간: 200ms
// 분류: 분리 집합, 다이나믹 프로그래밍, 배낭 문제
// 15508번 문제와 풀이 동일
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	n, m, k = scanInt(), scanInt(), scanInt()
	parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		Union(a, b)
	}
}

func Solve() {
	set := make(map[int]int)

	for i := 1; i <= n; i++ {
		set[Find(i)]++
	}

	groups := []int{}

	for _, v := range set {
		// 15508번 문제 풀이에서 추가된 부분
		// 집합의 크기가 n-k일 때도 나머지를 가지고 k를 만들어
		// 크기가 n-k, k인 집합 두 개를 만들 수 있다
		if v == k || v == n-k {
			fmt.Fprintln(writer, "SAFE")
			return
		}
		if v < k {
			groups = append(groups, v)
		}
	}

	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = -1
	}

	for i := 0; i < len(groups); i++ {
		for j := k; j >= 0; j-- {
			if dp[j] != -1 && j+groups[i] <= k {
				dp[j+groups[i]] = 1
			}
		}
	}

	if dp[k] == 1 {
		fmt.Fprintln(writer, "SAFE")
	} else {
		fmt.Fprintln(writer, "DOOMED")
	}
}

func Union(x, y int) {
	px, py := Find(x), Find(y)
	if px != py {
		parent[py] = px
	}
}

func Find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
