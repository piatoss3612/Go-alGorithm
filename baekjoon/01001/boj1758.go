package main

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
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i] > inputs[j]
	})
	result := 0
	for i := 0; i < len(inputs); i++ {
		tip := inputs[i] - i
		if tip >= 0 {
			result += tip
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
