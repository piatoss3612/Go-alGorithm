package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Route struct {
	start, end, width int
	index             int
}

type Routes []*Route

func (r Routes) Len() int { return len(r) }
func (r Routes) Less(i, j int) bool {
	return r[i].width > r[j].width
}
func (r Routes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
	r[i].index, r[j].index = i, j
}
func (r *Routes) Push(x interface{}) {
	*r = append(*r, x.(*Route))
}
func (r *Routes) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	p, w, c, v int
	parent     []int
)

// 메모리: 5612KB
// 시간: 32ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	p, w = scanInt(), scanInt()
	c, v = scanInt(), scanInt()

	parent = make([]int, p)
	for i := 1; i < p; i++ {
		parent[i] = i
	}

	routes := &Routes{}
	heap.Init(routes)

	var start, end, width int
	for i := 1; i <= w; i++ {
		start, end, width = scanInt(), scanInt(), scanInt()
		heap.Push(routes, &Route{start, end, width, i})
	}

	for len(*routes) > 0 {
		r := heap.Pop(routes).(*Route)
		union(r.start, r.end)

		if find(c) == find(v) {
			fmt.Fprintln(writer, r.width)
			return
		}
	}
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
