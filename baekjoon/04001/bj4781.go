package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 캔디의 칼로리와 가격
type candy struct {
	calories, price int
}

// 메모리: 9680KB
// 시간: 424ms
// 다이나믹 프로그래밍, 가격을 실수로 입력받는 것이 번거로운 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		N, floatM := scanInt(), scanFloat()
		// 프로그램 탈출
		if N == 0 && floatM == 0.00 {
			return
		}
		money := floatToInt(floatM) // 가지고 있는 돈을 정수형으로 변환 -> 실수 연산의 오차로 인한 문제 방지

		candys := make([]candy, N+1)
		for i := 1; i <= N; i++ {
			candys[i] = candy{scanInt(), floatToInt(scanFloat())}
		}

		dp := make([]int, money+1)

		// 가지고 있는 돈으로 얻을 수 있는 캔디의 최대 칼로리를 계산
		for i := 1; i <= N; i++ {
			c, p := candys[i].calories, candys[i].price
			for j := 0; j <= money; j++ {
				if j-p >= 0 {
					dp[j] = max(dp[j], dp[j-p]+c)
				}
			}
		}

		// dp를 내림차순으로 정렬
		sort.Slice(dp, func(i, j int) bool {
			return dp[i] > dp[j]
		})

		fmt.Fprintln(writer, dp[0])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}

func floatToInt(n float64) int {
	return int(n*100 + 0.5)
}
