package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Meeting struct {
	start, end int
}

type PQ []Meeting

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].end < pq[j].end }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Meeting))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	m := make([]Meeting, n)
	for i := 0; i < n; i++ {
		m[i] = Meeting{scanInt(), scanInt()}
	}

	sort.Slice(m, func(i, j int) bool {
		if m[i].start == m[j].start {
			return m[i].end < m[j].end
		}
		return m[i].start < m[j].start
	})

	t := &PQ{}
	heap.Init(t)

	for i := 0; i < n; i++ {
		if t.Len() == 0 {
			heap.Push(t, m[i])
			continue
		}

		tmp := heap.Pop(t).(Meeting)

		if tmp.end <= m[i].start {
			heap.Push(t, m[i])
		} else {
			heap.Push(t, tmp)
			heap.Push(t, m[i])
		}
	}

	fmt.Fprintln(writer, t.Len())
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
