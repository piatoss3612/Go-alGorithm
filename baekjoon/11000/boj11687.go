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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	// 시간 초과
	// for i := n * 5; i >= 0; i -= 5 {
	// 	sum := 0
	// 	for j := 5; j <= i; j *= 5 {
	// 		sum += i / j
	// 	}
	// 	if sum <= n {
	// 		if sum == n {
	// 			fmt.Fprintln(writer, i)
	// 		} else {
	// 			fmt.Fprintln(writer, -1)
	// 		}
	// 		return
	// 	}
	// }

	l, r := 1, n*5
	var mid, sum, ans int

	/*
		0을 가지려면 10으로 나누어 떨어져야 하는 수
		즉, 2와 5를 소인수로 가지고 있으면서 두 수의 갯수 중 작은 수가 0의 갯수가 된다

		n!이 끝에 0을 가지려면 2의 갯수는 부족함이 없기 때문에
		5를 소인수로 가지기 위해 5의 배수가 되어야 하며
		소인수로 5를 가지는 만큼 n!의 끝에 0을 가질 수 있다

		예시:
		5! -> 5 / 5 = 1개의 0
		10! -> 10 / 5 = 2개의 0
		15! -> 15 / 5 = 3개의 0
		...
		25! -> 25 / 5 + 25 / 25 = 6개의 0
		...
		125! -> 125 / 5 + 125 / 25 + 125 / 125 = 31개의 0
	*/

	for l <= r {
		sum = 0
		mid = (l + r) / 2
		for i := 5; i <= mid; i *= 5 {
			sum += mid / i
		}

		if n <= sum {
			if n == sum {
				// ans를 최솟값으로 갱신
				ans = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if ans != 0 {
		fmt.Fprintln(writer, ans)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
