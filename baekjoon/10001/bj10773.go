package bj10773

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
	n := nextInt()

	var nums []int
	for i := 0; i < n; i++ {
		m := nextInt()
		if m == 0 {
			nums = nums[:len(nums)-1]
		} else {
			nums = append(nums, m)
		}
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Fprintln(writer, sum)
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
