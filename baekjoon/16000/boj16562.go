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
	n, m, k int
	expense []int // 친구 비용
	parent  []int // 친구 그룹의 리더가 누구인지
)

// 메모리: 1172KB
// 시간: 8ms
// 친구비를 내면서까지 친구가 필요해?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, k = scanInt(), scanInt(), scanInt()
	expense = make([]int, n+1)
	parent = make([]int, n+1)

	for i := 1; i <= n; i++ {
		expense[i] = scanInt() // i와 친구가 되기 위해 필요한 비용;
		parent[i] = i          // i가 속한 친구 그룹의 리더를 i로 초기화
	}

	// 분리 집합
	var v, w int
	for i := 1; i <= m; i++ {
		v, w = scanInt(), scanInt()
		union(v, w) // v와 w는 같은 친구 그룹에 속하므로 union
	}

	total := 0
	// i번째 친구가 속한 그룹의 리더가 i인 경우
	// 즉, i에게 친구비용을 지불하면
	// 가장 적은 비용으로 i가 속한 그룹의 다른 사람들과도 친구가 될 수 있다
	for i := 1; i <= n; i++ {
		if parent[i] == i {
			total += expense[i]
		}
	}

	// 전체 비용이 가진 돈 k보다 많은 경우
	if total > k {
		fmt.Fprintln(writer, "Oh no")
	} else {
		fmt.Fprintln(writer, total)
	}
}

// x가 속한 친구 그룹의 리더를 find
func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

// x와 y를 같은 친구 그룹으로 union
func union(x, y int) {
	x, y = find(x), find(y) // x가 속한 그룹의 리더, y가 속한 그룹의 리더를 찾는다
	// 각각의 리더가 다르다면
	if x != y {
		// 친구 비용이 더 작은 리더가 전체 그룹의 리더가 된다
		if expense[x] > expense[y] {
			parent[x] = y
		} else {
			parent[y] = x
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
