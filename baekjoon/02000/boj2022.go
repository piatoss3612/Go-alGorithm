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
	x, y, c float64
)

// 메모리: 900KB
// 시간: 4ms
// 이분 탐색, 피타고라스 정리
// 참고: https://sw-ko.tistory.com/145
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	x, y, c = scanFloat(), scanFloat(), scanFloat()
	// 실수 연산이므로 l을 0으로 설정해도 divided by zero 런타임 오류가 발생하지 않음
	//
	var l, r float64 = 0, min(x, y)

	// x, y, c가 각각 30억 이하이므로 100회 반복만으로도 오차를 10^-3이하로 줄일 수 있다
	for i := 0; i < 100; i++ {
		mid := (l + r) / 2 // 두 빌딩 사이의 간격

		a := math.Sqrt(x*x - mid*mid) // 대각선 x가 포함된 삼각형의 높이
		b := math.Sqrt(y*y - mid*mid) // 대각선 y가 포함된 삼각형의 높이

		// 삼각형의 비율을 이용하여 임의의 mid에 대한 c, tempC를 구할 수 있다
		tempC := (a * b) / (a + b)

		// tempC가 c보다 큰 경우
		if tempC > c {
			// c의 길이를 줄여야 하므로 두 빌딩 사이의 간격을 늘려야 한다
			l = mid
		} else {
			// c의 길이를 늘려야 하므로 두 빌딩 사이의 간격을 줄여야 한다
			r = mid
		}
	}

	fmt.Fprintf(writer, "%0.3f\n", l) // 소숫점 3번째 자리까지 출력
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
