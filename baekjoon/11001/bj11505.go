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
	inp     [1000001]int
	seg     [2100000]int
	n, m, k int
)

const MOD = 1000000007

// 메모리: 30624KB
// 시간: 192ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, k = scanInt(), scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	SegmentTree(1, n, 1)

	var a, b, c int
	for i := 1; i <= m+k; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		if a == 1 {
			Update(b, c, 1, n, 1)
		} else {
			fmt.Fprintln(writer, Query(b, c, 1, n, 1))
		}
	}
}

func Query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return 1
	}

	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return (Query(left, right, nodeLeft, mid, node*2) * Query(left, right, mid+1, nodeRight, node*2+1)) % MOD
}

func Update(index, newVal, left, right, node int) int {
	if index < left || index > right {
		return seg[node]
	}

	if left == right {
		seg[node] = newVal
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = (Update(index, newVal, left, mid, node*2) * Update(index, newVal, mid+1, right, node*2+1)) % MOD
	return seg[node]
}

func SegmentTree(left, right, node int) int {
	if left == right {
		seg[node] = inp[left]
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = (SegmentTree(left, mid, node*2) * SegmentTree(mid+1, right, node*2+1)) % MOD
	return seg[node]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
