package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	supplyTo [26][]int // supplyTo[i]: i번 공급책으로부터 공급받는 공급책들
	checked  [26]bool  // 깊이 우선 탐색 여부
	inDegree [26]int   // i번 공급책에서 공급해주는 공급책의 수, 0인 경우 최상위 공급책
	ans      int       // 일부 공급책이 검거되고도 약을 공급받을 수 있는 공급책의 수
)

// 난이도: Gold 4
// 메모리: 920KB
// 시간: 4ms
// 분류: 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		from, to := scanSupplier(), scanSupplier()
		supplyTo[from] = append(supplyTo[from], to)
		inDegree[to]++
	}

	detected := scanInt()
	for i := 1; i <= detected; i++ {
		s := scanSupplier()
		checked[s] = true // 검거된 공급책은 깊이 우선 탐색에서 제외
	}
}

func Solve() {
	// 최상위 공급책 찾기
	rootSuppliers := []int{}
	for i := 0; i < 26; i++ {
		if !checked[i] && inDegree[i] == 0 {
			rootSuppliers = append(rootSuppliers, i)
		}
	}

	// 최상위 공급책들로터 깊이 우선 탐색
	for _, supplier := range rootSuppliers {
		DFS(supplier)
	}

	fmt.Fprintln(writer, ans)
}

func DFS(x int) {
	checked[x] = true

	for _, next := range supplyTo[x] {
		if !checked[next] {
			DFS(next)
			ans++
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanSupplier() int {
	scanner.Scan()
	s := int(scanner.Bytes()[0] - 'A')
	return s
}
