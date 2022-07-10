package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	sum     [56]int // i개의 비트 수 이하로 표현할 수 있는 모든 수가 포함한 1의 개수
	A, B    int
)

// sum 전처리
// 2^54 < 10^16 < 2^55
func init() {
	//
	sum[1] = 1
	// sum[2] = 3
	// for i := 3; i <= 55; i++ {
	//   sum[i] = sum[i-1] * 2 + 1 << (i-2)
	// }
	// for i := 2; i <= 55; i++ {
	//   sum[i] += sum[i-1]
	// }

	// 주석으로 남긴 위의 코드를 한 번의 반복문으로 개선
	for i := 2; i <= 55; i++ {
		sum[i] = sum[i-1]*2 + 1<<(i-1)
	}
}

// 메모리: 928KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	A, B = scanInt(), scanInt()
	// (1~B까지 비트 1의 개수의 합) - (1~A-1까지 비트 1의 개수의 합)
	fmt.Fprintln(writer, check(B)-check(A-1))
}

// 시간 초과가 나는 코드
// 로직은 아래의 check 함수와 같은데 시간복잡도가 log2N * log2N 이라서 시간초과가 난걸까?

// func check(n int) int {
//   res := 0

//   for n > 0 {
//     k := int(math.Log2(float64(n)))

//     res += sum[k] + n - (1 << k) + 1
//     n -= 1 << k
//   }

//   return res
// }

// n이하의 모든 수를 비트로 표현했을 때 포함되는 모든 1의 개수를 구하는 함수
func check(n int) int {
	// n을 비트로 표현했을 때 각 자리에 해당하는 값을 슬라이스로 변환
	// 자릿수를 1부터 세기 위해서 0번째 자리에 더미값을 넣어준다
	bit := []int{0}
	temp := n
	for temp > 0 {
		bit = append(bit, temp%2)
		temp /= 2
	}

	res := 0
	// 최상위 비트부터 탐색
	for i := len(bit) - 1; i >= 1; i-- {
		// i번째 비트가 1이라면
		if bit[i] == 1 {
			// sum[i-1]: i번째 비트가 1이라면 i-1개의 비트 수 이하로 표현되는 값을 반드시 표현할 수 있음이 보장된다
			// n - 1<<(i-1) + 1: n 이하이면서 최상위 비트가 n과 같은 수의 개수
			res += sum[i-1] + n - 1<<(i-1) + 1
			// n에서 최상위 비트 제거
			n -= 1 << (i - 1)
		}
	}

	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
