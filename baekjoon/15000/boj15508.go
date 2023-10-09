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

// 난이도: Gold 4
// 메모리: 916KB
// 시간: 4ms
// 분류: 분리 집합, 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	n, m, k = scanInt(), scanInt(), scanInt()

	// 각 원소가 속한 집합의 부모 원소를 자기자신으로 초기화
	parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		Union(a, b) // a가 속한 집합과 b가 속한 집합을 유니온
	}
}

func Solve() {
	set := make(map[int]int) // 각 집합의 부모 원소의 번호와 집합에 속한 원소들의 개수를 매핑

	for i := 1; i <= n; i++ {
		set[Find(i)]++
	}

	groups := []int{} // k보다 작은 원소의 개수를 가진 집합들의 원소의 개수의 모음

	for _, v := range set {
		if v <= k {
			// 한 집합에 속한 원소의 개수가 k개인 경우는 항상 k개의 원소를 포함한 집합과 n-k개의 원소를 포함한 집합을 만들 수 있다
			if v == k {
				fmt.Fprintln(writer, "SAFE")
				return
			}
			groups = append(groups, v)
		}
	}

	dp := make([]int, k+1) // k보다 작은 원소의 개수를 가진 집합들의 원소의 개수를 조합하여 1~k를 만들 수 있는지에 대한 메모이제이션
	for i := 1; i <= k; i++ {
		dp[i] = -1
	}

	// 배낭 문제
	for i := 0; i < len(groups); i++ {
		for j := k; j >= 0; j-- {
			if dp[j] != -1 && j+groups[i] <= k {
				dp[j+groups[i]] = 1
			}
		}
	}

	// 원소의 개수를 조합하여 k를 만들 수 있는 경우와 그렇지 않은 경우
	if dp[k] == 1 {
		fmt.Fprintln(writer, "SAFE")
	} else {
		fmt.Fprintln(writer, "DOOMED")
	}
}

// x가 속한 집합과 y가 속한 집합을 유니온한다
func Union(x, y int) {
	px, py := Find(x), Find(y)
	// x가 속한 집합의 부모 원소와 y가 속한 집합의 부모 원소가 다를 경우
	if px != py {
		parent[py] = px // 유니온
	}
}

// x가 속한 집합의 부모 원소의 번호를 찾는다
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
