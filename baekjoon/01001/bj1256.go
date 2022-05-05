package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	bino    [201][201]int    // 이항 계수를 저장하는 배열
	M       = 1000000000 + 1 // k의 최댓값이 10억이므로 이항 계수의 오버플로우가 발생하지 않게 10억보다 큰 값은 10억 1로 할당한다
)

/*
이항 계수를 사용하는 이유:

1# a 2개, z 2개로 사전을 만든다고 가정.
a 2개, z 2개로 만들 수 있는 경우의 수는 사전 순으로 aazz, azaz, azza, zaaz, zaza, zzaa 6개, 이항 계수 4C2와 같다

2# 사전의 4번째 문자열을 구한다고 가정.

skip(k) = 3

a로 시작하는 문자열은 3C1(aazz, azaz, azza), 3가지가 있다
우리는 4번째 문자열을 구하고 싶으므로 a로 시작하는 3개의 문자열은 궂이 구해보지 않아도 된다는 것을 알 수 있다

즉, 모든 문자열을 구하지 않고 a로 시작하는 문자열의 수 만큼 건너뜀(skip-3)으로써
skip(k)는 0이라는 숫자, 즉 우리가 구하고 싶은 문자열의 순서에 도달하게 된다

결과적으로 a로 시작하는 문자열이 아닌, z로 시작하는 문자열의 사전 순서 첫번째에 있는 문자열을 출력한다
*/
func init() {
	bino[0][0] = 1
	for i := 1; i <= 200; i++ {
		bino[i][0] = 1
		bino[i][i] = 1
		for j := 1; j < i; j++ {
			bino[i][j] = min(bino[i-1][j-1]+bino[i-1][j], M)
		}
	}
}

// 메모리: 1312KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()

	skip := scanInt() - 1 // 사전의 k번째 문자열은 곧 사전에서 1번부터 k-1개 만큼 건너뛴 경우

	// n개의 a와 m개의 z로 만들 수 있는 문자열의 경우의 수는 이항계수 bino[n+m][n] = n+m C n이다
	// 이 값이 skip(k)보다 작거나 같으면 k번째 문자열은 사전의 범위를 벗어나 만들 수 없는 문자열인 셈이다
	// 같은 경우를 제외하면 런타임에러 발생
	if skip >= bino[n+m][n] {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, kth(n, m, skip))
}

// k번째 문자열을 구하는 재귀 함수
func kth(n, m, skip int) string {
	// a의 개수가 0이라면 z를 m개만큼 반복한 문자열을 반환
	if n == 0 {
		return strings.Repeat("z", m)
	}

	// a로 시작하는 문자열에서 답을 찾을 수 있는 경우
	if skip < bino[n+m-1][n-1] {
		return "a" + kth(n-1, m, skip)
	}
	// a로 시작하는 문자열을 궂이 탐색하지 않아도 되는 경우
	return "z" + kth(n, m-1, skip-bino[n+m-1][n-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
