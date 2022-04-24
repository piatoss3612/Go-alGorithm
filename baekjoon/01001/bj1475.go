package bj1475

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
	n := strings.Split(scanner.Text(), "")

	digits := [10]int{}
	for _, v := range n {
		idx, _ := strconv.Atoi(v)
		if idx == 9 || idx == 6 {
			if digits[6] <= digits[9] {
				digits[6]++
			} else {
				digits[9]++
			}
		} else {
			digits[idx]++
		}

	}
	result := getMax(digits)
	fmt.Fprintln(writer, result)
}

func getMax(n [10]int) int {
	max := 0
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}
