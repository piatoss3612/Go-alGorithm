package bj1205

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

	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	p, _ := strconv.Atoi(scanner.Text())

	ranks := make([]int, 0, p)
	if n > 0 {
		for i := 0; i < n; i++ {
			scanner.Scan()
			ranked, _ := strconv.Atoi(scanner.Text())
			ranks = append(ranks, ranked)
		}
	} else {
		fmt.Fprintln(writer, 1)
		return
	}

	if (n == p) && ranks[n-1] >= s {
		fmt.Fprintln(writer, -1)
		return
	}

	rank := 1
	for _, v := range ranks {
		if v > s {
			rank++
		}
	}
	fmt.Fprintln(writer, rank)
}
