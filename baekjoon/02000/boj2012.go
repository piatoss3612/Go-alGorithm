package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	// 슬라이스를 오름차순으로 정렬하면 최적의 해를 구할 수 있다
	sort.Ints(input)

	result := 0
	for i := 0; i < n; i++ {
		tmp := input[i] - (i + 1)
		if tmp < 0 {
			tmp *= -1
		}
		result += tmp
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
