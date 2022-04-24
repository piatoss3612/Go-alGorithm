package bj3273

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
	// 오름차 순으로 정렬
	sort.Ints(input)
	x := scanInt()
	cnt := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if input[i]+input[j] > x {
				break
			}
			if input[i]+input[j] == x {
				cnt += 1
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
