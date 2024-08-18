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
	K       int
	elems   []int
	tree    [][]int
)

// 9934번: 완전 이진 트리
// hhttps://www.acmicpc.net/problem/9934
// 난이도: 실버 1
// 메모리: 912 KB
// 시간: 4 ms
// 분류: 트리, 재귀
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	K = scanInt()
	elemsCnt := 1<<K - 1
	elems = make([]int, elemsCnt+1)
	tree = make([][]int, K+1)
	for i := 1; i <= elemsCnt; i++ {
		elems[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, 1<<K-1
	buildTree(l, r, 1)

	for i := 1; i <= K; i++ {
		for _, v := range tree[i] {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	}
}

func buildTree(l, r, depth int) {
	if l == r {
		tree[depth] = append(tree[depth], elems[l])
		return
	}

	mid := (l + r) / 2
	tree[depth] = append(tree[depth], elems[mid])

	buildTree(l, mid-1, depth+1)
	buildTree(mid+1, r, depth+1)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
