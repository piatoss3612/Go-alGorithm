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
	A, B, C float64
)

// 메모리: 912KB
// 시간: 4ms
// 이분 탐색, 수치 해석
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	A, B, C = scanFloat64(), scanFloat64(), scanFloat64()

	err := math.Pow(10, -9) // 오차 범위 10^-9

	// sin(x)가 표현하는 값의 범위가 -1~1이므로
	// Ax+Bsin(x)=C를 만족하는
	// x의 최솟값은 (C - B) / A
	lo := (C - B) / A
	// x의 최댓값은 (C + B) / A
	hi := (C + B) / A

	// 이분 탐색
	// 두 개의 양끝값 lo, hi를 사용하여 중간값 mid를 구하고
	// mid가 Ax+Bsin(x)=C를 만족하는 x인지 판별한다
	for {
		mid := (lo + hi) / 2
		temp := A*mid + B*math.Sin(mid)
		if temp > C {
			hi = mid
		} else {
			lo = mid
		}

		// 반복문 종료 조건:
		// temp와 C의 오차의 절댓값이 err 이하인 경우
		if math.Abs(temp-C) <= err {
			fmt.Fprintln(writer, mid)
			return
		}
	}
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
