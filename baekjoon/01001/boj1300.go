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
	n, k    int
)

// 메모리: 904KB
// 시간: 36ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	k = scanInt()

	// 배열 B의 k번째 수는 항상 k보다 작거나 같으므로 n*n이 아닌 k를 right로 설정하고 이분 탐색함으로써 탐색 시간을 단축할 수 있다
	l, r := 0, k

	// 이분 탐색 - lower bound: getOrder(m)이 k보다 크거나 같은 최솟값 m을 구한다
	for l < r {
		m := (l + r) / 2
		if getOrder(m) >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	fmt.Fprintln(writer, r)
}

/*
	# 예제 입력:
	3
	7

	# 2차원 배열 A:
	1 2 3
	2 4 6
	3 6 9

	# 1차원 배열 B:
	1 2 2 3 3 4 6 6 9

	# m 이하의 수의 개수 - getOrder(m):
	getOrder(1) = 1
	getOrder(2) = 3
	getOrder(3) = 5
	getOrder(4) = 6
	getOrder(5) = 6
	getOrder(6) = 8
	getOrder(7) = 8
	getOrder(8) = 8
	getOrder(9) = 9

	getOrder(4)와 getOrder(5)의 값이 같은 이유는 배열 B에 5라는 값이 존재하지 않기 때문
	마찬가지로 배열 B에는 7과 8이 존재하지 않는다

	# 배열 B의 k번째 수
	배열 B의 k번째 수는 곧 getOrder(m)이 k보다 크거나 같은 m의 최솟값이다

	getOrder(6)은 8이고 getOrder(4)는 6으로 4이하의 수가 배열 B의 6번째 자리까지 차지하고 있고
	그다음으로 6이하의 수가 8번째 자리까지 차지하고 있는데

	getOrder(4)와 getOrder(5)의 값이 같기 때문에 배열 B에는 5가 존재하지 않는다는 것을 알 수 있다

	따라서 배열 B의 7번째, 8번째 자리를 6이 차지하고 있다는 것을 알 수 있다


	참고 - https://www.acmicpc.net/board/view/37110
*/

// n*n 2차원 배열 A에서 v이하의 수의 개수를 구하는 함수
func getOrder(v int) int {
	temp := 0
	for i := 1; i <= n; i++ {
		temp += min(v/i, n)
	}
	return temp
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
