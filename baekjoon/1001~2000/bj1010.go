package bj1010

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
	t := scanInt()

	mx := [31][31]int{}

	for i := 1; i <= 30; i++ {
		for j := i; j <= 30; j++ {
			if i == 1 {
				mx[i][j] = j
			} else {
				mx[i][j] = sum(mx[i-1][:j])
			}
		}
	}

	for k := 0; k < t; k++ {
		a, b := scanInt(), scanInt()
		fmt.Fprintln(writer, mx[a][b])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
