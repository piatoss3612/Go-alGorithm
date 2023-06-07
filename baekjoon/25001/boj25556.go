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
	N       int
	arr     []int
	s       [4][]int
)

// 난이도: Gold 5
// 메모리: 5400KB
// 시간: 24ms
// 분류: 그리디 알고리즘, 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	// 입력받은 순열의 모든 수를 4개의 스택에 오름차순으로 집어넣을 수 있으면 순열을 청소할 수 있다는 뜻
	for i := 0; i < N; i++ {
		flag := false
		for j := 0; j < 4; j++ {
			if len(s[j]) == 0 || s[j][len(s[j])-1] < arr[i] {
				s[j] = append(s[j], arr[i])
				flag = true
				break
			}
		}

		if !flag {
			fmt.Fprintln(writer, "NO")
			return
		}
	}

	fmt.Fprintln(writer, "YES")
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
