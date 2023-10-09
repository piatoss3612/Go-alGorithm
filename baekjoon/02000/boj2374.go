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

// 메모리: 888KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	ans := 0
	var s []int // 스택

	// 값을 입력받으면서 스택의 값들이 오름차순인 경우를 처리
	for i := 0; i < n; i++ {
		tmp := scanInt()

		// 스택의 길이가 0이거나 입력된 값이 스택의 마지막 값보다 작거나 같은 경우
		if len(s) == 0 || s[len(s)-1] >= tmp {
			s = append(s, tmp)
			continue
		}

		// 입력된 값의 스택의 마지막 값보다 큰 경우
		if s[len(s)-1] < tmp {
			ans += tmp - s[len(s)-1] // 입력된 값 - 스택의 마지막 값 만큼 Add 연산
			// 입력된 값보다 큰 값이 나올 때까지 스택에서 Pop 연산
			for len(s) != 0 && s[len(s)-1] < tmp {
				s = s[:len(s)-1]
			}
			s = append(s, tmp) // 입력된 값 스택에 추가
		}
	}

	// 스택의 값이 내림차순(같은 값 포함)인 경우를 처리
	for i := len(s) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			ans += s[i-1] - s[i]
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
