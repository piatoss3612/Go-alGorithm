package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 메모리: 884KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	b := scanBytes()
	b = append([]byte{'0'}, b...) // 연산의 편의를 위해 0의 자리에 더미 추가
	var dp [41][3]int

	/*
		예제 입력:
		27123

		풀이:

		dp[1]: 2 하나로 만들 수 있는 숫자의 경우 = dp[1][1] = 1

		dp[2]: dp[1]에 7 하나를 더해 만들 수 있는 숫자의 경우 = dp[2][1] = dp[1][1] + dp[1][2] = 1 + 0 = 1
			   27을 더해 만들 수 있는 숫자의 경우 = dp[2][2] = dp[1][1] = 1

		dp[2]: dp[2]에 1 하나를 더해 만들 수 있는 숫자의 경우 = dp[3][1] = dp[2][1] + dp[2][2] = 1 + 1 = 2
			   7과 1을 더해 만든 수는 34보다 크므로 dp[2][2] = 0

		...

		dp[5]: dp[4]에 3 하나를 더해 만들 수 있는 숫자의 경우 = dp[5][1] = dp[4][1] + dp[4][2] = 2 + 2 = 4
			   23을 더해 만들 수 있는 숫자의 경우 = dp[5][2] = dp[4][1] = 2

		예제 출력:
		6
	*/

	/*
		예외 입력:
		1101

		풀이:

		dp[1]: dp[1][1] = 1, dp[1][2] = 0
		dp[2]: dp[2][1] = 1, dp[2][2] = 1
		dp[3]: dp[3][1] = 0, dp[3][2] = 1

		0은 그자체로 카드가 아니며, 앞의 수와 조합하여야만 카드로 존재할 수 있는 수
		따라서 기존의 경우의 수에 0만 추가할 수는 없으므로 dp[3][1]은 0이 될 수 밖에 없다
		dp[3][2]는 1과 0을 붙여 [1, 10] 경우의 수를 구할 수 있다

		dp[4]: dp[4][1] = 1, dp[4][2] = 0

		dp[4]의 경우는 앞서 0으로 인해 발생한 예외 상황으로 인해 [1, 10] 경우의 수에 1을 추가하는 양상이 된다
		prev * 10 + cur에서 궂이 예외 처리를 해줄 필요가 없는 이유는 이미 dp[3][1]이 0이므로 dp[4][2]는 결국 0이 되기 때문이다
	*/

	dp[1][1] = 1

	n := len(b) - 1

	for i := 2; i <= n; i++ {
		prev := int(b[i-1] - '0')
		cur := int(b[i] - '0')

		if cur != 0 {
			dp[i][1] += dp[i-1][1] + dp[i-1][2]
		}

		if prev*10+cur <= 34 {
			dp[i][2] += dp[i-1][1]
		}
	}
	fmt.Fprintln(writer, dp[n][1]+dp[n][2])
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
