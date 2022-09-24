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
	N, S, D  int
	tree     [][]int // 입력으로 주어지는 트리
	depth    []int   // 트리의 각 노드의 깊이
	maxDepth []int   // 각 노드가 속한 경로의 최대 깊이
	ans      = 0     // 이동해야 하는 거리
)

// 메모리: 19404KB
// 시간: 72ms
// 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, S, D = scanInt(), scanInt(), scanInt()

	// 1. 트리 입력
	tree = make([][]int, N+1)
	for i := 1; i < N; i++ {
		x, y := scanInt(), scanInt()
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}

	// 2. 깊이, 최대 깊이 슬라이스 초기화
	depth = make([]int, N+1)
	maxDepth = make([]int, N+1)
	for i := 1; i <= N; i++ {
		depth[i] = -1
	}

	// 3. 노드 S에서 시작하여 트리의 깊이, 최대 깊이 구하기
	treeDepth(S, 0)

	// 4. 노드 S에서 시작하여 이동해야 하는 최단 거리 구하기
	solve(S)

	// 5. 결과 출력
	fmt.Fprintln(writer, ans*2) // 이동해야 하는 최단 거리는 왕복하는 거리이므로 *2를 해준다
}

// 3. 깊이, 최대 깊이 구하기
func treeDepth(x, d int) int {
	depth[x] = d // 노드 x의 깊이를 d로 초기화

	md := d // 최대 깊이

	for _, next := range tree[x] {
		if depth[next] == -1 {
			md = max(md, treeDepth(next, d+1)) // 최대 깊이 비교
		}
	}

	maxDepth[x] = md // 노드 x가 속한 경로의 최대 깊이 초기화

	return md
}

// 4. 이동해야 하는 최단 거리 구하기
func solve(x int) {
	for _, child := range tree[x] {
		// 노드 x의 자식 노드만 탐색하기 위해서 깊이 비교
		if depth[child] > depth[x] {
			// 자식 노드가 속한 경로의 최대 깊이가 노드 x의 깊이와 D를 더한 값보다 큰 경우
			// 즉, 노드 x가 속한 경로에 속한 모든 노드에 전단지를 돌리지 못한 경우
			if maxDepth[child]-depth[x] > D {
				ans++        // 자식 노드로 이동해야 하므로 이동 거리 증가
				solve(child) // 자식 노드로 이동하여 남은 노드에 전단지를 돌린다
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
