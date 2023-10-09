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
	seg     [270000]int
	inp     [100001]int
	N, M    int
)

// 메모리: 6976KB
// 시간: 76ms
// 세그먼트 트리를 사용해 최솟값의 인덱스 찾기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}

	M = scanInt()

	minIdx := SegmentTree(1, N, 1) // 세그먼트 트리 초기화 및 최솟값 인덱스 찾기

	for i := 1; i <= M; i++ {
		q := scanInt()

		switch q {
		case 1:
			a, b := scanInt(), scanInt()
			inp[a] = b                  // a번째 인덱스의 값을 b로 업데이트
			minIdx = Update(a, 1, N, 1) // 세그먼트 트리 및 최솟값 인덱스 업데이트
		case 2:
			fmt.Fprintln(writer, minIdx) // 최솟값 인덱스 출력
		}
	}
}

func SegmentTree(left, right, node int) int {
	if left == right {
		seg[node] = left
		return seg[node]
	}

	mid := (left + right) / 2
	onLeft := SegmentTree(left, mid, node*2)
	onRight := SegmentTree(mid+1, right, node*2+1)

	// 왼쪽과 오른쪽 최솟값 인덱스의 값 비교
	if inp[onLeft] <= inp[onRight] {
		seg[node] = onLeft
	} else {
		seg[node] = onRight
	}
	return seg[node]
}

func Update(target, left, right, node int) int {
	// 값이 변경된 인덱스가 범위에 포함되지 않는 경우
	if target < left || right < target {
		return seg[node]
	}

	// 값이 변경된 인덱스 자기자신인 경우
	if left == right {
		return seg[node]
	}

	mid := (left + right) / 2
	onLeft := Update(target, left, mid, node*2)
	onRight := Update(target, mid+1, right, node*2+1)

	// 왼쪽과 오른쪽 최솟값 인덱스의 값 비교
	if inp[onLeft] <= inp[onRight] {
		seg[node] = onLeft
	} else {
		seg[node] = onRight
	}
	return seg[node]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
