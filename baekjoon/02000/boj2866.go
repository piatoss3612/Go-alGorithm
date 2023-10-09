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
	R, C    int
	strs    [][]byte
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	R, C = scanInt(), scanInt()

	// 테이블의 열 방향으로 문자열 전처리
	strs = make([][]byte, C)
	for i := 0; i < C; i++ {
		strs[i] = make([]byte, 0, R)
	}

	for i := 0; i < R; i++ {
		inp := scanBytes()
		for j := 0; j < C; j++ {
			strs[j] = append(strs[j], inp[j])
		}
	}

	// 이분 탐색
	l, r := 0, R-1

	for l <= r {
		mid := (l + r) / 2

		// 각 문자열의 mid 번째 문자부터 시작하는 동일한 부분 문자열이 있는지 체크
		check := make(map[string]bool)
		flag := true

		for i := 0; i < C; i++ {
			if _, ok := check[string(strs[i][mid:])]; ok {
				// 동일한 부분 문자열이 있다면
				// flag 값을 false로 변경하고 반복문 종료
				flag = false
				break
			} else {
				// 동일한 부분 문자열이 없다면
				// 맵에 현재 문자열 추가
				check[string(strs[i][mid:])] = true
			}
		}

		// 동일한 부분 문자열이 없는 경우
		if flag {
			l = mid + 1 // l을 앞으로 이동
		} else {
			// 동일한 부분 문자열이 있는 경우
			r = mid - 1 // r을 뒤로 이동
		}
	}

	// l은 문자열에서 동일한 부분 문자열이 없는 경우(mid)의 최댓값 + 1
	// 따라서 count값은 l-1이다
	fmt.Fprintln(writer, l-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
