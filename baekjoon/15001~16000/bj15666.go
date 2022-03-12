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
	seq     []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	input = make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	BackTracking(0, 0)
}

func BackTracking(length, prev int) {
	if length == m {
		printSeq()
		return
	}

	picked := 0
	for i := 0; i < n; i++ {
		if prev > input[i] || picked == input[i] {
			continue
		}
		picked = input[i]
		seq = append(seq, input[i])
		BackTracking(length+1, input[i])
		seq = seq[:len(seq)-1]
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
