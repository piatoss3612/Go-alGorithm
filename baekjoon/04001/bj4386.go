package bj4386

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// 정점에 대한 정보를 저장할 구조체
type node struct {
	x, y float64
}

// 간선에 대한 정보를 저장한 구조체
type link struct {
	a, b int
	c    float64
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	nodes   []node
	links   []link
	parent  []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	nodes = make([]node, n+1)
	parent = make([]int, n+1)

	// 정점의 x, y 좌표 입력 받기
	for i := 1; i <= n; i++ {
		nodes[i] = node{scanFloat(), scanFloat()}
	}

	// 부모 초기화
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	// 모든 정점 사이의 가능한 거리 정보를 계산하여 정점 a, 정점 b, 길이 c를 슬라이스에 저장
	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			links = append(links, link{i, j, getDistance(nodes[i], nodes[j])})
		}
	}

	// 간선의 길이가 짧은 것부터 오름차순으로 정렬
	sort.Slice(links, func(i, j int) bool {
		return links[i].c < links[j].c
	})

	var ans float64
	numberOfLinks := 0

	// 크루스칼 알고리즘
	for _, link := range links {
		if numberOfLinks == n-1 {
			break
		}

		if parent[link.a] != parent[link.b] {
			ans += link.c
			numberOfLinks += 1
			prev := parent[link.b]
			current := parent[link.a]
			for i := 1; i <= n; i++ {
				if parent[i] == prev {
					parent[i] = current
				}
			}
		}
	}
	fmt.Fprintf(writer, "%0.2f\n", ans)
}

func getDistance(a, b node) float64 {
	tmpX := math.Pow(b.x-a.x, 2)
	tmpY := math.Pow(b.y-a.y, 2)
	return math.Sqrt(tmpX + tmpY)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
