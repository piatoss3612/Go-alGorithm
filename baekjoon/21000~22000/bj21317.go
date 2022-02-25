package bj21317

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	input   [][]int
	n, k    int
	ans     int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	input = make([][]int, n+1)
	for i := 1; i <= n-1; i++ {
		input[i] = []int{scanInt(), scanInt()}
	}

	k = scanInt()

	cross(0, 1, 1)

	fmt.Fprintln(writer, ans)
}

func cross(sum, idx, cnt int) {
	if idx == n {
		ans = getMin(ans, sum)
		return
	}

	if idx+1 <= n {
		cross(sum+input[idx][0], idx+1, cnt)
	}

	if idx+2 <= n {
		cross(sum+input[idx][1], idx+2, cnt)
	}

	if idx+3 <= n && cnt == 1 {
		cross(sum+k, idx+3, cnt-1)
	}
}

func getMin(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
