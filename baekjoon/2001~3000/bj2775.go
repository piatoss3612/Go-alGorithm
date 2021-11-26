package bj2775

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
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		apart := make([][]int, 0, k)

		for i := 0; i < k; i++ {
			line := make([]int, 0, n)
			if i == 0 {
				for k := 0; k < n; k++ {
					if k == 0 {
						line = append(line, 1)
					} else {
						line = append(line, line[k-1]+k+1)
					}
				}
				apart = append(apart, line)
				continue
			}
			for j := 0; j < n; j++ {
				if j == 0 {
					line = append(line, 1)
				} else {
					line = append(line, sum(apart[i-1], j))
				}
			}
			apart = append(apart, line)
		}
		fmt.Fprintln(writer, apart[k-1][n-1])
	}
}

func sum(line []int, j int) int {
	sum := 0
	for i, v := range line {
		if i > j {
			break
		}
		sum += v
	}
	return sum
}
