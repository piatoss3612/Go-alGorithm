package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	list    []string
	t, n    int
)

// 메모리: 5524KB -> 5532KB
// 시간: 1604ms -> 76ms
// 문자열 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t = scanInt()
	for i := 1; i <= t; i++ {
		TestCase()
	}
}

func TestCase() {
	n = scanInt()
	list = make([]string, n)
	for i := 0; i < n; i++ {
		list[i] = scanString()
	}

	// 숫자 문자열 슬라이스 정렬
	sort.Strings(list)

	// O(n^2) -> O(n)으로 시간 복잡도 단축
	/*
		문자열 슬라이스의 오름차순 정렬 특성상

		정렬 전: [12 123 11 112 134]
		정렬 후: [11 112 12 123 134]

		공통으로 포함되어 있는 문자에 대하여 더 짧은 문자열이 앞으로 가고
		긴 문자열은 뒤로 가게 된다
		따라서 O(n) 시간에 i번째 문자열과 i+1번째 문자열만 비교함으로써 시간을 단축할 수 있다
	*/
	for i := 0; i < n-1; i++ {
		// 앞에 있는 번호의 길이가 더 짧은 경우만: 전화번호가 같은 경우는 없다
		if len(list[i]) < len(list[i+1]) {
			// 뒤에 있는 번호가 접두사로 앞에 있는 번호를 포함하는 경우
			if strings.HasPrefix(list[i+1], list[i]) {
				fmt.Fprintln(writer, "NO")
				return
			}
		}
	}

	fmt.Fprintln(writer, "YES")
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
