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
	n, m := scanInt(), scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	// 오름차순으로 정렬해 인덱스가 0과 1인 최솟값끼리 더해주는 것을 반복하면 최적해를 찾을 수 있다
	for j := 0; j < m; j++ {
		sort.Ints(input)
		tmp := input[0]
		input[0] += input[1]
		input[1] += tmp
	}
	sum := 0
	for _, v := range input {
		sum += v
	}
	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
