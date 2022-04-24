package bj2960

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
	n, k := scanInt(), scanInt()
	checked := make([]int, 1001)
	cnt := 0
	for i := 2; i <= n; i++ {
		if checked[i] == 0 {
			checked[i] = 1
			cnt += 1
		}
		if cnt == k {
			fmt.Fprintln(writer, i)
			return
		}
		for j := i + i; j <= n; j += i {
			if checked[j] == 0 {
				checked[j] = 1
				cnt += 1
			}
			if cnt == k {
				fmt.Fprintln(writer, j)
				return
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
