package bj9372

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	tree    [][]int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	n, m := scanInt(), scanInt()
	tree = make([][]int, n+1)
	for j := 0; j < m; j++ {
		a, b := scanInt(), scanInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	visited = make([]bool, n+1)     // 방문 여부 확인
	fmt.Fprintln(writer, DFS(1, 0)) // 1번 정점부터 순회

	// 모든 정점이 연결되어 있으므로
	// 신장 트리의 경우로 생각하면 정점의 갯수 - 1이 답이된다
}

func DFS(idx, cnt int) int {
	if visited[idx] == true { // 이미 방문했다면, cnt + 1을 받았으므로 cnt - 1을 반환한다
		return cnt - 1
	}
	visited[idx] = true
	for _, v := range tree[idx] { // idx에 해당하는 정점에 연결된 정점들을 순회
		cnt = DFS(v, cnt+1)
	}
	return cnt // 순회한 횟수를 반환
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
