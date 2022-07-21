package bj18870

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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, n) // 정렬할 입력값
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	tmp := make([]int, n)      // 정렬하기 전의 입력값을 저장할 슬라이스
	copy(tmp, input)           // 정렬하기 전의 입력값을 복사
	sort.Ints(input)           // 입력값 오름차 순으로 정렬
	order := make(map[int]int) // 압축할 좌표를 저장할 맵
	cnt := 0
	order[input[0]] = cnt //  최솟값은 0으로 좌표 압축
	for i := 1; i < n; i++ {
		if input[i] == input[i-1] { // 이전 값하고 같다면 좌표를 추가하지 않는다
			continue
		}
		cnt += 1 // 이전 값과 다르다면 좌표를 1늘리고 압축된 좌표를 맵에 저장
		order[input[i]] = cnt
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", order[tmp[i]]) // 해당 좌표에 해당하는 압축 좌표를 출력
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
