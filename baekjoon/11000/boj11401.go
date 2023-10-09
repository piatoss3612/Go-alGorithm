package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	n, k      int
	factorial [4000001]int // 1~n까지의 팩토리얼 값을 저장하는 배열
	inverse   map[int]int  // (n-k)!k!의 역원을 구하기 위한 과정에서 분할 제곱값을 저장하는 맵
	p         = 1000000007 // nCk를 나누는 값, 소수
)

// 메모리: 31896KB
// 시간: 88ms
// 이항 계수 nCk를 p(=1000000007)로 나누는 문제, 페르마의 소정리

// p가 소수이므로
// 페르마의 소정리에 따라 nCk^p ≡ nCk (mod p) 이다

// 또한 nCk는 p를 소인수로 포함하지 않으므로 p와 nCk는 서로소이다
// 따라서 nCk^p-1 = 1 (mod p) 이다

// nCk = n!÷(n-k)!k!이며 n!^p ≡ n! (mod p) 이다
// 그리고 (a/b)%p == (a%p)/(b%p)가 성립하지 않으므로
// (n-k)!k!의 모듈러 곱셈 역원을 구하여 n!과 곱해주어야 한다

// (n-k)!k! -> A로 치환하면 A와 p가 서로소이므로 A^p-1 ≡ 1 (mod p) 이다
// A^p-1은 곧 A * A^p-2이므로 A의 모듈러 역원은 A^p-2 이다

// 즉, nCk = (n!*((n-k!)*k!)^p-2) % p 이다

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k = scanInt(), scanInt()
	factorial[0], factorial[1] = 1, 1
	for i := 2; i <= n; i++ {
		factorial[i] = (factorial[i-1] * i) % p
	}

	inverse = make(map[int]int)
	inverse[1] = (factorial[n-k] * factorial[k]) % p

	ans := (factorial[n] * rec(p-2)) % p

	fmt.Fprintln(writer, ans)
}

// 분할, 거듭 제곱을 사용해 (n-k)!k!의 역원을 구하는 함수
func rec(x int) int {
	_, ok := inverse[x]
	if ok {
		return inverse[x]
	}

	if x%2 == 0 {
		_, ok = inverse[x/2]
		if !ok {
			inverse[x/2] = rec(x / 2)
		}
		inverse[x] = (inverse[x/2] * inverse[x/2]) % p
		return inverse[x]
	}

	_, ok = inverse[x-1]
	if !ok {
		inverse[x-1] = rec(x - 1)
	}
	inverse[x] = (inverse[x-1] * inverse[1]) % p
	return inverse[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
