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
	n, m    int
	parent  []int
)

// 메모리: 4144KB
// 시간: 40ms
// 분리 집합, 모든 뉴런을 트리 형태로 연결하려면, 사이클은 끊고 떨어져 있는 것들은 연결해야 한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	cycle := 0 // 사이클 찾기

	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		cycle += union(a, b)
	}

	// 뉴런의 분리 집합의 개수 찾기
	cnt := 0
	for i := 1; i <= n; i++ {
		if find(i) == i {
			cnt++
		}
	}

	// 모든 뉴런을 하나의 트리 형태로 연결하기 위한 최소 연산 횟수
	// 사이클을 끊는 횟수 + (분리된 집합의 수 - 1, 즉 집합들을 연결하기 위한 최소 시냅스의 수)
	fmt.Fprintln(writer, cycle+cnt-1)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) int {
	x, y = find(x), find(y)
	if x != y {
		parent[y] = x
		return 0
	}

	// x와 y가 같으면 사이클을 이룬다는 의미
	return 1
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
