package bj2751

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
	n := nextInt()

	nums := []int{}

	for i := 1; i <= n; i++ {
		m := nextInt()
		nums = append(nums, m)
	}
	sort.Ints(nums)
	printNums(nums)
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func printNums(n []int) {
	for _, v := range n {
		fmt.Fprintf(writer, "%d\n", v)
	}
}
