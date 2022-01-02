package bj18310

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
	sort.Ints(input)
	if n%2 == 0 {
		// 짝수인 경우의 최적해는 2개의 중간값 중 작은 값
		fmt.Fprintln(writer, input[(n/2)-1])
	} else {
		// 홀수인 경우의 최적해는 중간값
		fmt.Fprintln(writer, input[n/2])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
