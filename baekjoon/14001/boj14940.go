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
	n, m    int
	graph   [][]int
	visited [][]bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	graph = make([][]int, n+1)
	visited = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, m+1)
		visited[i] = make([]bool, m+1)
	}

	var x, y int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			graph[i][j] = scanInt()
			if graph[i][j] == 2 {
				x, y = i, j
			}
		}
	}

	BFS(x, y, 0)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				fmt.Fprintf(writer, "%d ", -1)
			} else {
				fmt.Fprintf(writer, "%d ", graph[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
}

func BFS(x, y, cnt int) {
	visited[x][y] = true

	queue := [][]int{{x, y, cnt}}

	var curX, curY, curCnt int
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		curX, curY, curCnt = front[0], front[1], front[2]
		graph[curX][curY] = curCnt

		if validCoord(curX, curY-1) {
			visited[curX][curY-1] = true
			queue = append(queue, []int{curX, curY - 1, curCnt + 1})
		}

		if validCoord(curX, curY+1) {
			visited[curX][curY+1] = true
			queue = append(queue, []int{curX, curY + 1, curCnt + 1})
		}

		if validCoord(curX-1, curY) {
			visited[curX-1][curY] = true
			queue = append(queue, []int{curX - 1, curY, curCnt + 1})
		}

		if validCoord(curX+1, curY) {
			visited[curX+1][curY] = true
			queue = append(queue, []int{curX + 1, curY, curCnt + 1})
		}
	}
}

func validX(x int) bool {
	if x >= 1 && x <= n {
		return true
	}
	return false
}

func validY(y int) bool {
	if y >= 1 && y <= m {
		return true
	}
	return false
}

func validCoord(x, y int) bool {
	if validX(x) && validY(y) && !visited[x][y] && graph[x][y] != 0 {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
