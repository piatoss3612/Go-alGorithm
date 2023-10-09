package bj6064

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
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	m, n, x, y := scanInt(), scanInt(), scanInt(), scanInt()

	end := lcm(m, n) // 최소공배수를 기점으로 다시 <1:1>부터 반복되는 달력

	// x + m  * i를 기준으로 해를 구할 수 있다
	for i := x; i <= end; i += m {
		if i%n == y%n {
			fmt.Fprintln(writer, i)
			return
		}
	}
	fmt.Fprintln(writer, -1)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
