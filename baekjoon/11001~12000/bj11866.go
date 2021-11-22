package bj11866

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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	seq := []int{}

	for i := k - 1; len(seq) != n; i += k {
		if i >= len(nums) {
			i %= len(nums)
		}
		seq = append(seq, nums[i])
		nums = append(nums[:i], nums[i+1:]...)
		i--
	}
	fmt.Fprint(writer, "<")
	for i, v := range seq {
		fmt.Fprint(writer, v)
		if i != len(seq)-1 {
			fmt.Fprint(writer, ", ")
		}
	}
	fmt.Fprint(writer, ">\n")
}
