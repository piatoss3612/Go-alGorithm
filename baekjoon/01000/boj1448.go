package bj1448

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

	max := 0
	for i := len(inputs) - 1; i >= 2; i-- {
		if inputs[i] < inputs[i-1]+inputs[i-2] {
			max = inputs[i] + inputs[i-1] + inputs[i-2]
			break
		}
	}

	if max == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
