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
	inp     [100001]int
	seg     [270000]int
	n, q    int
)

// 메모리: 9068KB
// 시간: 184ms
// 2042번과 유사한 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, q = scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	segmentTree(1, n, 1)

	var x, y, a, b int
	for i := 1; i <= q; i++ {
		x, y, a, b = scanInt(), scanInt(), scanInt(), scanInt()

		// x가 y보다 큰 경우는 y~x구간을 탐색
		if x > y {
			fmt.Fprintln(writer, query(y, x, 1, n, 1))
		} else {
			fmt.Fprintln(writer, query(x, y, 1, n, 1))
		}

		updateTree(a, b-inp[a], 1, n, 1) // 구간합 갱신
		inp[a] = b                       // 입력값 갱신
	}
}

func query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || left > nodeRight {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return query(left, right, nodeLeft, mid, node*2) + query(left, right, mid+1, nodeRight, node*2+1)
}

func updateTree(idx, change, left, right, node int) {
	if idx < left || idx > right {
		return
	}

	if left == idx && right == idx {
		seg[node] += change
		return
	}

	if left <= idx && idx <= right {
		seg[node] += change
		mid := (left + right) / 2
		updateTree(idx, change, left, mid, node*2)
		updateTree(idx, change, mid+1, right, node*2+1)
		return
	}
}

func segmentTree(left, right, node int) int {
	if left == right {
		seg[node] = inp[left]
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = segmentTree(left, mid, node*2) + segmentTree(mid+1, right, node*2+1)
	return seg[node]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
