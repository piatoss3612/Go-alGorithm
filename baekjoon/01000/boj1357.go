package bj1357

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintf(writer, "%d\n", Rev(Rev(x)+Rev(y)))
}

func Rev(n int) int {
	s := strconv.Itoa(n)
	letters := strings.Split(s, "")
	var builder strings.Builder
	for i := len(letters) - 1; i >= 0; i-- {
		builder.WriteString(letters[i])
	}
	result, _ := strconv.Atoi(builder.String())
	return result
}
