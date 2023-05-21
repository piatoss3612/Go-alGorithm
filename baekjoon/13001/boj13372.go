package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	T, N int
	tree [MAX * 4]int
)

const MAX = 100000

// 난이도: Platinum 5
// 메모리: 8864KB
// 시간: 480ms
// 분류: 세그먼트 트리
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
	tree = [MAX * 4]int{}
}

func Solve() {
	cross := 0

	for i := 1; i <= N; i++ {
		conn := scanInt()
		cross += query(conn+1, N, 1, N, 1)
		update(conn, 1, 1, N, 1)
	}

	fmt.Fprintln(writer, cross)
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

func update(target, diff, left, right, node int) {
	if target < left || right < target {
		return
	}

	tree[node] += diff

	if left != right {
		mid := (left + right) / 2
		update(target, diff, left, mid, node*2)
		update(target, diff, mid+1, right, node*2+1)
	}
}
*/

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	T, N    int
	fenwick [MAX + 1]int
)

const MAX = 100000

// 난이도: Platinum 5
// 메모리: 6476KB
// 시간: 224ms
// 분류: 펜윅 트리
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
	fenwick = [MAX + 1]int{}
}

func Solve() {
	cross := 0

	for i := 1; i <= N; i++ {
		conn := scanInt()
		cross += sum(N) - sum(conn)
		add(conn, 1)
	}

	fmt.Fprintln(writer, cross)
}

func sum(pos int) (ret int) {
	for pos > 0 {
		ret += fenwick[pos]
		pos &= pos - 1
	}
	return
}

func add(pos, val int) {
	for pos <= N {
		fenwick[pos] += val
		pos += pos & -pos
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}