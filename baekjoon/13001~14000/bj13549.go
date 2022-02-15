package bj13549

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	minTime int = 100001
	visited [100001]bool
	n, k    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k = scanInt(), scanInt()

	BFS()
	fmt.Fprintf(writer, "%d\n", minTime)
}

func BFS() {
	visited[n] = true
	queue := [][]int{{n, 0}}

	for len(queue) > 0 {
		idx := queue[0][0]
		time := queue[0][1]

		visited[idx] = true
		queue = queue[1:]

		if idx == k {
			if time < minTime {
				minTime = time
			}
		}

		if valid(idx-1) && !visited[idx-1] {
			queue = append(queue, []int{idx - 1, time + 1})
		}

		if valid(idx+1) && !visited[idx+1] {
			queue = append(queue, []int{idx + 1, time + 1})
		}

		if valid(idx*2) && !visited[idx*2] {
			queue = append(queue, []int{idx * 2, time})
		}
	}
}

func valid(v int) bool {
	return v >= 0 && v <= 100000
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
