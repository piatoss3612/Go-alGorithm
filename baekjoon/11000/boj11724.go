package main

import (
	"bufio"
	_ "bytes"
	_ "container/heap"
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
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
	}
	visited = make([]bool, n+1)
	for i := 0; i < m; i++ {
		u, v := scanInt(), scanInt()
		graph[u][v] = 1
		graph[v][u] = 1
	}
	cnt := 0 // 깊이 우선 탐색을 실행할 때마다 하나의 연결된 요소를 찾게 되므로 이를 카운트 해준다
	for i := 1; i <= n; i++ {
		if visited[i] == false {
			cnt += 1
			DFS(i)
		}
	}
	fmt.Fprintln(writer, cnt)
}

func DFS(v int) {
	if visited[v] == true {
		return
	}
	visited[v] = true
	for idx, check := range graph[v] {
		if check == 1 {
			DFS(idx)
		}
	}
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
