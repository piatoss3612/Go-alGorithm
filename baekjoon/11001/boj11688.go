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

	a, b, L int
)

// 난이도: Gold 4
// 메모리: 980KB
// 시간: 16ms
// 분류: 정수론, 유클리드 호제법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	a, b, L = scanInt(), scanInt(), scanInt()
}

func Solve() {
	// LCM(a, b, c) = LCM(LCM(a, b), c) = (LCM(a, b) / GCD(LCM(a, b), c)) * c = L
	// 따라서 c는 L의 약수이다.

	divisors := []int{} // L의 약수

	for c := 1; c*c <= L; c++ {
		if L%c == 0 {
			divisors = append(divisors, c)
			if c*c != L {
				divisors = append(divisors, L/c)
			}
		}
	}

	sort.Ints(divisors) // 오름차순 정렬

	k := lcm(a, b) // LCM(a, b) = k

	ans := -1
	// LCM(k, c) = L을 만족하는 최소의 c를 찾는다.
	for _, d := range divisors {
		if lcm(k, d) == L {
			ans = d
			break
		}
	}
	fmt.Fprintln(writer, ans)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
