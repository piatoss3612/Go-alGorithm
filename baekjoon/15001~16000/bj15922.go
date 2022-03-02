package bj15922

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

type line struct {
	start int
	end   int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	lines := make([]line, n)
	for i := 0; i < n; i++ {
		lines[i] = line{scanInt(), scanInt()}
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].start < lines[j].start
	})

	ans := 0
	left := -1000000000
	right := -1000000000
	for i := 0; i < n; i++ {
		if right < lines[i].start {
			ans += right - left
			left = lines[i].start
			right = lines[i].end
		} else {
			right = getMax(right, lines[i].end)
		}
	}

	ans += right - left
	fmt.Fprintln(writer, ans)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
