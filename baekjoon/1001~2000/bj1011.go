package bj1011

import (
	"bufio"
	"fmt"
	"math"
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

	for i := 1; i <= t; i++ {
		testCase()
	}
}

/*
이동해야 할 거리: y - x = tmp
처음 이동하는 거리와 마지막에 이동하는 거리는 반드시 1

tmp = 1인 경우: 1 : 1번이동
tmp = 2인 경우: 1 1 : 2번이동
tmp = 3인 경우: 1 1 1 : 3번이동
tmp = 4인 경우: 1 2 1 : 3번이동
tmp = 5인 경우: 1 2 1 1 : 4번이동
tmp = 6인 경우: 1 2 2 1 : 4번이동
tmp = 7인 경우: 1 2 2 1 1 : 5번이동
tmp = 8인 경우: 1 2 2 2 1 : 5번이동
tmp = 9인 경우: 1 2 3 2 1 : 5번이동
...

tmp의 제곱근 n을 반올림 했을 때의 결과 출력:
1) n + 1이 되는 경우: (n + 1) * 2 - 1
2) n이 되는 경우: n * 2
*/

func testCase() {
	x, y := scanInt(), scanInt()
	tmp := y - x
	n := int(math.Round(math.Sqrt(float64(tmp))))

	if tmp > n*n {
		fmt.Fprintln(writer, n*2)
	} else {
		fmt.Fprintln(writer, n*2-1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
