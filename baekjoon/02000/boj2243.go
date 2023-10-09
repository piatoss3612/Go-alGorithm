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

	N    int
	tree [MAX*4 + 1]int
)

const MAX = 1000000

// 난이도: Platinum 5
// 메모리: 18088KB
// 시간: 100ms
// 분류: 세그먼트 트리, 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	var a, b, c int

	for i := 1; i <= N; i++ {
		a = scanInt()
		switch a {
		case 1:
			b = scanInt()
			pos := query(b, 1, MAX, 1)
			update(pos, -1, 1, MAX, 1)
			fmt.Fprintln(writer, pos)
		case 2:
			b, c = scanInt(), scanInt()
			update(b, c, 1, MAX, 1)
		}
	}
}

func update(target, diff, left, right, node int) {
	if target < left || right < target {
		return
	}

	tree[node] += diff
	if left == right {
		return
	}

	mid := (left + right) / 2
	update(target, diff, left, mid, node*2)
	update(target, diff, mid+1, right, node*2+1)
}

func query(target, left, right, node int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2
	if target <= tree[node*2] {
		return query(target, left, mid, node*2)
	}
	return query(target-tree[node*2], mid+1, right, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
