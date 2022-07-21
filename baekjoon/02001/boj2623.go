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
	graph    [][]int // 연결 상태
	inDegree []int   // 진입차수
)

// 메모리: 1048KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	graph = make([][]int, n+1)
	inDegree = make([]int, n+1)

	// singers: i번째 보조 PD가 담당하는 가수의 수
	// prev: current보다 먼저 출연하는 가수의 번호
	// current: 새롭게 입력받는 가수의 번호
	// 그래프는 prev -> current, prev가 current로 진입하는 형태
	var singers, prev, current int
	for i := 1; i <= m; i++ {
		singers = scanInt()
		for j := 1; j <= singers; j++ {
			current = scanInt()
			if j != 1 {
				graph[prev] = append(graph[prev], current)
				inDegree[current] += 1
			}
			prev = current
		}
	}

	topologicalSort()
}

// 위상 정렬
func topologicalSort() {
	var res, queue []int

	// 진입차수가 0인 가수의 번호를 큐에 추가
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		x := queue[0] // pop
		queue = queue[1:]
		res = append(res, x)

		// x번 가수의 다음 순서가 될 수 있는 가수들을 탐색
		for _, v := range graph[x] {
			inDegree[v] -= 1
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// 모든 가수들을 순서대로 정렬하지 못한 경우
	if len(res) != n {
		fmt.Fprintln(writer, 0)
		return
	}

	// 모든 가수들을 순서대로 정렬한 경우
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
