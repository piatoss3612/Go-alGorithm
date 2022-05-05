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
	catalan [2501]int // 길이가 5000인 올바른 괄호문자열의 개수는 2500번째 카탈란 수
	// bino [5001][5001]int // n개의 카탈란 수를 구하기 위해 2n으로 시작하는 조합 필요
	MOD = 1000000007
)

// 카탈란 수 나무위키: https://namu.wiki/w/%EC%B9%B4%ED%83%88%EB%9E%91%20%EC%88%98

// ###카탈란 수를 활용하는 문제###

/*
카탈란 수 구하기: 조합론
메모리: 117584KB
시간: 188ms

n번째 카탈란 수 = 2n C n - 2n C n-1
조합(이항 계수)을 먼저 구하고, 카탈란 수를 구한다
*/

// func init() {
//   bino[0][0] = 1
//   for i := 1; i <= 5000; i++ {
//     bino[i][0] = 1
//     bino[i][i] = 1
//     for j := 1; j < i; j++ {
//       bino[i][j] = (bino[i-1][j-1] + bino[i-1][j]) % MOD
//     }
//   }

//   catalan[0] = 1
//   for i := 1; i <= 2500; i++ {
// 이항 계수를 구할 때 나머지 연산을 실행했으므로
// 음수가 나오는 문제를 방지하기 위해 MOD를 한 번 더해주고 다시 나머지 연산을 해준다
//     catalan[i] = (bino[2*i][i] + MOD - bino[2*i][i-1]) % MOD
//   }
// }

/*
카탈란 수 구하기: 점화식
메모리: 932KB
시간: 56ms
*/

func init() {
	catalan[0] = 1
	catalan[1] = 1

	for i := 2; i <= 2500; i++ {
		for j := 0; j < i; j++ {
			catalan[i] += (catalan[j] * catalan[i-j-1])
			catalan[i] %= MOD
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()

	var n int

	for i := 1; i <= t; i++ {
		n = scanInt()

		// 길이가 홀수인 올바른 괄호문자열은 존재할 수 없다
		// 길이가 짝수인 경우에만 길이/2 번째 카탈란 수를 결과로 출력한다
		if n%2 == 0 {
			fmt.Fprintln(writer, catalan[n/2])
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
