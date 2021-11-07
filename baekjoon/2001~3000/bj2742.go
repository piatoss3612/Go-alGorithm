package bj2742

import (
	"bufio"
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
	for i := n; i > 0; i-- {
		// fmt.Fprintf(writer, "%d\n", i)
		writer.WriteString(strconv.Itoa(i))
		writer.WriteByte('\n')
	}
}
