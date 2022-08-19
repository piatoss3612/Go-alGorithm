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
	graph   [201][201]int
	friends [201]int
	N, M, K int
	INF     = 987654321
)

func init() {
	for i := 1; i <= 200; i++ {
		for j := 1; j <= 200; j++ {
			if i == j {
				continue
			}
			graph[i][j] = INF
		}
	}
}

// 메모리: 1588KB
// 시간: 32ms
// 플로이드 와샬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		a, b, t := scanInt(), scanInt(), scanInt()
		graph[a][b] = t
	}

	// 플로이드 와샬 알고리즘으로 최단 경로를 먼저 탐색한다
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	K = scanInt()
	for i := 1; i <= K; i++ {
		friends[i] = scanInt()
	}

	// 준형이와 친구들의 왕복시간 들 중 최대가 최소가 되는 도시 X를 선택한다.
	// 내가 심한 난독증인가 처음에 이게 무슨 소린가 한참 생각했다

	minVal := INF      // 왕복시간 중 최대가 최소인 값
	minList := []int{} // 왕복시간 중 최대가 최소일 때 이동할 수 있는 도시들

	for i := 1; i <= N; i++ {
		// 준형이와 친구들의 왕복시간 중 최댓값 구하기
		temp := 0
		for j := 1; j <= K; j++ {
			temp = max(temp, graph[friends[j]][i]+graph[i][friends[j]])
		}

		// 왕복시간 중 최대가 최소가 되는 경우
		if temp < minVal {
			minVal = temp      // 최솟값 갱신
			minList = []int{i} // 이동할 수 있는 도시 리스트 새로 생성
		} else if temp == minVal {
			minList = append(minList, i)
		}
	}

	for _, v := range minList {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
