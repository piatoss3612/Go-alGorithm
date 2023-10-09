package bj14921

import (
	"bufio"
	"fmt"
	"math"
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

	min := 200000001
	s, e := 0, n-1
	for s < e {
		tmp := input[s] + input[e]
		if tmp < 0 {
			s += 1
			min = getAbsMin(tmp, min)
		} else {
			e -= 1
			min = getAbsMin(tmp, min)
		}
	}
	fmt.Fprintln(writer, min)
}

func getAbsMin(a, b int) int {
	if math.Abs(float64(a)) < math.Abs(float64(b)) {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
