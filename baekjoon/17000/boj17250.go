package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, M       int
	galaxy     []int // 연결된 행성들의 수
	connection []int // 연결 정보
)

// 메모리: 5468KB
// 시간: 76ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	galaxy = make([]int, N+1)
	connection = make([]int, N+1)
	for i := 1; i <= N; i++ {
		galaxy[i] = scanInt()
		connection[i] = i
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		union(a, b)
	}
}

// 루트 요소 찾기
func find(x int) int {
	if connection[x] == x {
		return x
	}
	connection[x] = find(connection[x])
	return connection[x]
}

// x, y 유니온
func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		galaxy[x] += galaxy[y] // 연결된 행성의 수 갱신
		galaxy[y] = galaxy[x]
		connection[y] = x // y의 부모 요소를 x로 변경
	}

	fmt.Fprintln(writer, galaxy[x]) // 연결된 행성의 수 출력
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
