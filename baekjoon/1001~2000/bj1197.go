package bj1043

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type link struct {
	a, b, c int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	v, e    int
	links   []link
	parent  []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	v, e = scanInt(), scanInt()

	links = make([]link, e)
	parent = make([]int, v+1)
	for i := 0; i < e; i++ {
		links[i] = link{scanInt(), scanInt(), scanInt()}
	}

	for i := 1; i <= v; i++ {
		parent[i] = i
	}

	sort.Slice(links, func(i, j int) bool {
		return links[i].c < links[j].c
	})

	var ans int
	numberOfLinks := 0

	for _, link := range links {
		if numberOfLinks == v-1 {
			break
		}
		if !sameParent(link.a, link.b) {
			union(link.a, link.b)
			ans += link.c
			numberOfLinks += 1
		}
	}
	fmt.Fprintf(writer, "%d\n", ans)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		parent[y] = x
	}
}

func sameParent(x, y int) bool {
	x = find(x)
	y = find(y)
	if x == y {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
