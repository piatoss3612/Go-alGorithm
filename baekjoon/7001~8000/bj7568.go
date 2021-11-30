package bj7568

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

type big struct {
	weight int
	height int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	bs := make([][2]int, 0, n)

	for i := 0; i < n; i++ {
		w, h := scanInt(), scanInt()
		bs = append(bs, [2]int{w, h})
	}

	ranks := make([]int, n)

	for i := 0; i < n; i++ {
		ranks[i] = 1
		for j := 0; j < n; j++ {
			if (bs[i][0] < bs[j][0]) && (bs[i][1] < bs[j][1]) {
				ranks[i] += 1
			}
		}
	}

	for _, v := range ranks {
		fmt.Fprint(writer, v, " ")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
