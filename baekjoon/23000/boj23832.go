package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n       int
)

// 메모리: 908KB
// 시간: 40ms
// 오일러 피 함수를 사용해 1부터 N까지 각 정점과 서로소인 정점의 개수의 총합을 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	ans := 0

	// 1보다 작은 값을 가지면서 서로소인 정점은 없으므로 제외
	for i := n; i >= 2; i-- {
		ans += EulerPhi(i)
	}
	fmt.Fprintln(writer, ans)
}

func EulerPhi(x int) int {
	res := x

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			res = res * (i - 1) / i
			for x%i == 0 {
				x /= i
			}
		}
	}

	if x != 1 {
		res = res * (x - 1) / x
	}

	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
