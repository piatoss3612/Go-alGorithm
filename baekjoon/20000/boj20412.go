package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	scanner         = bufio.NewScanner(os.Stdin)
	writer          = bufio.NewWriter(os.Stdout)
	M, Seed, X1, X2 int
)

// 난이도: Gold 2
// 메모리: 956KB
// 시간: 8ms
// 분류: 수학, 정수론, 페르마의 소정리, 모듈러 연산
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	M, Seed, X1, X2 = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	// (a * Seed + c) = X1 (mod M)
	// (a * X1 + c) = X2 (mod M)
	// a * (X1 - Seed) = X2 - X1 (mod M)

	// (X1 - Seed)는 항상 M보다 작고 M은 소수이므로 (X1 - Seed)와 M은 서로소
	// 따라서 페르마의 소정리에 의해 (X1 - Seed)^(M-1) = 1 (mod M)
	// (X1 - Seed)^-1 = (X1 - Seed)^(M-2) (mod M)

	// a = (X2 - X1) / (X1 - Seed) (mod M)
	// a = (X2 - X1) * (X1 - Seed)^-1 (mod M)
	bigM := big.NewInt(int64(M))

	temp1 := new(big.Int).Mod(big.NewInt(int64(X2-X1)), bigM)
	temp2 := new(big.Int).Mod(big.NewInt(int64(X1-Seed)), bigM)

	invTemp2 := new(big.Int).ModInverse(temp2, bigM)

	// temp2가 0이면 invTemp2가 nil이 되어버림
	if invTemp2 == nil {
		invTemp2 = big.NewInt(0)
	}

	a := new(big.Int).Mod(big.NewInt(0).Mul(temp1, invTemp2), bigM)

	// c = X1 - (a * Seed) % M
	c := new(big.Int).Mod(new(big.Int).Sub(big.NewInt(int64(X1)), new(big.Int).Mul(a, big.NewInt(int64(Seed)))), bigM)

	fmt.Fprintf(writer, "%d %d\n", a, c)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
