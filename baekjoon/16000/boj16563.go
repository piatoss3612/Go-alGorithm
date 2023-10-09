package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	minFactor [5000001]int // 5000000보다 작거나 같은 수 i의 가장 작은 소인수를 저장
)

// minFactor 슬라이스 전처리
func init() {
	eratosthenes()
}

// 메모리: 47720KB -> 46308KB
// 시간: 1164ms -> 1048ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	n := scanInt()

	for i := 1; i <= n; i++ {
		factorization(scanInt())
	}
}

// 에라토스테네스의 체
// i가 소수인 경우 minFactor[i]는 i 자기자신,
// 그렇지 않은 경우 minFactor[i]는 i의 가장 작은 소인수가 된다
func eratosthenes() {
	minFactor[0] = -1
	minFactor[1] = -1

	// i의 가장 작은 소인수를 자기자신으로 초기화
	for i := 2; i <= 5000000; i++ {
		minFactor[i] = i
	}

	sqrtn := int(math.Sqrt(5000000))
	for i := 2; i <= sqrtn; i++ {
		// i가 소수인 경우
		if minFactor[i] == i {
			// 5000000보다 작거나 같은 i의 배수 j는 소수가 아니므로
			// 체로 걸러내는데, i의 배수라는 것은 j의 가장 작은 소인수가 i라는 뜻이다
			// j가 i*i부터 시작하는 이유는 불필요한 연산을 줄이기 위해서 이다
			for j := i * i; j <= 5000000; j += i {
				// j의 가장 작은 소인수가 갱신되지 않은 경우
				if minFactor[j] == j {
					// j의 가장 작은 소인수를 i로 갱신
					minFactor[j] = i
				}
			}
		}
	}
}

// 소인수 구하기
func factorization(x int) {
	/*
		factors := make([]int, 0, 0)
		for x > 1 {
			factors = append(factors, minFactor[x])
			x /= minFactor[x]
		}

		for _, f := range factors {
			fmt.Fprintf(writer, "%d ", f)
		}
		fmt.Fprintln(writer)
	*/

	// 정렬이 필요한가? -> 정렬은 필요하지 않았다

	// 왜?
	// minFactor[x]는 x의 가장 작은 소인수이므로
	// x를 minFactor[x]로 나눈 몫인 x'의 minFactor[x']는
	// 항상 minFactor[x]보다 크거나 같다

	// 따라서 소인수를 저장하고 정렬하는 과정없이 바로 출력함으로써 시간과 메모리를 절약할 수 있다

	for x > 1 {
		fmt.Fprintf(writer, "%d ", minFactor[x])
		x /= minFactor[x]
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
