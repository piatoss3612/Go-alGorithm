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
	k, c    int
)

const MAX = 1000000000

// 메모리: 908KB
// 시간: 4ms
// 확장된 유클리드 호제법을 사용하여 kx+cy=1을 만족하는 y값을 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		k, c = scanInt(), scanInt()

		// k명의 사람들에게 c개의 사탕이 들어있는 사탕 봉지를 y개를 구매한 뒤
		// 공정하게 나누어 주고도  1개의 사탕이 남아야되는 상황
		// 즉, c*y ≡ 1 (mod k)를 만족하는 사탕봉지의 개수 y를 구해야 한다

		// 사탕 봉지에 사탕이 1개(c=1) 들어있는 경우
		if c == 1 {
			// k+1개의 사탕 봉지가 필요
			if k+1 > MAX {
				fmt.Fprintln(writer, "IMPOSSIBLE") // 최댓값의 한계보다 큰 경우
			} else {
				fmt.Fprintln(writer, k+1)
			}
			continue
		}

		// 사람이 1명만 있는 경우
		// 사탕봉지는 1개만 구매하면 된다 (c >= 2)
		// k=1, c=1인 경우는 앞에서 c가 1인 경우를 검사할 때 처리된다
		if k == 1 {
			fmt.Fprintln(writer, 1)
			continue
		}

		// 확장된 유클리드 알고리즘을 사용해 c ≡ 1 (mod k) 연산의 모듈러 곱셈 역원 y를 구한다
		ExEuclidean(k, c)
	}
}

// 확장된 유클리드 호제법
func ExEuclidean(a, b int) {
	var q int
	var y1, y2 int = 0, 1
	for b > 0 {
		q = a / b
		a, b = b, a%b
		y1, y2 = y2, y1-y2*q
	}

	// 최대공약수가 1이 아닌 경우 또는 모듈러 곱셈 역원 y가 MAX보다 큰 경우
	if a != 1 || y1 > MAX {
		fmt.Fprintln(writer, "IMPOSSIBLE")
		return
	}
	fmt.Fprintln(writer, (k+y1)%k) // y1이 음수라면 덧셈의 역원을 출력한다
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
