package bj10989

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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	slice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		slice = append(slice, m)
	}
	sort.Ints(slice)
	for _, v := range slice {
		fmt.Fprintln(writer, v)
	}
}
