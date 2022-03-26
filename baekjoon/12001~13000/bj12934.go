package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	x, y := scanInt(), scanInt()

	// x + y는 1부터 i까지의 합: i * (i + 1) / 2
	// 즉, (x + y) * 2 = i * (i + 1)
	// i * (i + 1)은 i * i와 (i + 1) * (i + 1) 사이에 있으므로
	// 제곱근을 구하면 i.xxxx이므로 소숫점을 버리면 i를 구할 수 있다
	n := int(math.Sqrt(float64((x + y) * 2)))

	if (x+y)*2 != n*(n+1) {
		fmt.Fprintln(writer, -1)
		return
	}

	// 그리디 알고리즘:
	// x가 최소로 이겨야하는 경우를 구하는 방법은 뺄 수 있는 가장 큰 수부터 빼는 것이다
	ans := 0
	for i := n; i > 0; i-- {
		if x-i >= 0 {
			x -= i
			ans += 1
		}

		if x == 0 {
			break
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
