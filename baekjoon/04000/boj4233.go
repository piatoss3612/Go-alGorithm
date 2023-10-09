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
	temp    map[int]int
)

// 메모리: 1080KB
// 시간: 8ms
// 분할 정복, 페르마의 소정리에서 설명하는 유사소수를 찾는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		p, a := scanInt(), scanInt()

		if p == 0 && a == 0 {
			return
		}

		// 1. 우선적으로 소수 판별
		if isPrime(p) {
			fmt.Fprintln(writer, "no")
			continue
		}

		// 2. 분할 정복 과정에서 생성되는 값을 저장할 맵 초기화
		temp = make(map[int]int)
		temp[1] = a

		// 3. 분할 정복
		v := DAC(p, a, p)

		if v == a {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}

// 소수 판별 함수
func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 분할 정복 함수
func DAC(p, a, mod int) int {
	if v, ok := temp[p]; ok {
		return v
	}

	if p%2 == 0 {
		v := DAC(p/2, a, mod)
		temp[p] = (v * v) % mod
		return temp[p]
	}

	temp[p] = (DAC(p-1, a, mod) * DAC(1, a, mod)) % mod
	return temp[p]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
