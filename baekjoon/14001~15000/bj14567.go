package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Subject struct {
	hasParent bool
	children  []int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	subs    []Subject
	ans     []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()

	subs = make([]Subject, n+1)
	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		subs[b].hasParent = true
		subs[a].children = append(subs[a].children, b)
	}
	ans = make([]int, n+1)

	for i := 1; i <= n; i++ {
		if !subs[i].hasParent {
			search(i, 1)
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%d ", ans[i])
	}
	fmt.Fprintln(writer)
}

func search(idx, cnt int) {
	ans[idx] = cnt

	for _, v := range subs[idx].children {
		if ans[v] <= cnt {
			search(v, cnt+1)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
