package bj2178

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]byte
	visited [][]bool
	n, m    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	graph = make([][]byte, n+1)
	visited = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]byte, m+1)
		visited[i] = make([]bool, m+1)
	}
	for i := 1; i <= n; i++ {
		scanner.Scan()
		graph[i] = append([]byte{0}, (scanner.Bytes())...)
	}
	BFS(1, 1)
}

func BFS(i, j int) {
	visited[i][j] = true
	queue := [][]int{{i, j, 1}}

	for len(queue) > 0 {
		front := queue[0]
		r, c, cnt := front[0], front[1], front[2]
		if r == n && c == m {
			fmt.Fprintln(writer, cnt)
			return
		}
		queue = queue[1:]

		if validCoord(r-1, c) {
			visited[r-1][c] = true
			queue = append(queue, []int{r - 1, c, cnt + 1})
		}

		if validCoord(r, c-1) {
			visited[r][c-1] = true
			queue = append(queue, []int{r, c - 1, cnt + 1})
		}

		if validCoord(r, c+1) {
			visited[r][c+1] = true
			queue = append(queue, []int{r, c + 1, cnt + 1})
		}

		if validCoord(r+1, c) {
			visited[r+1][c] = true
			queue = append(queue, []int{r + 1, c, cnt + 1})
		}
	}
}

func validCoord(r, c int) bool {
	if validRow(r) && validCol(c) && !visited[r][c] && (graph[r][c] == 49) {
		return true
	}
	return false
}

func validRow(v int) bool {
	if v >= 1 && v <= n {
		return true
	}
	return false
}

func validCol(v int) bool {
	if v >= 1 && v <= m {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
