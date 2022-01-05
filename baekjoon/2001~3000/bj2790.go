package bj2790

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
	// 내림차순 정렬
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})
	result := 0 // 우승할 가능성이 있는 사람의 수
	tmp := 0    // 이것만 넘으면 우승이 가능한 수
	for i := 0; i < n; i++ {
		// 최고점은 받았는데 우승 하기 위해 필요한 수 보다 큰 경우
		if input[i]+n >= tmp {
			result += 1
		}
		// 이것만 넘으면 우승 가능한 수를 갱신
		if input[i]+i+1 > tmp {
			tmp = input[i] + i + 1
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
