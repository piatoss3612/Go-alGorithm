package bj2491

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
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = scanInt()
	}
	incDp := make([]int, n)
	incMax := 1
	decDp := make([]int, n)
	decMax := 1
	incDp[0], decDp[0] = 1, 1
	for i := 1; i < n; i++ {
		if seq[i-1] <= seq[i] {
			incDp[i] = incDp[i-1] + 1
			if incDp[i] > incMax {
				incMax = incDp[i]
			}
		} else {
			incDp[i] = 1
		}

		if seq[i-1] >= seq[i] {
			decDp[i] = decDp[i-1] + 1
			if decDp[i] > decMax {
				decMax = decDp[i]
			}
		} else {
			decDp[i] = 1
		}
	}
	if decMax > incMax {
		fmt.Fprintln(writer, decMax)
	} else {
		fmt.Fprintln(writer, incMax)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
