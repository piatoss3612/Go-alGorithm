package bj3052

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

	remainders := make(map[int]int)

	for i := 0; i < 10; i++ {
		scanner.Scan()
		s := scanner.Text()
		n, _ := strconv.Atoi(s)
		n %= 42
		_, ok := remainders[n]
		if ok {
			remainders[n]++
		} else {
			remainders[n] = 1
		}
	}
	fmt.Fprintf(writer, "%d\n", len(remainders))
}
