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
	inp     []int
)

// 메모리: 57568KB -> 44048KB
// 시간: 412ms -> 316ms
// 분할 정복과 병합 정렬을 사용해 버블 정렬에서 숫자의 위치가 스왑된 횟수를 구한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp = make([]int, n+1)

	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	cnt := MergeSort(1, n)
	fmt.Fprintln(writer, cnt)
}

// 분할 정복을 사용한 병합 정렬
func MergeSort(left, right int) int {
	// 기저 사례: 정렬할 슬라이스의 길이가 1인 경우
	if left == right {
		return 0
	}

	// 왼쪽, 오른쪽으로 분할하여 정렬
	mid := (left + right) / 2
	ret := MergeSort(left, mid) + MergeSort(mid+1, right) // 왼쪽 슬라이스 정렬 스왑 횟수 + 오른쪽 슬라이스 정렬 스왑 횟수

	mergeSorted := make([]int, right-left+1) // append를 사용해 메모리가 낭비되는 것을 방지하기 위해 left~right만큼의 길이로 할당
	idx := 0

	segL, segR := inp[left:mid+1], inp[mid+1:right+1]

	// left~right 전체 슬라이스 정렬
	for len(segL) > 0 || len(segR) > 0 {

		// 각각의 슬라이스의 맨 앞에 있는 수를 비교
		// 왼쪽이 오른쪽보다 큰 경우
		if len(segL) > 0 && len(segR) > 0 && segL[0] > segR[0] {
			ret += len(segL) // 오른쪽에 있는 수가 현재 왼쪽 슬라이스의 길이만큼 앞으로 이동
			mergeSorted[idx] = segR[0]
			segR = segR[1:]
			idx += 1
			continue
		}

		// 버블 정렬은 동일한 수의 위치가 뒤바뀌지 않는 안정된 정렬 방식이므로
		// 왼쪽과 오른쪽의 값이 같은 경우를 고려해 왼쪽에 있는 수를 먼저 정렬한다
		if len(segL) > 0 {
			mergeSorted[idx] = segL[0]
			segL = segL[1:]
			idx += 1
			continue
		}

		if len(segR) > 0 {
			mergeSorted[idx] = segR[0]
			segR = segR[1:]
			idx += 1
			continue
		}
	}

	copy(inp[left:right+1], mergeSorted) // 정렬된 슬라이스를 원래 슬라이스로 복사

	return ret // 누적 정렬 횟수 반환
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
