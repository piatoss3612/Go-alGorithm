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

	N, K int
	arr  [250001]int
	tree [200000]int
)

const MAX = 65535

// 난이도: platinum 5
// 메모리: 5332KB
// 시간: 148ms
// 분류: 세그먼트 트리, 이분 탐색
// 비고: 1572번과 동일한 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
}

func Solve() {
	sum := 0
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
		update(arr[i], 1, 0, MAX, 1)

		if i >= K {
			sum += query((K+1)/2, 0, MAX, 1)
			update(arr[i-K+1], -1, 0, MAX, 1)
		}
	}

	fmt.Fprintln(writer, sum)
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

func query(cnt, left, right, node int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2
	if tree[node*2] >= cnt {
		return query(cnt, left, mid, node*2)
	}
	return query(cnt-tree[node*2], mid+1, right, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
