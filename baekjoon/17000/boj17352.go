package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	group   []int
)

// 메모리: 14568KB
// 시간: 132ms
// 분리 집합

// N-1개로 연결된 섬들은 본래 1개의 집합이었을 테지만
// 다리를 1개 부숴버린 지금은 반드시 2개의 집합으로 나뉘었을 것이고
// 이 2개의 집합 각각 속한 아무 섬들 중에 2개의 섬은 연결해야
// 어떤 두 섬 사이든 왕복할 수 있게 만들 수 있다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	N = scanInt()
	group = make([]int, N+1)
	for i := 1; i <= N; i++ {
		group[i] = i
	}

	for i := 1; i <= N-2; i++ {
		x, y := scanLine()
		union(x, y)
	}

	for i := 1; i <= N; i++ {
		if find(i) == i {
			fmt.Fprintf(writer, "%d ", i)
		}
	}
	fmt.Fprintln(writer)
}

func find(x int) int {
	if group[x] == x {
		return x
	}
	group[x] = find(group[x])
	return group[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		group[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanLine() (int, int) {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	return x, y
}
