package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	Q       int                  // 쿼리의 개수
	gorilla map[string]*MaxValue // 맵의 [키]:값으로 [고릴라의 이름]:최대 힙에 저장된 고릴라가 가지고 있는 정보의 가치를 저장
	ans     = 0                  // 호석이가 얻게 되는 정보 가치의 총합
)

// 최대 힙 타입 정의 및 인터페이스 구현
type MaxValue []int

func (m MaxValue) Len() int { return len(m) }
func (m MaxValue) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m MaxValue) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m *MaxValue) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MaxValue) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// 메모리: 9496KB
// 시간: 296ms
// 맵, 최대 힙
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Q = scanInt()
	gorilla = make(map[string]*MaxValue)

	for i := 1; i <= Q; i++ {
		query := scanInt() // 쿼리 번호 입력
		switch query {
		case 1:
			// 쿼리 번호가 1인 경우, 고릴라가 새로운 정보를 얻는 이벤트 발생
			gorillaGotInformation()
		case 2:
			// 쿼리 번호가 2인 경우, 호석이가 고릴라에게 정보를 구매하는 이벤트 발생
			transaction()
		}
	}

	fmt.Fprintln(writer, ans)
}

func gorillaGotInformation() {
	name, k := scanName(), scanInt() // 정보를 얻는 고릴라의 이름과 얻게 되는 정보의 수 입력
	g, ok := gorilla[name]           // 고릴라가 새로운 정보상인지 확인

	if ok {
		// 새로운 정보상이 아닌 경우
		for i := 1; i <= k; i++ {
			heap.Push(g, scanInt())
		}
	} else {
		// 새로운 정보상인 경우
		h := &MaxValue{}
		for i := 1; i <= k; i++ {
			heap.Push(h, scanInt())
		}
		gorilla[name] = h
	}
}

func transaction() {
	name, number := scanName(), scanInt() // 호석이와 거래할 고릴라와 거래할 정보의 수 입력
	info, ok := gorilla[name]             // 호석이와 거래할 고릴라가 정보를 가지고 있는지 확인

	if !ok {
		return // 정보를 가지고 있지 않다면 정보를 구매하지 않고 거래 종료
	}

	// 정보를 가지고 있다면 가지고 있는 한에서 number 개수 만큼 정보를 구매
	for info.Len() > 0 && number > 0 {
		ans += heap.Pop(info).(int)
		number--
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanName() string {
	scanner.Scan()
	return scanner.Text()
}
