package bj2292

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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	sum := 1
	cnt := 1

	if n == 1 {
		fmt.Fprintln(writer, cnt)
		return
	}

	for i := 1; ; i++ {
		sum += i * 6
		cnt++
		if sum >= n {
			fmt.Fprintln(writer, cnt)
			break
		}
	}
}
