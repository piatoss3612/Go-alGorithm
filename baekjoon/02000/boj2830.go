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
	onebit  [20]int // 2^19 < 10^6 < 2^20, 2진수로 표현된 이름은 1백만 이하이므로 못해도 20개의 비트로 표현할 수 있다
	N       int
)

// 메모리: 5588KB
// 시간 240ms
// 비트마스크
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	// 외계인들의 이름을 2중 반복문으로 xor연산하여 누적값을 구해도 되지만
	// 이 경우는 입력값이 1백만 개 이므로 O(N^2)은 시간 초과가 발생할 수 있다

	// 따라서 O(N^2)보다 빠른 시간에 풀 수 있는 방법을 생각해 보았고
	// 오랫동안 고민한 끝에 O(N*logN)에 풀 수 있는 방법을 찾았다
	// 그것은 2진수의 i번째 비트에 포함된 모든 1의 개수를 구하고
	// xor 연산의 특성을 이용해 결과를 구하는 것이다

	for i := 1; i <= N; i++ {
		name := scanInt()
		// shift하여 곱하는 과정을 고려해 시작 인덱스는 0으로 설정하였다
		idx := 0
		// 이름의 idx번째 비트가 1이라면 onebit에 누적해서 더해준다
		// 이 과정은 O(logN)의 시간이 소요된다
		for name > 0 {
			// 전체 N개의 이름의 idx번째 비트에 포함된 1의 개수를 알고 있다면
			// 0의 개수는 N-(1의 개수)로 구할 수 있으므로 비트값이 1인 경우만 고려
			if name%2 != 0 {
				onebit[idx]++
			}
			name /= 2
			idx++
		}
	}

	ans := 0

	for i := 0; i < 20; i++ {
		// i번째 비트에 포함된 1의 개수가 0이라면
		// i번째 비트가 모두 0이거나 i번째 비트가 포함된 이름이 없는 경우이다
		if onebit[i] > 0 {
			// i번째 비트에 포함된 1의 개수 * 0의 개수 * i번째 비트에 해당하는 2의 제곱수
			ans += onebit[i] * (N - onebit[i]) * (1 << i)

			/*
				xor의 특성:
				A xor B가 1인 경우는 A=1,B=0 또는 A=0,B=1
				xor연산으로 1이 나오는 경우는 1과 0의 조합의 개수와 동일하다

				따라서 (i번째 비트에 포함된 1의 개수 * 0의 개수)는
				2중 반복문을 통해 xor연산을 하는 것과 동일한 결과를 도출하면서도 더 빠른 방법이다
			*/
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
