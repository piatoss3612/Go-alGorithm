package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	routes    [201][201]int // 경로가 있는 경우 1, 없는 경우 0
	reachable [201]int      // i번 도시로 오는데 경유하는 루트 도시
	plan      []int         // 여행 계획
	N, M      int
)

// 메모리: 1220KB
// 시간: 8ms
// 그래프 탐색, 분리 집합 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		reachable[i] = i // 루트를 자기자신을 초기화
		for j := 1; j <= N; j++ {
			routes[i][j] = scanInt()
		}
	}

	// 그래프는 (i,i) 대각선을 두고 대칭인 그래프
	// 따라서 절반만 탐색해도 된다
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			if routes[i][j] == 1 { // i에서 j로 경로가 있는 경우
				union(i, j)
			}
		}
	}

	plan = make([]int, M+1)
	for i := 1; i <= M; i++ {
		plan[i] = scanInt()
	}

	// 예외: 방문 계획 도시가 1개인 경우는 항상 여행 계획을 완수할 수 있다
	if M == 1 {
		fmt.Fprintln(writer, "YES")
		return
	}

	possible := true // flag
	// 여행 계획 2번째 도시부터 이전 도시와 비교
	for i := 2; i <= M; i++ {
		x, y := find(plan[i-1]), find(plan[i])
		// i-1번째로 계획한 도시와 i번째로 계획한 도시가 동일한 루트 도시를 가지지 않는 경우
		if x != y {
			possible = false
			break // 여행 계획은 불가능하다고 판명, 반복문 탈출
		}
	}

	if possible {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func find(x int) int {
	if reachable[x] == x {
		return x
	}
	reachable[x] = find(reachable[x])
	return reachable[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		// 반복문을 통해 넘겨받는 x는 항상 y보다 작은 값이다
		reachable[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
