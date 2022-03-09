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

	if n <= 10 {
		fmt.Fprintln(writer, n-1)
		return
	}

	queue := []int{}

	for i := 1; i <= 9; i++ {
		queue = append(queue, i)
	}

	cnt := 10

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

	if cnt < n {
		fmt.Fprintln(writer, -1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
