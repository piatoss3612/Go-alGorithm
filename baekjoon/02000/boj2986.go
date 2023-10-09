package bj2986

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
	f := 1 // 소수인 경우는 1을 빼야하므로 초깃값을 1로 설정
	for i := 2; i*i <= n; i++ {
		// 처음으로 나누어지는 수를 찾으면
		// n을 그 수로 나누면 n이 가진 가장 큰 소인수를 구할 수 있다
		if n%i == 0 {
			f = n / i
			break
		}
	}
	// 결과는 n에서 가장 큰 소인수를 뺀 값
	fmt.Fprintln(writer, n-f)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
