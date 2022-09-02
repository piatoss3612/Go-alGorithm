package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	N, M        int
	family      []string            // 여러 가문의 사람들
	inDegree    map[string]int      // 조상 -> 후손 진입 차수
	descendants map[string][]string // 후손들
	children    map[string][]string // 직계 자손들
	ancestors   []string            // 가문의 시조들
)

// 메모리: 28080KB
// 시간: 324ms
// 위상 정렬, 해시를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	family = make([]string, N)
	inDegree = make(map[string]int)
	descendants = make(map[string][]string)
	children = make(map[string][]string)

	for i := 0; i < N; i++ {
		name := scanName()
		family[i] = name
		inDegree[name] = 0
		descendants[name] = []string{}
	}

	M = scanInt()
	for i := 0; i < M; i++ {
		x, y := scanName(), scanName()
		inDegree[x]++                              // 후손의 진입 차수 증가
		descendants[y] = append(descendants[y], x) // 조상과 후손 관계 저장
	}

	queue := []string{} // 일반 큐를 사용

	for k, v := range inDegree {
		// 진입 차수가 0인 경우는 가문의 시조
		if v == 0 {
			ancestors = append(ancestors, k)
			queue = append(queue, k) // 큐에 가문의 시조들을 추가
		}
	}

	// 큐의 길이가 0이될 때까지 반복
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]

		// 큐에서 꺼내온 사람의 자손들을 탐색
		for _, child := range descendants[name] {
			inDegree[child]-- // 자손의 진입 차수 감소
			// 자손의 진입 차수가 0이라면 직계 자손이라는 의미
			if inDegree[child] == 0 {
				queue = append(queue, child)                   // 진입 차수가 0이된 사람의 이름을 큐에 추가
				children[name] = append(children[name], child) // 직계 자손 추가
			}
		}
	}

	// 가문의 개수와 각 가문의 시조들의 이름을 출력
	sort.Strings(ancestors)              // 이름은 사전순으로 출력해야 하므로 정렬을 꼭 해줘야 한다
	fmt.Fprintln(writer, len(ancestors)) // 가문의 수 = 시조들의 수
	for _, name := range ancestors {
		fmt.Fprintf(writer, "%s ", name)
	}
	fmt.Fprintln(writer)

	// 사전순으로 사람의 이름과 자식의 수, 사전순으로 자식들의 이름을 공백으로 구분하여 출력
	sort.Strings(family)
	for _, name := range family {
		fmt.Fprintf(writer, "%s ", name)
		if len(children[name]) == 0 {
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintf(writer, "%d ", len(children[name]))
			sort.Strings(children[name])
			for _, child := range children[name] {
				fmt.Fprintf(writer, "%s ", child)
			}
			fmt.Fprintln(writer)
		}
	}
}

func scanName() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
