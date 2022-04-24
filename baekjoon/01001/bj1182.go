package bj1182

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	cnt     = 0
	n, s    int
	input   []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, s = scanInt(), scanInt()
	input = make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	getSubSum(0, 0)
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getSubSum(idx, sum int) {
	if idx >= n {
		return
	}

	sum += input[idx]

	if sum == s {
		cnt += 1
	}
	getSubSum(idx+1, sum-input[idx])
	getSubSum(idx+1, sum)
}
