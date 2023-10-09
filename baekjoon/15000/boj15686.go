package bj15686

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	houses   [][]int
	chickens [][]int
	picked   [][]int
	visited  []bool
	n, m     int
	ans      = 100000000
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	chickens = [][]int{}
	houses = [][]int{}

	n, m = scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			tmp := scanInt()
			if tmp == 1 {
				houses = append(houses, []int{i, j})
			} else if tmp == 2 {
				chickens = append(chickens, []int{i, j})
			}
		}
	}

	visited = make([]bool, len(chickens))

	solve(0, 0)
	fmt.Fprintln(writer, ans)
}

func solve(idx, cnt int) {
	if cnt == m {
		tmp := 0
		for i := 0; i < len(houses); i++ {
			min := 10000000
			for j := 0; j < len(picked); j++ {
				x := math.Abs(float64(picked[j][0] - houses[i][0]))
				y := math.Abs(float64(picked[j][1] - houses[i][1]))
				min = getMin(min, int(x+y))
			}
			tmp += min
		}
		ans = getMin(ans, tmp)
		return
	}

	for i := idx; i < len(chickens); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		picked = append(picked, chickens[i])
		solve(i, cnt+1)
		visited[i] = false
		picked = picked[:len(picked)-1]
	}
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
