package bj1654

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	"os"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	k, n := scanInt(), scanInt()

	lans := make([]int, 0, k)
	sum := 0
	for i := 0; i < k; i++ {
		lan := scanInt()
		sum += lan
		lans = append(lans, lan)
	}

	start := 1
	end := sum / n

	for start <= end {
		mid := (start + end) / 2
		tmp := 0
		for _, v := range lans {
			tmp += v / mid
		}
		if tmp >= n {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	fmt.Fprintln(writer, start-1)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
