package bj1439

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
	digits := scanDigits()
	abb := sameDigits(digits)
	zero := strings.Count(abb, "0")
	one := strings.Count(abb, "1")
	if zero < one {
		fmt.Fprintf(writer, "%d\n", zero)
	} else {
		fmt.Fprintf(writer, "%d\n", one)
	}
}

func scanDigits() []string {
	scanner.Scan()
	s := strings.Split(scanner.Text(), "")
	return s
}

func sameDigits(digits []string) string {
	var builder strings.Builder
	same := digits[0]
	for _, v := range digits {
		if v != same {
			builder.WriteString(same)
			same = v
		}
	}
	builder.WriteString(same)
	return builder.String()
}
