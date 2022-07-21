package bj2606

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	c := scanInt()
	n := scanInt()
	graph := make([][]int, c+1)
	for i := 0; i <= c; i++ {
		row := make([]int, c+1)
		graph[i] = row
	}
	for i := 0; i < n; i++ {
		a, b := scanInt(), scanInt()
		graph[a][b] = 1
		graph[b][a] = 1
	}

	infected := make([]bool, c+1)
	infected[1] = true

	queue := []int{}
	queue = append(queue, 1)

	result := []int{}

	for len(queue) != 0 {
		check := queue[0]
		if len(queue) == 1 {
			queue = []int{}
		} else {
			queue = queue[1:]
		}

		if infected[check] == true {
			for i, v := range graph[check] {
				if v == 1 && infected[i] == false {
					infected[i] = true
					queue = append(queue, i)
					result = append(result, i)
				}
			}
		}
	}
	fmt.Fprintln(writer, len(result))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
