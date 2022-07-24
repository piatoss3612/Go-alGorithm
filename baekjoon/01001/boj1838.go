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
	inp     []Bubble
	N       int
)

type Bubble struct {
	beforeSort int // 정렬되기 이전의 인덱스
	value      int // 입력값
}

// 메모리: 13480KB
// 시간: 196ms
// 버블 정렬의 반복문이 종료되는 i를 찾는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]Bubble, N)
	for i := 0; i < N; i++ {
		inp[i] = Bubble{i, scanInt()}
	}

	// 버블 정렬: stable sort
	// 중복된 값의 순서가 변경되지 않는 정렬 방법
	sort.Slice(inp, func(i, j int) bool {
		if inp[i].value == inp[j].value {
			return inp[i].beforeSort < inp[j].beforeSort
		}
		return inp[i].value < inp[j].value
	})

	// 버블 정렬 과정에서 j, j+1 번째 값이 스왑됨으로 인해
	// 앞쪽으로 이동한 값의 이전 인덱스와 현재 인덱스의 차의 최댓값이
	// 곧 버블 정렬의 반복문이 종료되는 i이다
	ans := 0
	for i := 0; i < N; i++ {
		ans = max(ans, inp[i].beforeSort-i)
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
