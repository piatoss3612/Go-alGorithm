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
	islands []Island
	tree    [][]int
	dp      []int
)

type Island struct {
	isWolf bool
	amount int
}

// 난이도: Gold 3
// 메모리: 27472KB
// 시간: 104ms
// 분류: 트리, 깊이 우선 탐색, 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	islands = make([]Island, N+1)
	tree = make([][]int, N+1)
	dp = make([]int, N+1)

	// 2번째 섬부터 N번째 섬까지의 정보 입력
	for i := 2; i <= N; i++ {
		t, a, p := scanByte(), scanInt(), scanInt()
		islands[i] = Island{t == 'W', a}
		tree[p] = append(tree[p], i)
	}
}

func Solve() {
	fmt.Fprintln(writer, DFS(1)) // 1번 섬에서 깊이 우선 탐색 실행
}

// 깊이 우선 탐색을 통해 x번째 섬에	있는 양들의 수를 구한다
func DFS(x int) int {
	// x번째 섬과 연결된 다른 섬으로부터 x번째 섬으로 이동하는 양들의 수 구하기
	for _, next := range tree[x] {
		dp[x] += DFS(next)
	}

	// x번째 섬이 늑대들의 섬인 경우
	if islands[x].isWolf {
		dp[x] = max(0, dp[x]-islands[x].amount) // 늑대가 양을 잡아먹는다!
	} else {
		dp[x] += islands[x].amount // x번째 섬의 양들도 구명보트를 타고 이동한다
	}

	return dp[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}

func scanByte() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}
