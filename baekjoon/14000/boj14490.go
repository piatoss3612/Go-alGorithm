package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	input []byte
)

// 14490번: 백대열
// https://www.acmicpc.net/problem/14490
// 난이도: 실버 5
// 메모리: 852 KB
// 시간: 4 ms
// 분류: 수학, 문자열, 정수론, 유클리드 호제법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	input = scanBytes()
}

func Solve() {
	nums := bytes.Split(input, []byte{':'})

	a, b := bytesToInt(nums[0]), bytesToInt(nums[1])

	g := gcd(a, b)

	fmt.Fprintf(writer, "%d:%d\n", a/g, b/g)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func bytesToInt(b []byte) int {
	res := 0
	for _, v := range b {
		res = res*10 + int(v-'0')
	}
	return res
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}