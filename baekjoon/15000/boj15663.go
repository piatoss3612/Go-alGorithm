package main

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
	n, m    int
	input   []int
	visited []bool
	seq     []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	input = make([]int, n)
	visited = make([]bool, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	BackTracking(0)
}

/*
예제 입력:
4 2
9 7 9 1

프로세스:
9 7 9 1 -> 정렬 1 7 9 9

1: 1 7, 1 9, (1 9)
7: 7 1, 7 9, (7 9)
9: 9 1, 9 7, 9 9
9: (9 1), (9 7), (9 9)

예제 출력:
1 7
1 9
7 1
7 9
9 1
9 7
9 9
*/

func BackTracking(length int) {
	if length == m {
		printSeq()
		return
	}

	prev := 0
	for i := 0; i < n; i++ {
		if visited[i] || prev == input[i] {
			continue
		}
		visited[i] = true
		prev = input[i]
		seq = append(seq, input[i])
		BackTracking(length + 1)
		seq = seq[:len(seq)-1]
		visited[i] = false
	}
}

func printSeq() {
	for _, v := range seq {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
