package bj1343

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s := scanner.Text()
	if strings.Contains(s, "XXXX") {
		s = strings.ReplaceAll(s, "XXXX", "AAAA")
	}

	if strings.Contains(s, "XX") {
		s = strings.ReplaceAll(s, "XX", "BB")
	}

	if strings.Contains(s, "X") {
		fmt.Fprintf(writer, "%d\n", -1)
		return
	}
	fmt.Fprintln(writer, s)
}
