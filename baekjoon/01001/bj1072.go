package bj1072

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
	x, y := scanInt(), scanInt()
	z := (y * 100) / x

	if z >= 99 {
		fmt.Fprintln(writer, -1)
		return
	}

	start := 0
	end := 1000000000
	for start <= end {
		mid := (start + end) / 2
		tmpX := x + mid
		tmpY := y + mid
		tmpZ := (tmpY * 100) / tmpX
		if tmpZ > z {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	fmt.Fprintln(writer, start)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
