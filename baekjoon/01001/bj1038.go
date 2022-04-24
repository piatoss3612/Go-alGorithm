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

	if n <= 9 {
		fmt.Fprintln(writer, n)
		return
	}

	queue := []int{}

	for i := 1; i <= 9; i++ {
		queue = append(queue, i)
	}

	cnt := 9

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		for i := 0; i < front%10; i++ {
			tmp := front*10 + i
			cnt += 1
			if cnt == n {
				fmt.Fprintln(writer, tmp)
				return
			}
			queue = append(queue, tmp)
		}
	}

	// 줄어드는 수를 저장하는 슬라이스는 특정 시점,
	// 즉 9876543210이 구해지는 순간 멈추게 될 것이므로
	// n번째 줄어드는 수를 구할 수 없는 경우는 반복문 밖에서 구할 수 있다

	if cnt < n {
		fmt.Fprintln(writer, -1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
