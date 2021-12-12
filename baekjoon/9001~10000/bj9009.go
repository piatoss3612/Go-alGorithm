package bj9009

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
	fib     = []int{1, 1}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for i := 2; ; i++ {
		tmp := fib[i-1] + fib[i-2]
		if tmp > 1000000000 {
			break
		} else {
			fib = append(fib, tmp)
		}
	}

	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		tmp := n
		result := []int{}
		for tmp != 0 {
			for j := len(fib) - 1; j > 0; j-- {
				if tmp-fib[j] >= 0 {
					result = append(result, fib[j])
					tmp -= fib[j]
				}
			}
		}
		sort.Ints(result)
		for _, v := range result {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
