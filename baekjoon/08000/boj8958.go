package bj8958

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
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		getScore(scanner.Text())
	}
}

func getScore(s string) {
	ox := strings.Split(s, "")
	total := 0
	score := 1
	if ox[0] == "O" {
		total += score
	}
	for i := 1; i < len(ox); i++ {
		if ox[i] == "O" && ox[i] == ox[i-1] {
			score += 1
			total += score
		} else if ox[i] == "O" {
			total += score
		} else {
			score = 1
		}
	}
	fmt.Fprintf(writer, "%d\n", total)
}
