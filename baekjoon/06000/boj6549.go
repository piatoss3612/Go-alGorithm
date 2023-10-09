package main

import (
	"bufio"
	"fmt"
	"os"
	_ "sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	inp     []int
	n       int
)

// 메모리: 6016KB
// 시간: 68ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		n = scanInt()
		if n == 0 {
			return
		}

		inp = make([]int, n+1)

		for i := 1; i <= n; i++ {
			inp[i] = scanInt()
		}
		fmt.Fprintln(writer, DAC(1, n))
	}
}

// 분할 정복
func DAC(left, right int) int {
	// 직사각형의 너비가 1인 경우
	if left == right {
		return inp[left]
	}

	// 가장 큰 직사각형이 왼쪽 또는 오른쪽에 있는 경우
	mid := (left + right) / 2
	ret := max(DAC(left, mid), DAC(mid+1, right))

	// 가장 큰 직사각형이 mid를 기준으로 좌우로 걸쳐져있는 경우
	segL, segR := mid, mid+1
	squareHeight := min(inp[segL], inp[segR]) // 직사각형을 만들 수 있는 최대 높이
	ret = max(ret, squareHeight*2)

	// 왼쪽 또는 오른쪽으로 1씩 이동하며 직사각형 넓이의 최댓값 갱신
	for left < segL || segR < right {
		// 오른쪽으로 이동하는 경우
		if segR < right && (segL == left || inp[segL-1] < inp[segR+1]) {
			segR += 1
			squareHeight = min(squareHeight, inp[segR])
		} else { // 왼쪽으로 이동하는 경우
			segL -= 1
			squareHeight = min(squareHeight, inp[segL])
		}

		/*
			#높이가 더 높은 방향으로 먼저 이동하는 이유
			1. 직사각형의 최대 높이를 유지하면서 넓이가 1*최대 높이만큼 증가
			2. 최대 높이가 줄어드는 것을 최소화
		*/

		ret = max(ret, (segR-segL+1)*squareHeight) // 최댓값 갱신
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
