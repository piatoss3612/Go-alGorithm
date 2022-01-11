package bj11867

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
	n, m := scanInt(), scanInt()
	// 2개의 박스 중 하나라도 짝수인 갯수의 돌이 들어있다면
	// 먼저 시작한 사람은 반드시 1,1을 만들 수 있다
	if n%2 == 0 || m%2 == 0 {
		fmt.Fprintln(writer, "A")
	} else {
		fmt.Fprintln(writer, "B")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
