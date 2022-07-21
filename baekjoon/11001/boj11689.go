package main

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
)

// 메모리: 912KB
// 시간: 16ms
// GCD(n,k) = 1, 즉 n보다 작으면서 n과 서로소인 수 k의 개수를 찾는 문제

// 오일러 피 함수:
// n 이하의 양의 정수 중 n과 서로소인 양의 정수의 개수는
// n * (1 - 1/p1) * (1 - 1/p2) * ... (1 - 1/pn)
// p1, p2, ..., pn은 n의 소인수
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	ans := n

	// 소인수 분해
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			ans /= i
			ans *= i - 1
		}
	}

	// 소인수가 남아있는 경우
	if n != 1 {
		ans /= n
		ans *= n - 1
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
