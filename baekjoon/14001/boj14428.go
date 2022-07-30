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
	tree    [270000]int
	inp     [100001]int
	N, M    int
)

// 메모리: 5736KB
// 시간: 100ms
// 최솟값의 인덱스를 찾는 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}

	SegmentTree(1, N, 1)

	M = scanInt()

	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()

		switch a {
		case 1:
			inp[b] = c
			Update(b, 1, N, 1)
		case 2:
			fmt.Fprintln(writer, Query(b, c, 1, N, 1))
		}
	}
}

func SegmentTree(left, right, node int) int {
	if left == right {
		tree[node] = left
		return tree[node]
	}

	mid := (left + right) / 2
	onLeft := SegmentTree(left, mid, node*2)
	onRight := SegmentTree(mid+1, right, node*2+1)

	if inp[onLeft] <= inp[onRight] {
		tree[node] = onLeft
	} else {
		tree[node] = onRight
	}
	return tree[node]
}

func Update(target, left, right, node int) int {
	if target < left || right < target {
		return tree[node]
	}

	if left == right {
		tree[node] = left
		return tree[node]
	}

	mid := (left + right) / 2
	onLeft := Update(target, left, mid, node*2)
	onRight := Update(target, mid+1, right, node*2+1)

	if inp[onLeft] <= inp[onRight] {
		tree[node] = onLeft
	} else {
		tree[node] = onRight
	}
	return tree[node]
}

func Query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return left
	}

	if left <= nodeLeft && nodeRight <= right {
		return tree[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	onLeft := Query(left, right, nodeLeft, mid, node*2)
	onRight := Query(left, right, mid+1, nodeRight, node*2+1)

	if inp[onLeft] <= inp[onRight] {
		return onLeft
	}

	return onRight
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
