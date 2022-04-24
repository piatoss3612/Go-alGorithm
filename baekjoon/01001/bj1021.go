package bj1021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Deque []int

func (d *Deque) pop() {
	*d = (*d)[1:]
}

func (d *Deque) rotate_left() {
	*d = append((*d)[1:], (*d)[0])
}

func (d *Deque) rotate_right() {
	*d = append((*d)[len(*d)-1:], (*d)[0:len(*d)-1]...)
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	var deque Deque = make([]int, n)
	// 꺼내야할 순서를 저장
	for i := 1; i <= m; i++ {
		deque[scanInt()-1] = i
	}
	cnt := 0  // 큐 연산 횟수
	turn := 1 // 현재 찾으려는 수
	for {
		if turn > m {
			break
		}

		var idx int
		for i := 0; i < len(deque); i++ {
			if deque[i] == turn {
				idx = i
				break
			}
		}
		// 큐의 길이를 절반보다 작거나 같은 위치에 있으면 왼쪽으로 로테이션
		if idx <= len(deque)/2 {
			for i := 0; i < idx; i++ {
				deque.rotate_left()
				cnt += 1
			}
			deque.pop()
			turn += 1
		} else { // 오른쪽으로 로테이션
			for i := 0; i < len(deque)-idx; i++ {
				deque.rotate_right()
				cnt += 1
			}
			deque.pop()
			turn += 1
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
