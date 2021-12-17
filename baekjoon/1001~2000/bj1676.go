package bj1676

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
	facTest(n)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func facTest(n int) {
	// 팩토리얼을 직접 구해서 계산하면 너무 큰 수가 나올 수 있으므로
	// 곱해지는 모든 항의 소인수가 2이거나 5인 경우를 카운트
	result := map[int]int{2: 0, 5: 0}
	for i := n; i >= 2; i-- {
		tmp := i
		for {
			if tmp%5 == 0 {
				tmp /= 5
				result[5] += 1
			} else if tmp%2 == 0 {
				tmp /= 2
				result[2] += 1
			} else {
				break
			}
		}
	}
	// 2와 5 중 작은 수만큼 10을 만들 수 있으므로 작은 수를 출력
	if result[2] >= 1 && result[5] >= 1 {
		if result[2] > result[5] {
			fmt.Fprintln(writer, result[5])
		} else {
			fmt.Fprintln(writer, result[2])
		}
	} else {
		fmt.Fprintln(writer, 0)
	}
}
