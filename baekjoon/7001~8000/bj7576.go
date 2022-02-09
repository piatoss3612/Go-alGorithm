package bj7576

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	box     [][]int
	visited [][]bool
	cntZero int // 박스에 0이 남았는지 판별하기 위한 변수
	m, n    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	m, n = scanInt(), scanInt()
	box = make([][]int, n+1)
	visited = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		box[i] = make([]int, m+1)
		visited[i] = make([]bool, m+1)
		if i != 0 {
			for j := 1; j <= m; j++ {
				box[i][j] = scanInt()
				if box[i][j] == 0 {
					cntZero += 1
				}
			}
		}
	}

	// 1이 들어있는 칸들을 모아서 동시에 깊이 우선 탐색을
	// 진행해야 하므로 queue에 담아서 넘겨준다
	queue := [][]int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if box[i][j] == 1 {
				queue = append(queue, []int{i, j, 0})
				visited[i][j] = true
			}
		}
	}
	BFS(queue)
}

func BFS(queue [][]int) {
	ans := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y, cnt := front[0], front[1], front[2]
		ans = cnt

		if validCoord(x, y-1) {
			cntZero -= 1
			visited[x][y-1] = true
			queue = append(queue, []int{x, y - 1, cnt + 1})
		}

		if validCoord(x, y+1) {
			cntZero -= 1
			visited[x][y+1] = true
			queue = append(queue, []int{x, y + 1, cnt + 1})
		}

		if validCoord(x-1, y) {
			cntZero -= 1
			visited[x-1][y] = true
			queue = append(queue, []int{x - 1, y, cnt + 1})
		}

		if validCoord(x+1, y) {
			cntZero -= 1
			visited[x+1][y] = true
			queue = append(queue, []int{x + 1, y, cnt + 1})
		}
	}

	if cntZero != 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func validCoord(x, y int) bool {
	if validX(x) && validY(y) && box[x][y] == 0 && !visited[x][y] {
		return true
	}
	return false
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
