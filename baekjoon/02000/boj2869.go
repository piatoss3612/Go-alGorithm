package bj2869

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
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text())

	cnt := (v - b) / (a - b)

	if (v-b)%(a-b) != 0 {
		cnt++
	}

	fmt.Fprintln(writer, cnt)
}
