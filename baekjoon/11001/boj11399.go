package bj11399

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
	min := 0
	stack := 0
	for _, v := range inputs {
		stack += v
		min += stack
	}
	fmt.Fprintln(writer, min)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
