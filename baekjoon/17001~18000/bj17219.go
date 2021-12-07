package bj17219

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

	n, m := scanInt(), scanInt()
	pwds := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		url := scanner.Text()
		scanner.Scan()
		pwd := scanner.Text()
		pwds[url] = pwd
	}

	for j := 0; j < m; j++ {
		scanner.Scan()
		fmt.Fprintln(writer, pwds[scanner.Text()])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
