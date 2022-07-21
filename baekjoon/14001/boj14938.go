package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [101][101]int
	INF     int = 10000000
)

// 그래프 초기화
func init() {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if i == j {
				continue
			}
			graph[i][j] = INF
		}
	}
}

// 메모리: 988KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, r := scanInt(), scanInt(), scanInt()

	items := make([]int, n+1) // 지역마다 아이템을 얻을 수 있는 아이템의 개수를 저장
	for i := 1; i <= n; i++ {
		items[i] = scanInt()
	}

	var a, b, l int
	for i := 1; i <= r; i++ {
		a, b, l = scanInt(), scanInt(), scanInt()
		// 그래프는 양방향 그래프!
		graph[a][b] = l
		graph[b][a] = l
	}

	// 플로이드 와샬
	// i->j로 최단거리를 갱신
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	total := make([]int, n+1) // 어떤 지역에서 최대 얻을 수 있는 아이템의 수를 저장

	// 플로이드 와샬을 수행해 계산한 최단거리를 사용해
	// i에서 시작하여 특정 지역까지의 거리가 m보다 작거나 같은 지역을 구하고
	// 해당 지역에서 얻을 수 있는 아이템의 수를 total에 누적해서 더해준다
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if graph[i][j] <= m {
				total[i] += items[j]
			}
		}
	}

	sort.Ints(total)               // 오름차순으로 정렬하여
	fmt.Fprintln(writer, total[n]) // 가장 아이템을 많이 얻을 수 있는 경우를 출력한다
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
