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
	n, m     int
	inDegree []int      // 진입 차수
	graph    []building // 건물을 짓는데 필요한 정보
	dp       []int      // 건물이 완성되기까지 걸리는 최소 시간
)

type building struct {
	time int   // 건물을 짓기위해 걸리는 시간
	prev []int // 선행하여 지어야 하는 건물들
}

// 메모리: 2308KB
// 시간: 12ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()

	graph = make([]building, n+1)
	inDegree = make([]int, n+1)
	dp = make([]int, n+1)

	for i := 1; i <= n; i++ {
		graph[i].time = scanInt()
		precede := scanInt()
		for precede != -1 {
			graph[i].prev = append(graph[i].prev, precede)
			inDegree[precede] += 1
			precede = scanInt()
		}
	}

	topologicalSort()

	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, dp[i])
	}
}

// 위상 정렬
func topologicalSort() {
	var q []int

	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		_ = rec(x) // 진입 차수가 0인 x번 건물을 짖기 위해 필요한 시간의 최솟값 구하기

		for i := 0; i < len(graph[x].prev); i++ {
			y := graph[x].prev[i]
			inDegree[y] -= 1
			if inDegree[y] == 0 {
				q = append(q, y)
			}
		}
	}
}

func rec(x int) int {
	ret := &dp[x]
	if *ret != 0 {
		return *ret
	}

	*ret = graph[x].time // 선행하여 지어야 할 건물이 없는 경우, x번 건물을 짓기 위해 필요한 시간을 반환
	for i := 0; i < len(graph[x].prev); i++ {
		/*
			선행하여 지어야 할 건물을 짓는데 걸리는 최소 시간 + 건물 x를 짓는데 걸리는 시간의 최댓값
			아니 최솟값을 구하라는데 왜 최댓값을 구하느냐???

			예제 입력:
			5
			10 -1
			10 1 -1
			4 1 -1
			4 3 1 -1
			3 3 -1

			풀이:

			4번 건물을 짓는데 필요한 최소 시간을 알아보자
			4번 건물은 3번 건물과 1번 건물을 선행하여 지어야 한다

			3번 건물을 짓는데 필요한 최소 시간은 3번 건물 + 1번 건물 = 14 시간 이다
			1번 건물을 짓는데 필요한 최소 시간은 1번 건물 = 10 시간 이다

			여기서 3번 건물은 1번 건물을 선행해서 지어야 하므로
			결과적으로 4번 건물은 1->3->4 순서로 지어져 17 시간이 걸리게 된다

			즉, 선행하는 조건들을 고려하다 보면 다른 선행 조건들을 포괄적으로 포함할 수 있는,
			최솟값(먼저 지어야 하는 건물을 짓는데 걸리는) 중에 최댓값을 구해야 하는 것이다
		*/
		*ret = max(*ret, rec(graph[x].prev[i])+graph[x].time)
	}

	return *ret
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
