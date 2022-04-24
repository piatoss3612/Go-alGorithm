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
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	ans := 0
	// 모든 입력값에 대해 두 포인터를 사용해 해당 입력값을 만들 수 있는지 확인
	// s 또는 e가 i와 동일한 경우는 다른 두 수가 아니므로 포함하지 않는다
	for i := 0; i < n; i++ {
		s, e, sum := 0, n-1, 0
		for s < e {
			sum = input[s] + input[e]
			if sum == input[i] {
				if s != i && e != i {
					ans += 1
					break
				} else if s == i {
					s += 1
				} else if e == i {
					e -= 1
				}
			} else if sum < input[i] {
				s += 1
			} else {
				e -= 1
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
