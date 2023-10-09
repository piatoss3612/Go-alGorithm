package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	T       int
	N, M, K int
	path    []Path
	tree    [MAX * 4]int
)

type Path struct {
	n, m int
}

const MAX = 1000

// 난이도: Platinum 5
// 메모리: 10800KB
// 시간: 284ms
// 분류: 자료 구조, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve(i)
	}
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	path = make([]Path, 0, K)
	tree = [MAX * 4]int{}

	for i := 1; i <= K; i++ {
		path = append(path, Path{n: scanInt(), m: scanInt()})
	}

	sort.Slice(path, func(i, j int) bool {
		if path[i].n == path[j].n {
			return path[i].m < path[j].m
		}
		return path[i].n < path[j].n
	})
}

func Solve(caseNo int) {
	total := 0
  for _, p := range path {
    total += query(p.m+1, M, 1, M, 1)
    update(p.m, 1, M, 1)
  }

	fmt.Fprintf(writer, "Test case %d: %d\n", caseNo, total)
}

func query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return tree[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return query(left, right, nodeLeft, mid, node*2) + query(left, right, mid+1, nodeRight, node*2+1)
}

func update(target, left, right, node int) {
	if target < left || right < target {
		return
	}

	tree[node] += 1

	if left == right {
		return
	}

	mid := (left + right) / 2
	update(target, left, mid, node*2)
	update(target, mid+1, right, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}