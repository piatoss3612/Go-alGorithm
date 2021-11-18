package bj1476

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
	e, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	if (e == s) && (s == m) {
		fmt.Fprintln(writer, e)
		return
	}

	for i := 15; ; i += 15 {
		year := i + e
		if ((year-s)%28) == 0 && ((year-m)%19) == 0 {
			fmt.Fprintln(writer, year)
			break
		}
	}
}
