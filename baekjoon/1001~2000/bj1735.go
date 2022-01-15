package bj1735

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
	a1, a2 := scanInt(), scanInt()
	b1, b2 := scanInt(), scanInt()
	c1, c2 := a1*b2+b1*a2, a2*b2 // a와 b의 합인 c
	d := gcd(c1, c2)             // c의 분자와 분모의 최대 공약수
	fmt.Fprintln(writer, c1/d, c2/d)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
