package bj2576

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	scanner.Split(bufio.ScanWords)

	min := 100
	sum := 0
	for i := 0; i < 7; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		check := checkOdd(n)
		if check {
			sum += n
			min = getMin(min, n)
		}
	}

	if sum == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Printf("%d\n%d\n", sum, min)

}

func checkOdd(n int) bool {
	if (n % 2) == 1 {
		return true
	}
	return false
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
