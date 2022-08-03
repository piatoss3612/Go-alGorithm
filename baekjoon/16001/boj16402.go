package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	Kingdoms map[string]string
	N, M     int
)

// 메모리: 1324KB
// 시간: 8ms
// 분리 집합, 부모 요소가 누구인지만 알 수 있으면 해당 왕국이 속국인지 아닌지 판별할 수 있다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	s := strings.Split(scanString(), " ")
	N, _ = strconv.Atoi(s[0])
	M, _ = strconv.Atoi(s[1])
	Kingdoms = make(map[string]string)

	for i := 1; i <= N; i++ {
		k := scanString()
		Kingdoms[k] = k
	}

	for i := 1; i <= M; i++ {
		q := strings.Split(scanString(), ",")
		n, _ := strconv.Atoi(q[2])
		x, y := Find(q[0]), Find(q[1])

		switch n {
		case 1:
			// 첫 번째 왕국의 종주국과 두 번째 왕국의 종주국이 같은 경우는
			// 속국이 종주국을 공격한 경우
			if x == y {
				Kingdoms[q[0]] = q[0]
				Kingdoms[y] = q[0]
			} else {
				Kingdoms[y] = x
			}
		case 2:
			if x == y {
				Kingdoms[q[1]] = q[1]
				Kingdoms[x] = q[1]
			} else {
				Kingdoms[x] = y
			}
		}
	}

	var res []string

	for k, v := range Kingdoms {
		if k == Find(v) {
			res = append(res, k)
		}
	}

	sort.Strings(res) // 사전순으로 정렬

	fmt.Fprintln(writer, len(res))

	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func Find(x string) string {
	if Kingdoms[x] == x {
		return x
	}

	Kingdoms[x] = Find(Kingdoms[x])
	return Kingdoms[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
