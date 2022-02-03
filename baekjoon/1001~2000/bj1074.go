package bj1074

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
	n, r, c := scanInt(), scanInt(), scanInt()

	result := 0

	/*
		사분면의 크기를 줄여가면서
		사분면의 가장 기본 형태인
		0 1
		2 3
		으로 치환해서 문제를 풀 수 있다
	*/

	for n > 0 {

		n -= 1
		square := int(math.Pow(2, float64(n)))

		// #1 사분면
		if r < square && c < square {
			result += square * square * 0
		}

		// #2 사분면
		if r < square && c >= square {
			result += square * square * 1
			c -= square
		}

		// #3 사분면
		if r >= square && c < square {
			result += square * square * 2
			r -= square
		}

		// #4 사분면
		if r >= square && c >= square {
			result += square * square * 3
			r -= square
			c -= square
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
