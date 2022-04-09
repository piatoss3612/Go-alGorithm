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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	// 음수(0 포함)와 양수의 곱은 의미가 없으므로
	// 계산의 편의를 위해 양수와 음수(0 포함)를 따로 입력받아 정렬했다
	positive := []int{}
	negative := []int{}

	for i := 1; i <= n; i++ {
		m := scanInt()
		if m > 0 {
			positive = append(positive, m)
		} else {
			negative = append(negative, m)
		}
	}

	sort.Ints(negative)
	sort.Slice(positive, func(i, j int) bool {
		return positive[i] > positive[j]
	})

	ans := 0
	var a, b, tmp int
	for len(negative) > 0 {
		a = negative[0]
		negative = negative[1:]

		if len(negative) == 0 {
			ans += a
			break
		}

		b = negative[0]
		negative = negative[1:]

		tmp = a * b
		// 음수(0 포함)끼리의 곱은 양수 또는 0
		// 0인 경우를 고려하여 tmp가 a와 b보다 크거나 같은 경우에 묶는다
		if tmp >= a && tmp >= b {
			ans += tmp
		} else {
			ans += a + b
		}
	}

	for len(positive) > 0 {
		a = positive[0]
		positive = positive[1:]

		if len(positive) == 0 {
			ans += a
			break
		}

		b = positive[0]
		positive = positive[1:]

		tmp = a * b
		// 양수의 곱 tmp는 항상 a와 b보다 커야 묶을 수 있다
		if tmp > a && tmp > b {
			ans += tmp
		} else {
			ans += a + b
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
