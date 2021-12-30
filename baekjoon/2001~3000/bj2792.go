package bj2792

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	left, right := 1, 0
	input := make([]int, m)
	for i := 0; i < m; i++ {
		input[i] = scanInt()
		if input[i] > right {
			right = input[i]
		}
	}
	for left <= right {
		mid := (left + right) / 2
		cnt := 0
		for i := 0; i < m; i++ {
			if input[i]%mid == 0 {
				cnt += input[i] / mid
			} else {
				cnt += input[i]/mid + 1
			}
		}
		if cnt > n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Fprintln(writer, left)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
