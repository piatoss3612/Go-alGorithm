package bj2468

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
	visited [][]bool
	n       int
	end     = 0
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	graph = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		if i != 0 {
			for j := 1; j <= n; j++ {
				graph[i][j] = scanInt()
				if graph[i][j] > end {
					end = graph[i][j]
				}
			}
		}
	}
	ans := 0
	for k := 0; k < end; k++ {
		initVisited()
		cnt := 0
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if graph[i][j] > k && !visited[i][j] {
					DFS(i, j, k)
					cnt += 1
				}
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	fmt.Fprintln(writer, ans)
}

func DFS(x, y, limit int) {
	visited[x][y] = true

	if validCoord(x+1, y) && graph[x][y] > limit {
		DFS(x+1, y, limit)
	}

	if validCoord(x-1, y) && graph[x][y] > limit {
		DFS(x-1, y, limit)
	}

	if validCoord(x, y+1) && graph[x][y] > limit {
		DFS(x, y+1, limit)
	}

	if validCoord(x, y-1) && graph[x][y] > limit {
		DFS(x, y-1, limit)
	}
}

func validCoord(x, y int) bool {
	return validX(x) && validY(y) && !visited[x][y]
}

func validX(x int) bool {
	return x >= 1 && x <= n
}

func validY(y int) bool {
	return y >= 1 && y <= n
}

func initVisited() {
	visited = make([][]bool, n+1)
	for i := 1; i <= n; i++ {
		visited[i] = make([]bool, n+1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
