package bj16435

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
	reader  = bufio.NewReader(os.Stdin)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	l := scanInt()

	inputs := make([]int, n)

	for i := 0; i < n; i++ {
		inputs[i] = scanInt()
	}

	sort.Ints(inputs)

	for _, v := range inputs {
		if v <= l {
			l += 1
		} else {
			break
		}
	}
	fmt.Fprintln(writer, l)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
