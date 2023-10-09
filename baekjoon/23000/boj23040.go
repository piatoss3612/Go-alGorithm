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
	N       int
	parent  []int // 정점 i가 속한 그룹(검은색, 빨간색)의 부모 정점
	family  []int // 정점 i가 속한 그룹(검은색, 빨간색)의 정점의 개수
	tree    [][]int
	visited []bool
	color   []byte
)

// 난이도: Gold 3
// 메모리: 26728KB
// 시간: 92ms
// 분류: 트리, 분리 집합, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, 200000), 200000) // color 슬라이스의 길이가 최대 100001이므로 버퍼 크기를 늘려줘야 한다
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	tree = make([][]int, N+1)
	parent = make([]int, N+1)
	family = make([]int, N+1)
	visited = make([]bool, N+1)

	for i := 1; i <= N-1; i++ {
		u, v := scanInt(), scanInt()
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
		parent[i] = i
		family[i] = 1
	}
	parent[N] = N
	family[N] = 1
	color = scanBytes()
}

func Solve() {
	DFS(1) // 사이클이 존재하지 않는 트리의 임의의 정점(1)에서 깊이 우선 탐색 실행

	ans := 0
	for i := 1; i <= N; i++ {
		// 정점 i가 검은색인 경우
		if color[i] == 'B' {
			for _, next := range tree[i] {
				// 검은색 정점 i와 연결된 정점 next가 빨간색인 경우
				if color[next] == 'R' {
					ans += family[Find(next)] // 정점 next가 속한 그룹의 부모 정점을 찾아 그룹에 속한 빨간 정점의 개수(=경로의 수)를 더한다
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func DFS(x int) {
	visited[x] = true

	for _, next := range tree[x] {
		if !visited[next] {
			// 정점 x 및 x와 연결된 정점 next가 모두 빨간색인 경우
			if color[x] == 'R' && color[next] == 'R' {
				Union(x, next) // 정점 x와 next를 유니온
			}

			DFS(next) // 정점 next에서 깊이 우선 탐색
		}
	}
}

// 정점 x의 부모 정점을 찾아 반환한다
func Find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

// 정점 x와 정점 y가 속한 그룹을 유니온한다
func Union(x, y int) {
	x, y = Find(x), Find(y)
	if x != y {
		family[x] += family[y]
		family[y] = family[x]
		parent[y] = x
	}
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	b := scanner.Bytes()
	res := make([]byte, 0, len(b)+1)
	res = append(res, 0)
	return append(res, b...)
}
