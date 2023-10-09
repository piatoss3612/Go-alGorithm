package bj11652

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
	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		num := scanInt()
		_, ok := counts[num]
		if ok {
			counts[num] += 1
		} else {
			counts[num] = 1
		}
	}
	key := 0
	max := 0
	for k, v := range counts {
		if v > max {
			max = v
			key = k
		}
		// 카운트가 같은 경우 키 값이 작은 값으로 변경
		if v == max {
			if k < key {
				key = k
			}
		}
	}
	fmt.Fprintln(writer, key)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
