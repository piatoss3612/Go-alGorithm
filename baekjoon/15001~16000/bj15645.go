package bj15645

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
	dp1 := make([][]int, n+1)
	dp2 := make([][]int, n+1)
	dp1[0] = make([]int, 4)
	dp2[0] = make([]int, 4)
	for i := 1; i <= n; i++ {
		dp1[i] = make([]int, 4)
		dp2[i] = make([]int, 4)
		for j := 1; j <= 3; j++ {
			tmp := scanInt()
			dp1[i][j] = tmp
			dp2[i][j] = tmp
		}
	}

	for i := 2; i <= n; i++ {
		dp1[i][1] = getMax(dp1[i][1]+dp1[i-1][1], dp1[i][1]+dp1[i-1][2])
		dp2[i][1] = getMin(dp2[i][1]+dp2[i-1][1], dp2[i][1]+dp2[i-1][2])

		dp1[i][2] = getMax(getMax(dp1[i][2]+dp1[i-1][1], dp1[i][2]+dp1[i-1][2]), dp1[i][2]+dp1[i-1][3])
		dp2[i][2] = getMin(getMin(dp2[i][2]+dp2[i-1][1], dp2[i][2]+dp2[i-1][2]), dp2[i][2]+dp2[i-1][3])

		dp1[i][3] = getMax(dp1[i][3]+dp1[i-1][2], dp1[i][3]+dp1[i-1][3])
		dp2[i][3] = getMin(dp2[i][3]+dp2[i-1][2], dp2[i][3]+dp2[i-1][3])
	}

	sort.Ints(dp1[n])
	sort.Ints(dp2[n])
	fmt.Fprintln(writer, dp1[n][3], dp2[n][1])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
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
