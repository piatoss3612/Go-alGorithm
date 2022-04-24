package bj2667

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
	graph   [][]byte
	visited [][]bool
	counts  []int
	n       int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	graph = make([][]byte, n+1)
	visited = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]byte, n+1)
		visited[i] = make([]bool, n+1)
		if i != 0 {
			scanner.Scan()
			graph[i] = append([]byte{0}, scanner.Bytes()...)
		}
	}
	complex := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if graph[i][j] == 49 && !visited[i][j] {
				counts = append(counts, DFS(i, j, 0))
				complex += 1
			}
		}
	}
	sort.Ints(counts)
	fmt.Fprintln(writer, complex)
	for _, v := range counts {
		fmt.Fprintln(writer, v)
	}
}

func DFS(i, j, cnt int) int {
	visited[i][j] = true
	cnt += 1

	if validCoord(i-1, j) {
		cnt += DFS(i-1, j, 0)
	}

	if validCoord(i, j-1) {
		cnt += DFS(i, j-1, 0)
	}

	if validCoord(i, j+1) {
		cnt += DFS(i, j+1, 0)
	}

	if validCoord(i+1, j) {
		cnt += DFS(i+1, j, 0)
	}
	return cnt
}

func validX(x int) bool {
	if x >= 1 && x <= n {
		return true
	}
	return false
}

func validY(y int) bool {
	if y >= 1 && y <= n {
		return true
	}
	return false
}

func validCoord(x, y int) bool {
	if validX(x) && validY(y) && (graph[x][y] == 49) && !visited[x][y] {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
