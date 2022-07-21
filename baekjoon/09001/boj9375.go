package bj9375

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
	t := scanInt()
	for i := 0; i < t; i++ {
		testFashion()
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func testFashion() {
	n := scanInt()
	styles := make(map[string]int)
	for i := 0; i < n; i++ {
		// 의상 종류에 해당하는 의상 이름은 중복되지 않으므로 개수만 카운트하여 맵에 저장
		_, part := scanString(), scanString()
		_, ok := styles[part]
		if ok {
			styles[part] += 1
		} else {
			styles[part] = 1
		}
	}
	result := 1
	for _, v := range styles {
		// 의상 종류의 개수 + 입지 않는 경우의 수를 곱해준다
		result *= v + 1
	}
	// 결과에서 모두 입지 않는 경우의 수 1을 빼준다
	fmt.Fprintln(writer, result-1)
}
