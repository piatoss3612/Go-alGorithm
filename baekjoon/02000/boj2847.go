package bj2847

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
	n := scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	if n == 1 {
		fmt.Fprintln(writer, 0)
		return
	}
	cnt := 0
	for i := n - 2; i >= 0; i-- {
		if input[i] >= input[i+1] {
			// i번째 항이 i+1번째 항보다 큰 경우에 1작은 값으로 변환
			tmp := input[i] - input[i+1] + 1
			input[i] -= tmp
			cnt += tmp
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
