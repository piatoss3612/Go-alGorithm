package bj12852

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   []int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	graph = make([]int, n+1)
	visited = make([]bool, n+1)
	BFS(n)
}

func BFS(v int) {
	visited[v] = true
	queue := [][]int{{v, 0, v}}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front[0] == 1 {
			fmt.Fprintln(writer, front[1])
			for _, v := range front[2:] {
				fmt.Fprintf(writer, "%d ", v)
			}
			fmt.Fprintln(writer)
			return
		}
		if front[0]-1 >= 1 && !visited[front[0]-1] {
			visited[front[0]-1] = true
			tmp := []int{front[0] - 1, front[1] + 1}
			tmp = append(tmp, front[2:]...)
			tmp = append(tmp, front[0]-1)
			queue = append(queue, tmp)
		}

		if front[0]%3 == 0 && !visited[front[0]/3] {
			visited[front[0]/3] = true
			tmp := []int{front[0] / 3, front[1] + 1}
			tmp = append(tmp, front[2:]...)
			tmp = append(tmp, front[0]/3)
			queue = append(queue, tmp)
		}

		if front[0]%2 == 0 && !visited[front[0]/2] {
			visited[front[0]/2] = true
			tmp := []int{front[0] / 2, front[1] + 1}
			tmp = append(tmp, front[2:]...)
			tmp = append(tmp, front[0]/2)
			queue = append(queue, tmp)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
