package bj1453

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
	seats := make(map[string]int)
	cnt := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		seat := scanner.Text()
		_, ok := seats[seat]
		if ok {
			cnt++
		} else {
			seats[seat] = 1
		}
	}
	fmt.Fprintln(writer, cnt)
}
