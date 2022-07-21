package bj10816

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
	cards := make(map[string]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num := scanner.Text()
		_, ok := cards[num]
		if ok {
			cards[num]++
		} else {
			cards[num] = 1
		}
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for j := 0; j < m; j++ {
		scanner.Scan()
		mum := scanner.Text()
		_, ok := cards[mum]
		if ok {
			fmt.Fprint(writer, cards[mum], " ")
		} else {
			fmt.Fprint(writer, 0, " ")
		}
	}
	fmt.Fprintln(writer)
}
