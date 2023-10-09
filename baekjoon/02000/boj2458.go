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
	INF     = 100000
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i != j {
				graph[i][j] = INF
			}
		}
	}

	for i := 1; i <= m; i++ {
		graph[scanInt()][scanInt()] = 1
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	ans := 0

	/*
		자신의 키가 몇 등인지 알기 위해서는
		나한테 날라오는 다른 학생들의 화살표 + 나에게서 다른 학생으로 뻗어나가는 화살표의 수가
		전체 학생의 수 - 나, 즉 n -1이 되어야 한다

		이는 플로이드 와샬 알고리즘을 통해 a에서 b로 가는 경로, b에서 a로 가는 경로가 있는 경우
		즉, graph[a][b] != INF && graph[b][a] != INF를 모두 만족하는 경우의 학생을 결괏값에 추가해준다
	*/

	for i := 1; i <= n; i++ {
		flag := true
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			if graph[i][j] == INF && graph[j][i] == INF {
				flag = false
				break
			}
		}
		if flag {
			ans += 1
		}
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
