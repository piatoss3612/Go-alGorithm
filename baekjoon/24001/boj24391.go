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
	N, M    int
	conn    []int // 건물 연결 정보
)

// 메모리: 3452KB
// 시간: 60ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	conn = make([]int, N+1)
	for i := 1; i <= N; i++ {
		conn[i] = i
	}

	for i := 1; i <= M; i++ {
		x, y := scanInt(), scanInt()
		union(x, y)
	}

	move := 0

	// conn[from]과 conn[to]가 서로다른 값일 때
	// from 건물에서 to 건물로 이동하기 위해 밖으로 나와야 한다
	from := scanInt()
	for i := 1; i < N; i++ {
		to := scanInt()
		if find(from) != find(to) {
			move++
		}
		from = to
	}

	fmt.Fprintln(writer, move)
}

func find(x int) int {
	if conn[x] == x {
		return x
	}
	conn[x] = find(conn[x])
	return conn[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		conn[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
