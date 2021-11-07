package bj2721

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
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		var result int64
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		for j := 1; j <= n; j++ {
			result += int64(j) * triNum(j+1)
		}
		fmt.Fprintf(writer, "%d\n", result)
	}
}

func triNum(n int) int64 {
	k := int64(n)
	return k * (k + 1) / 2
}
