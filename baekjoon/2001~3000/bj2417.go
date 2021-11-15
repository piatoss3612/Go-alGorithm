package bj2417

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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	result := int(math.Floor(math.Sqrt(float64(n))))
	if result*result < n {
		result += 1
	}
	fmt.Fprintf(writer, "%d\n", result)
}
