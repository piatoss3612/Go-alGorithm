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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	bookLeft := []int{}  // 0보다 작은 값을 저장
	bookRight := []int{} // 0이상인 값을 저장

	for i := 0; i < n; i++ {
		tmp := scanInt()
		if tmp < 0 {
			bookLeft = append(bookLeft, -tmp) // 0보다 작은 경우 계산이 용이하도록 양수로 바꿔준다
		} else {
			bookRight = append(bookRight, tmp)
		}
	}

	// 오름차순으로 정렬
	sort.Ints(bookLeft)
	sort.Ints(bookRight)

	ans := 0
	bll := len(bookLeft)
	brl := len(bookRight)
	//  0에서부터 가장 큰 값(m개의 책을 한 번에 들고)으로 왕복하는 거리를 누적해서 더해준다
	for i := bll - 1; i >= 0; i -= m {
		ans += bookLeft[i] * 2
	}

	for i := brl - 1; i >= 0; i -= m {
		ans += bookRight[i] * 2
	}

	// 절댓값이 가장 큰 값은 가장 마지막에 움직여서 0으로 돌아오지 않는 것이 최적의 경우이므로
	// 누적된 값에서 한 번 빼준다
	if bll == 0 {
		ans -= bookRight[brl-1]
	} else if brl == 0 {
		ans -= bookLeft[bll-1]
	} else {
		ans -= getMax(bookLeft[bll-1], bookRight[brl-1])
	}

	fmt.Fprintln(writer, ans)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
