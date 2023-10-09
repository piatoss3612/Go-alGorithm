package bj3036

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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = scanInt()
	}
	// 첫번째 원의 반지름에 비례하여 돌아가므로
	// 첫번째 원의 반지름 : i번째 원의 반지름을 최대공약수로 나누어
	// a/b의 형태로 출력
	for i := 1; i < n; i++ {
		tmp := gcd(arr[0], arr[i])
		fmt.Fprintf(writer, "%d/%d\n", arr[0]/tmp, arr[i]/tmp)
	}
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
