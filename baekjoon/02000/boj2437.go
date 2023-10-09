package main

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
	weight  []int // 각 저울추의 무게
	N       int   // 저울추의 개수
)

// 메모리: 916KB
// 시간: 4ms
// 그리디 알고리즘
// 도저히 안풀려서 질문을 참고했다 https://www.acmicpc.net/board/view/45841
// 범골도 노력하면 천재의 발치라도 닿을 수 있을까?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	weight = make([]int, N)
	for i := 0; i < N; i++ {
		weight[i] = scanInt()
	}

	sort.Ints(weight) // 저울추를 가벼운 것부터 오름차순으로 정렬

	measurable := 0

	// measurable은 [0, measurable] 연속된 범위의 최댓값으로
	for i := 0; i < N; i++ {
		// 오름차순으로 정렬한 i번째 저울추의 무게 weight[i]가 (측정가능한 범위의 최댓값+1)보다 작거나 같다면
		if measurable+1 >= weight[i] {
			// i번째 저울추의 무게를 포함하여 [0,measurable+weight[i]] 연속된 구간의 자연수들을
			//  0~i번째 저울추를 조합해 표현할 수 있다는 것을 귀납적으로 추론할 수 있다
			measurable += weight[i]
		} else {
			// i번째 저울추의 무게가  [0, measurable] 연속된 범위에 포함되지 않는다면
			// measurable+1은 표현할 수 없는 최솟값이 된다
			break
		}
	}

	fmt.Fprintln(writer, measurable+1)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
