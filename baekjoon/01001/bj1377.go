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

// 입력된 값과 정렬 전의 인덱스를 저장하는 구조체
type sorting struct {
	value int
	index int
}

// 메모리: 13500KB
// 시간: 248ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	sorted := make([]sorting, n) // 입력된 값과 정렬 전의 인덱스를 저장하는 sorting 구조체의 슬라이스 생성
	for i := 0; i < n; i++ {
		sorted[i] = sorting{scanInt(), i}
	}

	// 1. 값만 비교: 틀렸습니다
	// sort.Slice(sorted, func(i, j int) bool {
	// 	return sorted[i].value < sorted[j].value
	// })

	// 2. 값을 비교한느데 값이 같으면 인덱스를 비교: 맞았습니다
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].value == sorted[j].value {
			return sorted[i].index < sorted[j].index
		}
		return sorted[i].value < sorted[j].value
	})

	/*
		# 정렬 방법 1과 2의 차이

		버블 정렬은 값이 같으면 인덱스의 순서가 뒤바뀌지 않는 stable sort
		따라서 값이 같은 경우 인덱스가 뒤바뀌지 않는 다는 것을 보장해야 한다
	*/

	ans := 0

	for i, v := range sorted {
		// 정렬된 슬라이스의 i번째 값이 기존의 인덱스로부터 얼마나 앞으로 이동했는지의 최댓값을 구한다
		if ans < v.index-i {
			ans = v.index - i
		}
	}

	/*
		# 버블 정렬이 종료되는 i

		= 정렬된 슬라이스의 i번째 값이 정렬되지 않은 기존의 인덱스로부터 앞으로 이동한 차이의 최댓값?

		2 3 4 1 을 버블 정렬

		i = 1일 때, 2 3 1 4
		i = 2일 때, 2 1 3 4
		i = 3일 때, 1 2 3 4
		i = 4일 때, 버블 정렬 종료

		버블 정렬의 특성상 뒤에 있던 값이 앞으로 이동하는 경우는 한 번의 순회에서 한 번 밖에 발생하지 않는다
		따라서 기존의 인덱스 - 현재 인덱스의 값이 커질 수록 버블 정렬을 위해 슬라이스를 여러 번 순회하고 있다는 것을 의미한다


		# 반대로 앞에서 뒤로 이동하는 경우

		10 1 2 3 을 버블 정렬

		i = 1일 때, 1 2 3 10
		i = 2일 때, 버블 정렬 종료

		앞에 있던 값이 뒤로 이동하는 경우는 한 번의 순회에서 여러 번 발생할 수 있으므로
		버블 정렬이 종료되는 i가 될 수 없다
	*/

	fmt.Fprintln(writer, ans+1) // 문제에서 인덱스는 1부터 시작하므로 1을 더해준 값을 출력
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
