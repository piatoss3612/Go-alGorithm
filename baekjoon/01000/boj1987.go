// map을 사용한 경우
// 메모리: 928KB
// 시간: 3044ms
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var (
// 	scanner = bufio.NewScanner(os.Stdin)
// 	writer  = bufio.NewWriter(os.Stdout)
// 	n, m    int
// 	graph   [][]rune
// 	visited [][]bool
// 	checked map[rune]bool
// 	max     = 0
// )

// func main() {
// 	defer writer.Flush()
// 	scanner.Split(bufio.ScanWords)
// 	n, m = scanInt(), scanInt()

// 	graph = make([][]rune, n)
// 	visited = make([][]bool, n)
// 	checked = make(map[rune]bool)

// 	for i := 0; i < n; i++ {
// 		graph[i] = make([]rune, m)
// 		visited[i] = make([]bool, m)
// 		graph[i] = []rune(scanString())
// 	}

// 	visited[0][0] = true
// 	checked[graph[0][0]] = true
// 	dfs(0, 0, 1)

// 	fmt.Fprintln(writer, max)
// }

// func dfs(x, y, cnt int) {
// 	if cnt > max {
// 		max = cnt
// 	}

// 	if validPoint(x-1, y) && !visited[x-1][y] {
// 		if !checked[graph[x-1][y]] {
// 			checked[graph[x-1][y]] = true
// 			visited[x-1][y] = true
// 			dfs(x-1, y, cnt+1)
// 			checked[graph[x-1][y]] = false
// 			visited[x-1][y] = false
// 		}
// 	}

// 	if validPoint(x+1, y) && !visited[x+1][y] {
// 		if !checked[graph[x+1][y]] {
// 			checked[graph[x+1][y]] = true
// 			visited[x+1][y] = true
// 			dfs(x+1, y, cnt+1)
// 			checked[graph[x+1][y]] = false
// 			visited[x+1][y] = false
// 		}
// 	}

// 	if validPoint(x, y-1) && !visited[x][y-1] {
// 		if !checked[graph[x][y-1]] {
// 			checked[graph[x][y-1]] = true
// 			visited[x][y-1] = true
// 			dfs(x, y-1, cnt+1)
// 			checked[graph[x][y-1]] = false
// 			visited[x][y-1] = false
// 		}
// 	}

// 	if validPoint(x, y+1) && !visited[x][y+1] {
// 		if !checked[graph[x][y+1]] {
// 			checked[graph[x][y+1]] = true
// 			visited[x][y+1] = true
// 			dfs(x, y+1, cnt+1)
// 			checked[graph[x][y+1]] = false
// 			visited[x][y+1] = false
// 		}
// 	}
// }

// func validX(x int) bool {
// 	return x >= 0 && x < n
// }

// func validY(y int) bool {
// 	return y >= 0 && y < m
// }

// func validPoint(x, y int) bool {
// 	return validX(x) && validY(y)
// }

// func scanInt() int {
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	return n
// }

// func scanString() string {
// 	scanner.Scan()
// 	return scanner.Text()
// }

// []bool 슬라이스를 사용한 경우
// 메모리: 920KB
// 시간: 568ms
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
	graph   [][]rune
	visited [][]bool
	checked []bool
	max     = 0
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	graph = make([][]rune, n)
	visited = make([][]bool, n)
	checked = make([]bool, 26)

	for i := 0; i < n; i++ {
		graph[i] = make([]rune, m)
		visited[i] = make([]bool, m)
		graph[i] = []rune(scanString())
	}

	visited[0][0] = true
	checked[graph[0][0]-'A'] = true
	dfs(0, 0, 1)

	fmt.Fprintln(writer, max)
}

func dfs(x, y, cnt int) {
	if cnt > max {
		max = cnt
	}

	if validPoint(x-1, y) && !visited[x-1][y] {
		if !checked[graph[x-1][y]-'A'] {
			checked[graph[x-1][y]-'A'] = true
			visited[x-1][y] = true
			dfs(x-1, y, cnt+1)
			checked[graph[x-1][y]-'A'] = false
			visited[x-1][y] = false
		}
	}

	if validPoint(x+1, y) && !visited[x+1][y] {
		if !checked[graph[x+1][y]-'A'] {
			checked[graph[x+1][y]-'A'] = true
			visited[x+1][y] = true
			dfs(x+1, y, cnt+1)
			checked[graph[x+1][y]-'A'] = false
			visited[x+1][y] = false
		}
	}

	if validPoint(x, y-1) && !visited[x][y-1] {
		if !checked[graph[x][y-1]-'A'] {
			checked[graph[x][y-1]-'A'] = true
			visited[x][y-1] = true
			dfs(x, y-1, cnt+1)
			checked[graph[x][y-1]-'A'] = false
			visited[x][y-1] = false
		}
	}

	if validPoint(x, y+1) && !visited[x][y+1] {
		if !checked[graph[x][y+1]-'A'] {
			checked[graph[x][y+1]-'A'] = true
			visited[x][y+1] = true
			dfs(x, y+1, cnt+1)
			checked[graph[x][y+1]-'A'] = false
			visited[x][y+1] = false
		}
	}
}

func validX(x int) bool {
	return x >= 0 && x < n
}

func validY(y int) bool {
	return y >= 0 && y < m
}

func validPoint(x, y int) bool {
	return validX(x) && validY(y)
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
