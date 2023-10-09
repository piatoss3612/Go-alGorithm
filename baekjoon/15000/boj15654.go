package bj15654

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
	n, m    int
	input   []int
	seq     []int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	input = make([]int, n)
	visited = make([]bool, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	BackTracking(0)
}

func BackTracking(idx int) {
	if idx == m {
		printSeq()
		return
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		seq = append(seq, input[i])
		BackTracking(idx + 1)
		seq = seq[:len(seq)-1]
		visited[i] = false
	}
}

func printSeq() {
	for _, v := range seq {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
