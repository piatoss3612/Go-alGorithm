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
	set     []int
	n, m    int
)

// 메모리: 11004KB
// 시간: 68ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	// 0~n까지 집합의 부모를 원소 자신으로 초기화
	set = make([]int, n+1)
	for i := 1; i <= n; i++ {
		set[i] = i
	}

	var ops, a, b int
	for i := 1; i <= m; i++ {
		ops, a, b = scanInt(), scanInt(), scanInt()

		// 연산자 swtich문
		switch ops {
		case 0:
			union(a, b) // a가 포함된 집합과 b가 포함된 집합을 합친다
		case 1:
			inSameSet(a, b) // a와 b가 같은 집합에 포함되어 있는지 확인한다
		}
	}
}

func find(x int) int {
	if set[x] == x {
		return x
	}
	set[x] = find(set[x])
	return set[x]
}

func union(x, y int) {
	// x가 속한 집합과 y가 속한 집합을 합치기 위해
	// x와 y의 부모를 찾아 x, y에 재할당한다
	x, y = find(x), find(y)
	if x != y {
		set[y] = x
	}
}

func inSameSet(x, y int) {
	if find(x) == find(y) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
