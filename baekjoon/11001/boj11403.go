package bj11403

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = scanInt()
			if graph[i][j] == 0 {
				// 간선이 존재하지 않을 경우,
				// 최소 경로를 비교할 때 사용하기 위해 32비트 정수의 최댓값을 할당
				// 32비트 정수 최댓값 + 32비트 정수 최댓값은 오버플로가 발생하지 않는다!
				graph[i][j] = math.MaxInt32
			}
		}
	}
	// 플로이드 와샬 알고리즘
	// h는 거쳐가는 정점
	// i는 시작점 j는 종점
	for h := 0; h < n; h++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				graph[i][j] = getMin(graph[i][h]+graph[h][j], graph[i][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] < math.MaxInt32 {
				fmt.Fprint(writer, 1, " ")
			} else {
				fmt.Fprint(writer, 0, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

func getMin(a, b int) int {
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
