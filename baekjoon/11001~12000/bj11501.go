package main

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
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		input := make([]int, n)
		for j := 0; j < n; j++ {
			input[j] = scanInt()
		}
		result := 0
		max := input[n-1]
		// 최대 수익을 얻기 위해서 뒤에서부터 최댓값을 구하고
		// 최댓값보다 큰 수가 나오면 바꿔주고
		// 최댓값보다 작은 수가 나오면 빼서 수익에 더해주면
		// 최댓값, 즉 최적해를 찾을 수 있다
		for k := n - 2; k >= 0; k-- {
			if input[k] > max {
				max = input[k]
				continue
			}
			result += max - input[k]
		}
		fmt.Fprintln(writer, result)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
