package main

import (
	"bufio"
	_ "bytes"
	_ "container/heap"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	factorial [4000001]int
	inverse   [4000001]int
)

// 상수로 나머지 연산을 실행하니 무려 전체 시간이 244 -> 160ms로 줄어들었다!
const P = 1000000007

// 전처리 함수
func init() {
	// 1!~4000000! 팩토리얼 수를 전처리
	factorial[0], factorial[1] = 1, 1
	for i := 2; i <= 4000000; i++ {
		factorial[i] = (factorial[i-1] * i) % P
	}
	// 1!~4000000!의 모듈러 곱셈 역원을 전처리
	// (n!)^-1 = (n+1)*((n+1)!)^-1
	inverse[4000000] = rec(factorial[4000000], P-2)
	for i := 3999999; i >= 1; i-- {
		inverse[i] = (inverse[i+1] * (i + 1)) % P
	}
}

// x^y 분할-제곱 연산 함수
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

// 메모리: 65916KB
// 시간: 160ms
// 페르마의 소정리를 사용해 이항 계수를 구하는 문제
// 11401번 문제와 동일
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()

	var n, k, ans int
	for i := 1; i <= t; i++ {
		n, k = scanInt(), scanInt()
		// n과 k가 같거나 k가 0인 경우
		// 즉, nCn 또는 nC0인 경우 1을 출력
		if n == k || k == 0 {
			fmt.Fprintln(writer, 1)
			continue // 바보같이 break를 써버려서 몇 번 틀렸다...
		}
		// nCk = (n! * (n-k)!의 모듈려 곱셈 역원 * k!의 모듈려 곱셈 역원) % P
		ans = (inverse[n-k] * inverse[k]) % P
		ans = (ans * factorial[n]) % P
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
