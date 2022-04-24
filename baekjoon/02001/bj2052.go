package bj2052

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	f := big.NewFloat(math.Pow(2, float64(-n)))

	s := strings.TrimRight(f.Text('f', 300), "0")

	fmt.Println(s)
}
