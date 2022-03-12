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
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for i := 0; i < m; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		links[i] = link{a, b, c}
	}
	sort.Slice(links, func(i, j int) bool {
		return links[i].c < links[j].c
	})

	// 크루스칼 알고리즘
	// 최소 신장 트리: 정점의 갯수 - 1개의 간선이 필요

	ans := 0
	numberOfLinks := 0

	for _, link := range links {
		if numberOfLinks == n-1 { // 간선의 개수: n - 1개인 경우 종료
			break
		}
		// 정점 a와 정점 b의 부모가 다른 경우
		// 정점 b와 같은 부모를 가진 정점들의 부모 값을 정점 a의 부모로 바꾼다 (유니온)
		if parent[link.a] != parent[link.b] {
			ans += link.c
			numberOfLinks += 1
			prev := parent[link.b]
			current := parent[link.a]
			for i := 1; i <= n; i++ {
				if parent[i] == prev {
					parent[i] = current
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
