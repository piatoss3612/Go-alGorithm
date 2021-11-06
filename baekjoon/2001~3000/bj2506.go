package bj2506

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	scores := strings.Split(scanner.Text(), " ")

	result := 0
	score := 0

	for i := 0; i < n; i++ {
		if scores[i] == "1" {
			score += 1
			result += score
		} else {
			score = 0
		}
	}

	fmt.Println(result)
}
