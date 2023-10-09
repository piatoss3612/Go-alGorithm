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
	inDegree []int   // 진입 차수
	graph    [][]int // 노드들의 연결 여부 저장
)

// 메모리: 8812KB
// 시간: 52ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	inDegree = make([]int, n+1)
	graph = make([][]int, n+1)

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		graph[b] = append(graph[b], a) // a <- b인 비순환 방향 그래프
		inDegree[a] += 1               // b가 a로 진입하므로 a의 진입 차수 증가
	}

	topologicalSort() // 위상 정렬
}

func topologicalSort() {
	var res []int // 진입 차수가 0, 키가 가장 작은 학생부터 차례대로 저장
	var q []int

	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i) // 진입 차수가 0인 학생의 번호를 큐에 추가
		}
	}

	for len(q) > 0 {
		front := q[0] // 진입 차수가 0인 학생을 한 명씩 큐에서 꺼내오기
		q = q[1:]

		res = append(res, front)

		// 큐에서 꺼낸 학생과 연결된 학생들의 진입 차수를 1씩 줄이고
		// 진입 차수가 줄어든 학생의 진입 차수가 0이 되면 큐에 추가
		for i := 0; i < len(graph[front]); i++ {
			temp := graph[front][i]
			inDegree[temp] -= 1
			if inDegree[temp] == 0 {
				q = append(q, temp)
			}
		}
	}

	// 키가 큰 학생부터 출력
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", res[i])
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
