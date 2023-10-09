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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	gcd, lcm := scanInt(), scanInt()

	ab := lcm * gcd

	sqt := int(math.Floor(math.Sqrt(float64(ab)))) // a*b의 제곱근 -> 연산 횟수를 줄일 수 있다

	sum := 200000000
	var a, b int

	for i := sqt; i >= gcd; i-- {
		if i%gcd == 0 && ab%i == 0 { // i가 최대공약수로 나누어 떨어지고 a*b가 i로 나누어 떨어지는 경우
			tmp := ab / i
			if GCD(i, tmp) == gcd { // i와 tmp의 최대공약수가 gcd여야만 한다
				if i+tmp < sum {
					sum = i + tmp
					if tmp > i {
						a, b = i, tmp
					} else {
						a, b = tmp, i
					}
				}
			}
		}
	}
	fmt.Fprintln(writer, a, b)
}

// 최대공약수를 구하는 유클리드 호제법
func GCD(a, b int) int {
	if b > a {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
