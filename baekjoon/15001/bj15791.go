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

const P = 1000000007

// 메모리: 904KB
// 시간: 8ms
// 페르마의 소정리를 사용해 이항 계수를 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()

	if n == m {
		fmt.Fprintln(writer, 1)
		return
	}

	var a, b, c int
	a = 1
	for i := 1; i <= n; i++ {
		a *= i
		a %= P

		if i == n-m {
			b = a
		}

		if i == m {
			c = a
		}
	}

	// a * (b * c)^P-2 ≡ nCm (mod P)
	ans := rec((b*c)%P, P-2) % P // 주의# b*c도 %P 연산을 해줘야만 한다
	ans *= a
	ans %= P
	fmt.Fprintln(writer, ans)
}

// 모듈러 곱셈 역원을 구하는 함수
func rec(x, y int) int {
	res := 1
	for y > 0 {
		if y%2 != 0 {
			res *= x
			res %= P
		}
		x *= x
		x %= P
		y /= 2
	}
	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
