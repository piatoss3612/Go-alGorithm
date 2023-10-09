package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	K     int
	signs []byte
)

// 난이도: Silver 1
// 메모리: 916KB
// 시간: 8ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	K = scanInt()
	signs = make([]byte, K)
	for i := 0; i < K; i++ {
		signs[i] = scanString()[0]
	}
}

func Solve() {
	fmt.Fprintln(writer, maxNum())
	fmt.Fprintln(writer, minNum())
}

func maxNum() string {
	res := make([]int, 0, K+1)
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < K; i++ {
		switch signs[i] {
		case '<':
			cnt := 1 // 연속된 '<'의 개수
			for j := i + 1; j < K; j++ {
				if signs[j] == '<' {
					cnt++
				} else {
					break
				}
			}
			res = append(res, nums[len(nums)-cnt-1]) // 남은 수 중 cnt번째로 큰 수
			nums = append(nums[:len(nums)-cnt-1], nums[len(nums)-cnt:]...)
		case '>':
			res = append(res, nums[len(nums)-1]) // 남은 수 중 가장 큰 수
			nums = nums[:len(nums)-1]
		}
	}

	res = append(res, nums[len(nums)-1]) // 남은 수 중 가장 큰 수

	builder := strings.Builder{}
	for _, n := range res {
		builder.WriteString(strconv.Itoa(n))
	}

	return builder.String()
}

func minNum() string {
	res := make([]int, 0, K+1)
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < K; i++ {
		switch signs[i] {
		case '<':
			res = append(res, nums[len(nums)-1]) // 남은 수 중 가장 작은 수
			nums = nums[:len(nums)-1]
		case '>':
			cnt := 1 // 연속된 '>'의 개수
			for j := i + 1; j < K; j++ {
				if signs[j] == '>' {
					cnt++
				} else {
					break
				}
			}
			res = append(res, nums[len(nums)-cnt-1]) // 남은 수 중 cnt번째로 작은 수
			nums = append(nums[:len(nums)-cnt-1], nums[len(nums)-cnt:]...)
		}
	}

	res = append(res, nums[len(nums)-1]) // 남은 수 중 가장 작은 수

	builder := strings.Builder{}
	for _, n := range res {
		builder.WriteString(strconv.Itoa(n))
	}

	return builder.String()
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
