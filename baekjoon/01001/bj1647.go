package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type link struct {
	a, b, c int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n, m    int
	links   []link
	parent  []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	links = make([]link, m)
	parent = make([]int, n+1)
	for i := 0; i < m; i++ {
		links[i] = link{scanInt(), scanInt(), scanInt()}
	}

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	sort.Slice(links, func(i, j int) bool {
		return links[i].c < links[j].c
	})

	var ans int
	numberOfLinks := 0

	// 2개의 마을로 분리하는데 하나의 마을에 집이 최소 하나있어야 한다
	// 2개의 마을을 연결하는 선은 제거
	// 따라서 최소 신장 트리의 간선의 갯수(정점의 갯수 - 1)에서
	// 1개를 뺀만큼의 간선을 크루스칼 알고리즘을 통해 구하면 된다

	for _, link := range links {
		if numberOfLinks == n-2 {
			break
		}
		if !sameParent(link.a, link.b) {
			union(link.a, link.b)
			ans += link.c
			numberOfLinks += 1
		}
	}
	fmt.Fprintln(writer, ans)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		parent[y] = x
	}
}

func sameParent(x, y int) bool {
	x = find(x)
	y = find(y)
	if x == y {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
