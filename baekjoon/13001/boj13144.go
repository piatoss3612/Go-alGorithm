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
	count   [100001]int
	inp     []int
	N       int
)

// 메모리: 3872KB
// 시간: 20ms
// 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]int, N+1)
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}

	s, e := 1, 1 // 첫 번째 인덱스에서 시작

	sum := 0

	for s <= e && e <= N {
		// 1. 중복된 숫자가 등장하지 않은 경우
		// e를 앞으로 한 칸 이동하고
		// e번째 수를 포함하여 새롭게 만들 수 있는 부분 수열의 수를 더해준다
		if count[inp[e]] == 0 {
			count[inp[e]]++
			e++
			sum += e - s
		} else {
			// 2. 중복된 숫자가 등장한 경우
			// s를 앞으로 한 칸 이동
			count[inp[s]]--
			s++
		}
	}

	fmt.Fprintln(writer, sum)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
