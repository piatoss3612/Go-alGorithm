package bj1092

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
	crane := make([]int, n)
	for i := 0; i < n; i++ {
		crane[i] = scanInt()
	}

	m := scanInt()
	box := make([]int, m)
	for i := 0; i < m; i++ {
		box[i] = scanInt()
	}

	sort.Ints(crane)
	sort.Ints(box)

	if crane[n-1] < box[m-1] {
		fmt.Fprintln(writer, -1)
		return
	}

	/*
		오름차순으로 정렬된 크레인을 기준으로
		가장 큰 값의 크레인 ~ 가장 작은 값의 크레인까지 순회
		오름차순으로 정렬된 박스값의 뒤쪽에서부터
		해당하는 크레인보다 작거나 같은 경우
		해당 박스값을 슬라이스에서 제거해준다

		한 번 순회할 때마다 ans 값에 1을 더해주며
		박스 슬라이스에 값이 존재하지 않을 경우에 종료된다
	*/

	ans := 0
	for len(box) > 0 {
		ans += 1
		for i := n - 1; i >= 0; i-- {
			for j := len(box) - 1; j >= 0; j-- {
				if crane[i] >= box[j] {
					box = append(box[:j], box[j+1:]...)
					break
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
