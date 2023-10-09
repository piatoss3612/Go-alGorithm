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
	parent  []int
	n, m    int
)

// 메모리: 24172KB
// 시간: 416ms
// 선분을 연결하여 사이클 찾기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	parent = make([]int, n)
	for i := 1; i < n; i++ {
		parent[i] = i
	}

	cnt := 987654321 // 사이클이 몇 번째 차례에서 나타났는지 비교하기 위한 값

	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		pa, pb := find(a), find(b)
		// a와 b의 부모가 같다면 사이클이 형성되었다는 것을 의미한다
		if pa == pb {
			cnt = min(cnt, i)
		} else {
			// 부모가 다른 경우는 union 작업
			parent[pb] = pa
		}
	}

	// 사이클을 발견하지 못한 경우
	if cnt == 987654321 {
		fmt.Fprintln(writer, 0)
	} else {
		// 사이클이 최초로 발생한 차례를 출력
		fmt.Fprintln(writer, cnt)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
