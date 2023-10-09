package bj2798

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
	m, _ := strconv.Atoi(scanner.Text())

	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	result := blackJack(nums, n, m)
	fmt.Fprintln(writer, result)
}

func blackJack(nums []int, n, m int) int {
	var result int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum == m {
					return sum
				}
				if result < sum && sum < m {
					result = sum
				}
			}
		}
	}
	return result
}
