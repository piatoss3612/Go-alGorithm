package bj1449

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
	n, l := scanInt(), scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	// 오름차 순으로 정렬
	sort.Ints(input)
	if n == 1 {
		fmt.Fprintln(writer, 1)
		return
	}
	cnt := 1
	start := input[0]
	for i := 1; i < n; i++ {
		if input[i]-start > l-1 {
			cnt += 1
			start = input[i]
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
