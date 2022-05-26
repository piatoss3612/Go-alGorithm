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
	minTree [300001]int // 최솟값 세그먼트 트리
	maxTree [300001]int // 최댓값 세그먼트 트리
	n, m    int
)

// 메모리: 10696KB
// 시간: 144ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	segmentTree(1, n, 1)

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		x, y := query(a, b, 1, 1, n)
		fmt.Fprintln(writer, x, y)
	}
}

// left~right 범위에 해당하는 최소, 최댓값을 찾는 쿼리 함수
func query(left, right, node, nodeLeft, nodeRight int) (int, int) {
	// 범위를 완전히 벗어난 경우
	if nodeLeft > right || nodeRight < left {
		return 1000000001, 0 // 최솟값과 최댓값으로 범위보다 큰 값, 범위보다 작은 값 반환
	}

	// left <= nodeLeft <= nodeRight <= right인 경우
	// node가 가리키는 범위가 left~right 범위에 완전히 포함되는 경우 (교집합 상태)
	if left <= nodeLeft && nodeRight <= right {
		return minTree[node], maxTree[node]
	}

	// 왼쪽, 오른쪽 구간을 나누어 탐색한 값을 비교
	mid := (nodeLeft + nodeRight) / 2
	a, b := query(left, right, node*2, nodeLeft, mid)
	c, d := query(left, right, node*2+1, mid+1, nodeRight)
	return min(a, c), max(b, d)
}

/*
# segmentTree:

left: 구간의 시작 위치
right: 끝나는 위치
node: 트리의 인덱스

minTree예시:

5
30 5
30 38 20 5
30 100 38 50 51 20 81 5
75 30 0 0 0 0 0 0 51 52 0 0 0 0

주어진 입력값의 구간을 나누어 이러한 형태의 이진 트리를 1차원 배열에 저장한다
세그먼트 트리는 거의 가득찬 이진 트리(full binary tree)이므로 1차원 배열을 사용하는 것이 메모리 효율적이다
*/
func segmentTree(left, right, node int) (int, int) {
	// 길이가 1인 구간
	if left == right {
		minTree[node] = inp[left]
		maxTree[node] = inp[left]
		return minTree[node], maxTree[node]
	}

	// 왼쪽, 오른쪽으로 구간 나누어 탐색한 값을 비교
	mid := (left + right) / 2
	leftMin, leftMax := segmentTree(left, mid, node*2)
	rightMin, rightMax := segmentTree(mid+1, right, node*2+1)

	minTree[node] = min(leftMin, rightMin) // left~right 구간의 최솟값
	maxTree[node] = max(leftMax, rightMax) // left~right 구간의 최댓값
	return minTree[node], maxTree[node]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
