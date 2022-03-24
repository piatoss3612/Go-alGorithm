package bj1715

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	card := &Heap{}
	heap.Init(card)
	for i := 0; i < n; i++ {
		heap.Push(card, scanInt())
	}

	/*
		최솟값 2개를 뽑아서 더한 값을 슬라이스에 추가하고
		이를 슬라이스의 길이가 1이 될 때까지 반복하는 그리디 알고리즘

		단순히 슬라이스를 사용하면 시간초과가 발생하므로
		힙을 사용하여 시간 복잡도를 낮춰야 한다
	*/

	var a, b, tmp int
	ans := 0
	for card.Len() > 1 {
		a, b = heap.Pop(card).(int), heap.Pop(card).(int)
		tmp = a + b
		ans += tmp
		heap.Push(card, tmp)
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
