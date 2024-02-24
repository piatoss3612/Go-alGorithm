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
	N, M    int
)

// 23056번: 참가자 명단
// https://www.acmicpc.net/problem/23056
// 난이도: 실버 4
// 메모리: 888 KB
// 시간: 4 ms
// 분류: 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
}

type Student struct {
	class int
	name  string
}

func Solve() {
	class := make([]int, N+1)

	odd := []Student{}
	even := []Student{}

	for {
		i, name := scanInt(), scanString()

		if i == 0 && name == "0" {
			break
		}

		if class[i] == M {
			continue
		}

		class[i]++

		if i%2 == 0 {
			even = append(even, Student{i, name})
		} else {
			odd = append(odd, Student{i, name})
		}
	}

	sort.Slice(odd, func(i, j int) bool {
		if odd[i].class == odd[j].class {
			if len(odd[i].name) == len(odd[j].name) {
				return odd[i].name < odd[j].name
			}
			return len(odd[i].name) < len(odd[j].name)
		}

		return odd[i].class < odd[j].class
	})

	sort.Slice(even, func(i, j int) bool {
		if even[i].class == even[j].class {
			if len(even[i].name) == len(even[j].name) {
				return even[i].name < even[j].name
			}
			return len(even[i].name) < len(even[j].name)
		}

		return even[i].class < even[j].class
	})

	for _, s := range odd {
		fmt.Fprintf(writer, "%d %s\n", s.class, s.name)
	}

	for _, s := range even {
		fmt.Fprintf(writer, "%d %s\n", s.class, s.name)
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
