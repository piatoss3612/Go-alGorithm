package bj1389

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		graph[a][b] = 1
		graph[b][a] = 1
	}
	gameResult := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tmp := 0 // 정점 i의 케빈 베이컨 수
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			tmp += baconGame(i, j)      // 정점 i가 j와 케빈 베이컨 게임을 했을 때 나오는 단계를 더해준다
			visited = make([]bool, n+1) // 방문 여부 초기화
		}
		gameResult[i] = tmp // 케빈 베이컨 수를 저장
	}
	min := n * n
	idx := 0
	for i := 1; i <= n; i++ {
		if gameResult[i] < min { // 최솟값 구하기
			min = gameResult[i]
			idx = i
		}
	}
	fmt.Fprintln(writer, idx)
}

func baconGame(start, target int) int {
	visited[start] = true
	cnt := 0                       // 케빈 베이컨 단계
	queue := [][]int{{start, cnt}} // 이중 배열로 된 큐에 정점과 해당 정점에서 케빈 베이컨 단계를 저장
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for i := 1; i < len(graph[front[0]]); i++ {
			if graph[front[0]][i] == 1 && !visited[i] {
				visited[i] = true
				cnt = front[1] + 1
				queue = append(queue, []int{i, cnt})
				if i == target {
					return cnt
				}
			}
		}
	}
	return cnt
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
