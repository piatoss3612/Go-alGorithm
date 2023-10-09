package bj1325

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	visited []bool // 방문 여부 확인
	counts  []int  // 각 정점에서 해킹할 수 있는 컴퓨터의 최댓값을 저장
	cnt     int
	n, m    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	graph = make([][]int, n+1)
	counts = make([]int, n+1)

	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		graph[b] = append(graph[b], a)
	}

	max := 0

	for i := 1; i <= n; i++ {
		visited = make([]bool, n+1)
		DFS(i)
		counts[i] = cnt
		if cnt > max {
			max = cnt
		}
		cnt = 0
	}

	for i := 1; i <= n; i++ {
		if counts[i] == max {
			fmt.Fprintf(writer, "%d ", i)
		}
	}
	fmt.Fprintln(writer)
}

func DFS(i int) {
	visited[i] = true

	for _, v := range graph[i] {
		if !visited[v] {
			DFS(v)
			cnt += 1
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
