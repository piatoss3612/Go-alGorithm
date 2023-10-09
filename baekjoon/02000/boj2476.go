package bj2476

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

	scores := make([]int, n)

	for i := 0; i < n; i++ {
		dices := rollDices()
		dice1, _ := strconv.Atoi(dices[0])
		dice2, _ := strconv.Atoi(dices[1])
		dice3, _ := strconv.Atoi(dices[2])
		score := getScore(dice1, dice2, dice3)
		scores = append(scores, score)
	}
	var ms int = 0

	for _, v := range scores {
		ms = max(ms, v)
	}
	fmt.Println(ms)
}

func rollDices() []string {
	scanner.Scan()
	dices := strings.Split(scanner.Text(), " ")
	return dices
}

func getScore(d1, d2, d3 int) int {
	if d1 == d2 && d2 == d3 && d3 == d1 {
		return (d1 * 1000) + 10000
	} else if d1 != d2 && d2 != d3 && d3 != d1 {
		return max(max(d1, d2), d3) * 100
	}

	if d1 == d2 {
		return d1*100 + 1000
	} else if d2 == d3 {
		return d2*100 + 1000
	}
	return d3*100 + 1000
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
