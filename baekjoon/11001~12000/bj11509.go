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
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = scanInt()
	}

	// 1차 풀이: 메모리 11692KB 584ms
	// m := []int{}

	// for i := 0; i < n; i++ {
	// 	if len(m) == 0 {
	// 		m = append(m, h[i]-1)
	// 		continue
	// 	}
	// 	flag := true
	// 	for idx, v := range m {
	// 		if v == h[i] {
	// 			m[idx] -= 1
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		m = append(m, h[i]-1)
	// 	}
	// }

	// fmt.Fprintln(writer, len(m))

	// 2차 풀이: 메모리 11848KB 120ms
	check := make([]int, 1000001)

	ans := 0
	for i := 0; i < n; i++ {
		if check[h[i]] == 0 {
			ans += 1
			check[h[i]-1] += 1
		} else {
			check[h[i]] -= 1
			check[h[i]-1] += 1
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
