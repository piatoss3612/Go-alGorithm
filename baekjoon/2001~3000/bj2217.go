package bj2217

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

	inputs := make([]int, n)
	for i := 0; i < n; i++ {
		inputs[i] = scanInt()
	}
	sort.Ints(inputs)

	cnt := len(inputs)
	max := 0
	for i := 0; i < n; i++ {
		w := inputs[i] * cnt
		if w > max {
			max = w
		}
		cnt -= 1
	}
	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
