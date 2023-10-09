package bj14241

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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	result := 0
	for i := 1; i < n; i++ {
		result += input[i-1] * input[i]
		input[i] += input[i-1]
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
