package bj1920

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
	s1 := make(map[string]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s1[scanner.Text()] = 1
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for j := 0; j < m; j++ {
		scanner.Scan()
		if s1[scanner.Text()] == 1 {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
