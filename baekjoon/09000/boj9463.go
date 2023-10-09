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

	T, N int
	arr  [MAX + 1]int
	tree [MAX*4 + 1]int
)

const MAX = 100000

// 난이도: Platinum 5
// 메모리: 12092KB
// 시간: 2820ms
// 분류: 자료 구조, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N = scanInt()
	arr = [MAX + 1]int{}
	tree = [MAX*4 + 1]int{}

	for i := 1; i <= N; i++ {
		arr[scanInt()] = i
	}
}

func Solve() {
	total := 0
	for i := 1; i <= N; i++ {
		x := scanInt()
		total += query(arr[x]+1, N, 1, N, 1)
		update(arr[x], 1, 1, N, 1)
	}
	fmt.Fprintln(writer, total)
}

func query(left, right, nodeLeft, nodeRight, node int) int {
	if nodeRight < left || right < nodeLeft {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return tree[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return query(left, right, nodeLeft, mid, node*2) + query(left, right, mid+1, nodeRight, node*2+1)
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

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
