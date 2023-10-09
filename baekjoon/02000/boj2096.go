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
	n := scanInt()
	input := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = make([]int, 3)
		input[i][0] = scanInt()
		input[i][1] = scanInt()
		input[i][2] = scanInt()
	}

	min1, min2, min3 := input[1][0], input[1][1], input[1][2]
	max1, max2, max3 := input[1][0], input[1][1], input[1][2]
	var tmpMin1, tmpMin2, tmpMin3 int
	var tmpMax1, tmpMax2, tmpMax3 int

	for i := 2; i <= n; i++ {
		tmpMin1 = getMin(input[i][0]+min1, input[i][0]+min2)
		tmpMin2 = getMin(getMin(input[i][1]+min1, input[i][1]+min2), input[i][1]+min3)
		tmpMin3 = getMin(input[i][2]+min2, input[i][2]+min3)

		tmpMax1 = getMax(input[i][0]+max1, input[i][0]+max2)
		tmpMax2 = getMax(getMax(input[i][1]+max1, input[i][1]+max2), input[i][1]+max3)
		tmpMax3 = getMax(input[i][2]+max2, input[i][2]+max3)

		min1, min2, min3 = tmpMin1, tmpMin2, tmpMin3
		max1, max2, max3 = tmpMax1, tmpMax2, tmpMax3
	}

	max := getMax(getMax(max1, max2), max3)
	min := getMin(getMin(min1, min2), min3)

	fmt.Fprintf(writer, "%d %d\n", max, min)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
