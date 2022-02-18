package bj1064

import (
	"bufio"
	"fmt"
	"math"
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
	xa, ya := scanFloat64(), scanFloat64()
	xb, yb := scanFloat64(), scanFloat64()
	xc, yc := scanFloat64(), scanFloat64()

	// 정점들이 같은 직선 위에 있는지 기울기를 통해 확인
	if (ya-yb)*(xb-xc) == (yb-yc)*(xa-xb) {
		fmt.Fprintf(writer, "%0.1f\n", float64(-1))
		return
	}

	line1 := math.Sqrt((xc-xa)*(xc-xa) + (yc-ya)*(yc-ya))
	line2 := math.Sqrt((xb-xa)*(xb-xa) + (yb-ya)*(yb-ya))
	line3 := math.Sqrt((xc-xb)*(xc-xb) + (yc-yb)*(yc-yb))

	case1 := (line1 + line2) * float64(2)
	case2 := (line2 + line3) * float64(2)
	case3 := (line3 + line1) * float64(2)

	max := getMax(getMax(case1, case2), case3)
	min := getMin(getMin(case1, case2), case3)

	fmt.Fprintln(writer, max-min)
}

func getMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
}
