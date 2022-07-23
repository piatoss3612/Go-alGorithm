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
	gate    []int // 1~gi번 게이트 중에 도킹이 가능한 게이트의 수 (1 <= gi <= G)
	target  []int // pi번 비행기가 도킹하고자 하는 게이트의 번호
	G, P    int
)

// 메모리: 3324KB
// 시간: 20ms
// 분리 집합의 find 함수를 사용하여 1~gi번 게이트 중 도킹할 수 있는 게이트의 수 찾기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	G, P = scanInt(), scanInt()

	gate = make([]int, G+1)
	for i := 1; i <= G; i++ {
		gate[i] = i // 1~gi번 게이트에는 최대 gi개의 비행기가 도킹할 수 있다
	}

	target = make([]int, P+1)
	for i := 1; i <= P; i++ {
		target[i] = scanInt()
	}

	cnt := 0

	for i := 1; i <= P; i++ {
		// i번째 비행기가 도킹하려는 게이트 범위(1~target[i])에서
		// 도킹이 가능한 게이트의 수
		dockable := find(target[i])

		// 도킹이 가능한 게이트가 없다면
		if dockable == 0 {
			break
		}

		// 게이트의 범위(1~target[i])에서 도킹이 가능한 게이트의 수 갱신
		gate[dockable] = find(gate[dockable] - 1)
		cnt++
	}

	fmt.Fprintln(writer, cnt)
}

// 1~x번 게이트 범위에서 도킹가능한 게이트의 수 찾기
func find(x int) int {
	if gate[x] == x {
		return x
	}
	gate[x] = find(gate[x])
	return gate[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
