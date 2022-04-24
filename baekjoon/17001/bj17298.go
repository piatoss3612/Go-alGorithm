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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	s := []int{}
	ans := []int{}

	for len(a) > 0 {
		// 입력된 값을 pop하여 tmp에 임시 저장
		tmp := a[len(a)-1]
		a = a[:len(a)-1]

		for len(s) > 0 {
			// s의 마지막 인덱스(top)에 있는 값이 tmp보다 큰 경우
			// ans에 s의 top을 push
			// s에는 tmp를 push
			if s[len(s)-1] > tmp {
				ans = append(ans, s[len(s)-1])
				s = append(s, tmp)
				break
			} else {
				// s의 마지막 인덱스(top)에 있는 값이 tmp보다 작은 경우
				// s에서 pop
				s = s[:len(s)-1]
			}
		}

		// 스택 길이가 0라면 스택에 tmp를 push
		// ans에는 -1을 결과로 push
		if len(s) == 0 {
			s = append(s, tmp)
			ans = append(ans, -1)
			continue
		}
	}

	// ans를 역순으로 출력
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", ans[i])
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
