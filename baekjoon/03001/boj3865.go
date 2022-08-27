package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	academies map[string][]string // 학회 이름(key) : 학회 멤버들 (value)
	visited   map[string]bool     // 제일 처음에 주어지는 학회의 멤버들 탐색 여부
	n         int
)

// 메모리: 2540KB
// 시간: 8ms
// 해시 맵, 문자열 파싱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		n = scanInt()
		if n == 0 {
			return
		}

		academies = make(map[string][]string)
		var target string

		for i := 0; i < n; i++ {
			s := scanString()
			s = strings.TrimSuffix(s, ".")   // 접미사 . 제거
			splited := strings.Split(s, ":") // '학회 이름:멤버들 이름' 구조의 문자열을 ':'을 기준으로 문자열 슬라이스로 분리
			if i == 0 {
				target = splited[0] // 첫번째 확회명을 타겟으로 지정
			}
			academies[splited[0]] = strings.Split(splited[1], ",") // 학회 멤버들을 ','를 기준으로 문자열 슬라이스로 분리
		}

		visited = make(map[string]bool)
		visited[target] = true
		count := find(target, 0)

		fmt.Fprintln(writer, count)
	}
}

func find(name string, count int) int {
	// 매개변수로 전달받은 이름이 학회명인 경우
	if currentMembers, ok := academies[name]; ok {
		// 학회 멤버들 탐색
		for _, member := range currentMembers {
			// 현재 멤버 이름이 아직 확인되지 않은 경우
			if _, ok := visited[member]; !ok {
				visited[member] = true
				// 멤버 이름이 학회명이 아닌 경우
				if _, ok := academies[member]; !ok {
					count += 1

				} else {
					// 학회명인 경우
					// 해당 학회의 멤버들 탐색
					count += find(member, 0)
				}
			}
		}
	}
	return count
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
