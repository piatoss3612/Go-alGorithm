package bj3053

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
	n := scanInt()
	// 유클리드 기하학 - 원의 넓이: 반지름 * 반지름 * PI
	// 택시 기하학 - 원의 넓이: 반지름 * 반지름 * 2
	fmt.Fprintf(writer, "%0.6f\n", float64(n*n)*math.Pi)
	fmt.Fprintf(writer, "%0.6f\n", float64(n*n)*2)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
