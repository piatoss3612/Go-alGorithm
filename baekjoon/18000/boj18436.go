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
	tree    [100001]int
	inp     [100001]int
	N, M    int
)

// 메모리: 6304KB
// 시간: 80ms
// 세그먼트 트리로 푸니까 시간 초과가 나서 펜윅 트리로 풀었다. 아직 펜윅 트리의 로직을 제대로 이해한 것은 아니다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		inp[i] = scanInt() % 2 // 입력값을 2로 나눈 나머지를 저장: 홀수는 1, 짝수는 0
		if inp[i] == 1 {
			Update(i, 1) // 입력값이 1이라면 펜윅 트리 업데이트
		}
	}

	M = scanInt()

	for i := 1; i <= M; i++ {
		q, a, b := scanInt(), scanInt(), scanInt()
		switch q {
		case 1:
			temp := b % 2
			// 새로 변경할 값이 기존의 값과 다른 경우
			if temp != inp[a] {
				// inp[a]가 1이라면 0이 되므로 inp[a]가 포함된 구간에 -1
				if inp[a] == 1 {
					Update(a, -1)
				} else {
					Update(a, 1)
				}
				inp[a] = temp
			}
		case 2:
			odds := Sum(b) - Sum(a-1)        // 구간 a~b에 포함된 홀수의 개수
			fmt.Fprintln(writer, b-a+1-odds) // 짝수의 개수는 전체 수의 개수 - 홀수의 개수
		case 3:
			odds := Sum(b) - Sum(a-1)
			fmt.Fprintln(writer, odds)
		}
	}
}

// 펙윅 트리 구간합
func Sum(pos int) int {
	ret := 0
	for pos > 0 {
		ret += tree[pos]
		pos &= (pos - 1) // 최하위비트 제거
	}
	return ret
}

// 펜윅 트리 업데이트
func Update(pos, val int) {
	for pos <= N {
		tree[pos] += val
		pos += (pos & -pos) // 최하위비트 만큼 이동
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
