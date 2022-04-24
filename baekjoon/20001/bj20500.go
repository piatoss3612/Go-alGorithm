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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	dp := make([]int, 1516)

	dp[2] = 1
	dp[3] = 1

	/*
		15의 배수는 3의 배수(자릿수의 합이 3으로 나누어 떨어짐)이면서 5의 배수
		1과 5로 구성된 수이므로 1의 자릿수는 반드시 5로 고정

		dp[1] = 0
		dp[2] = 1 => 15
		dp[3] = 1 => 555

		dp[i]는...
		dp[2]부터 dp[i - 2]까지 각각의 수를 두 배한 값을 누적해서 더한 값 + 1

		왜???

		예시)
		dp[4] = 3 => 1515, 1155, 5115
		dp[4]의 1515는 dp[2]의 15의 뒤에 15를 붙여 만든 15의 배수
		dp[4]의 1155는 dp[2]의 15의 앞과 뒤에 1과 5를 붙여 만든 15의 배수
		이런 식으로 15의 배수를 만들 수 있다
		마지막으로 1을 더한 것은 자릿수 배치나 조합으로 인해 추가적으로 발생하는 경우의 수
	*/

	for i := 4; i <= 1515; i++ {
		tmp := 0
		for j := 2; j <= i-2; j++ {
			tmp = (tmp + 2*dp[j]) % 1000000007
		}
		dp[i] += tmp + 1
	}

	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
