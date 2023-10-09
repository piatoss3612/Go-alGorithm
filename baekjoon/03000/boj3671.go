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

	c       int
	isPrime [MAX]bool // 소수인지 아닌지 판별
	checked [MAX]bool // 이미 생성한 소수인지 아닌지 판별
	digits  [10]int // 종이 조각에 적힌 숫자의 개수
)

const MAX = 10000000

// 난이도: Gold 4
// 메모리: 20420KB
// 시간: 100ms
// 분류: 브루트포스, 백트래킹, 에라토스테네스의 체, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	eratosthenes() // 에라토스테네스의 체를 이용하여 1 ~ 10000000까지의 소수를 판별
}

func Solve() {
	c = scanInt()

	for i := 1; i <= c; i++ {
		checked = [MAX]bool{}
		digits = [10]int{}
		b := scanBytes()
		for _, v := range b {
			digits[v-'0'] += 1
		}
		fmt.Fprintln(writer, rec(0, len(b)))
	}
}

func rec(num, remain int) int {
	// 기저 사례: 종이 조각을 전부 사용했을 경우
	if remain == 0 {
		return 0
	}

	ret := 0
	for i := 0; i < 10; i++ {
		// i가 적힌 종이 조각이 남아 있을 경우
		if digits[i] > 0 {
			digits[i] -= 1
			next := num*10 + i // 종이 조각에 적힌 숫자를 이어 붙임
			// next가 소수이고, 아직 생성하지 않은 소수일 경우
			if isPrime[next] && !checked[next] {
				ret += 1
				checked[next] = true
			}
			ret += rec(next, remain-1) // 재귀 호출
			digits[i] += 1
		}
	}
	return ret
}

func eratosthenes() {
	for i := 2; i < MAX; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < MAX; i++ {
		if isPrime[i] {
			for j := i * i; j < MAX; j += i {
				isPrime[j] = false
			}
		}
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
