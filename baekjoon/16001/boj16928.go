package bj16928

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
	board   [101]int
	visited [101]bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()

	for i := 0; i < n; i++ {
		x, y := scanInt(), scanInt()
		board[x] = y
	}

	for i := 0; i < m; i++ {
		u, v := scanInt(), scanInt()
		board[u] = v
	}

	Search()
}

func Search() {
	visited[1] = true
	queue := [][]int{{1, 0}}
	ans := math.MaxInt64

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front[0] == 100 {
			ans = getMin(ans, front[1])
			continue
		}

		for i := 1; i <= 6; i++ {
			idx := front[0] + i
			if idx <= 100 && !visited[idx] {
				visited[idx] = true
				if board[idx] != 0 {
					queue = append(queue, []int{board[idx], front[1] + 1})
				} else {
					queue = append(queue, []int{idx, front[1] + 1})
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func getMin(a, b int) int {
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
