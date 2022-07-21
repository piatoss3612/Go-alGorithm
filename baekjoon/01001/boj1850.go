package bj1850

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
	a, b := scanInt64(), scanInt64()
	result := gcd(a, b)
	// a개의 1을 가진 수와 b개의 1을 가진 수의 최대공약수는
	// a와 b의 최대공약수 만큼의 1의 개수를 가지고 있다
	for i := 0; int64(i) < result; i++ {
		fmt.Fprint(writer, 1)
	}
	fmt.Fprintln(writer)
}

func scanInt64() int64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return int64(n)
}

func gcd(a, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
