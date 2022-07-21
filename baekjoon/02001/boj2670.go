package bj2670

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
	inputs := make([]float64, n)

	for i := 0; i < n; i++ {
		inputs[i] = scanFloat()
	}

	for j := 1; j < n; j++ {
		tmp := inputs[j] * inputs[j-1]
		if tmp > inputs[j] {
			inputs[j] = tmp
		}
	}
	var max float64 = 0

	for k := 0; k < n; k++ {
		if inputs[k] > max {
			max = inputs[k]
		}
	}

	fmt.Fprintf(writer, "%.3f\n", max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
