package main

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
	for i := 0; i < t; i++ {
		n := scanInt()
		input := make([]int, n)
		for j := 0; j < n; j++ {
			input[j] = scanInt()
		}
		var sum int64 = 0
		for j := 0; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum += int64(gcd(input[j], input[k]))
			}
		}
		fmt.Fprintln(writer, sum)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
