package bj9655

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
	// SK가 먼저 시작하고 홀수 개(1 또는 3)을 가져갈 수 있으므로 짝수 만큼 남는 경우 반드시 SK가 승리하게 된다
	// 따라서 홀수인 경우에는 SK, 짝수인 경우에는 CY가 승리한다
	if n%2 == 0 {
		fmt.Fprintln(writer, "CY")
	} else {
		fmt.Fprintln(writer, "SK")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
