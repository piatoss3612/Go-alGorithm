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

// 메모리: 30728KB
// 시간: 236ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, k = scanInt(), scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	segmentTree(1, n, 1)

	var a, b, c int
	for i := 1; i <= m+k; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()

		if a == 1 {
			updateTree(b, c-inp[b], 1, n, 1)
			inp[b] = c
		} else {
			fmt.Fprintln(writer, query(b, c, 1, n, 1))
		}
	}
}

// left~right 구간의 합을 구하는 함수
func query(left, right, nodeLeft, nodeRight, node int) int {
	// left~right 범위가 탐색 범위에 포함되지 않는 경우
	if right < nodeLeft || left > nodeRight {
		return 0
	}

	// left~right 범위가 탐색 범위에 완전히 포함되는 경우
	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	// 왼쪽, 오른쪽 구간 나누기
	mid := (nodeLeft + nodeRight) / 2
	return query(left, right, nodeLeft, mid, node*2) + query(left, right, mid+1, nodeRight, node*2+1)
}

// 세그먼트 트리에서 idx번째 입력값이 포함된 모든 누적합을 갱신
func updateTree(idx, change, left, right, node int) {
	// 범위를 벗어난 경우
	if idx < left || idx > right {
		return
	}

	// idx번째 입력값으로만 이루어진 세그먼트 값인 경우
	// 이부분 체크 안해주면 runtime error: index out of range가 발생
	if left == idx && right == idx {
		seg[node] += change
		return
	}

	// idx가 세그먼트 범위에 포함되는 경우
	if left <= idx && idx <= right {
		seg[node] += change
		mid := (left + right) / 2
		updateTree(idx, change, left, mid, node*2)
		updateTree(idx, change, mid+1, right, node*2+1)
		return
	}
}

// 입력값으로 구간합 세그먼트 트리를 구하는 함수
func segmentTree(left, right, node int) int {
	// 길이가 1인 구간합
	if left == right {
		seg[node] = inp[left]
		return seg[node]
	}

	// 왼쪽, 오른쪽 구간 나누기
	mid := (left + right) / 2
	seg[node] = segmentTree(left, mid, node*2) + segmentTree(mid+1, right, node*2+1)
	return seg[node]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
