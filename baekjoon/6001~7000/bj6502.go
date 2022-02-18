package bj6502

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
	r, w, l float64
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	cnt := 1
	for {
		r = scanFloat64()

		if r == 0 {
			return
		}

		w, l = scanFloat64(), scanFloat64()

		diagonal := math.Sqrt(w*w + l*l)

		if r*2 >= diagonal {
			fmt.Fprintf(writer, "Pizza %d fits on the table.\n", cnt)
		} else {
			fmt.Fprintf(writer, "Pizza %d does not fit on the table.\n", cnt)
		}
		cnt += 1
	}
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
}
