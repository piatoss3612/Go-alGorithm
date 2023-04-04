package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner                    = bufio.NewScanner(os.Stdin)
	writer                     = bufio.NewWriter(os.Stdout)
	T                          int
	H, W, O, F, XS, YS, XE, YE int
	maze                       [101][101]int
	visited                    [101][101]bool
	dx                         = []int{-1, +0, +1, +0}
	dy                         = []int{+0, +1, +0, -1}
)

// 난이도: Gold 4
// 메모리: 1940KB
// 시간: 12ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	H, W, O, F, XS, YS, XE, YE = scanInt(), scanInt(), scanInt(), scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	maze = [101][101]int{}
	visited = [101][101]bool{}

	for i := 1; i <= O; i++ {
		x, y, l := scanInt(), scanInt(), scanInt()
		maze[x][y] = l
	}
}

type Inseong struct {
	x, y, f int
}

func Solve() {
	/*
		# 너비 우선 탐색

		1. 인성이는 (XS, YS)에서 초기 힘 F인 상태로 출발 (출발 위치를 방문 처리하고 큐에 삽입)
		2. 큐가 빌 때까지 또는 인성이가 (XE, YE)에 도달할 때까지 다음 과정을 반복
		3. 인성이는 힘이 0이 되면 이동할 수 없다
		4. 인성이는 상하좌우로 이동할 수 있으며, 이동할 때마다 힘은 1씩 감소
			4-1. 인성이가 더 높은 곳으로 이동할 때, (이동할 곳의 높이 - 현재 위치의 높이) 이상의 힘을 가지고 있어야 한다
			4-2. 인성이가 방문하지 않은 곳으로 이동할 때, 방문 처리하고 큐에 삽입
	*/
	q := []Inseong{}
	q = append(q, Inseong{x: XS, y: YS, f: F})
	visited[XS][YS] = true

	goal := false

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		cx := front.x
		cy := front.y
		cf := front.f

		// 목적지에 도착했을 때
		if cx == XE && cy == YE {
			goal = true
			break
		}

		// 힘이 0이 되면 이동할 수 없다
		if cf == 0 {
			continue
		}

		// 상하좌우로 이동
		for i := 0; i < 4; i++ {
			nx, ny := cx+dx[i], cy+dy[i]
			// 범위 내에 있고, 방문하지 않았을 때
			if valid(nx, ny) && !visited[nx][ny] {
				// 더 높은 곳으로 이동할 때
				if maze[nx][ny] > maze[cx][cy] {
					// (이동할 곳의 높이 - 현재 위치의 높이) 이상의 힘을 가지고 있어야 한다
					if cf >= maze[nx][ny]-maze[cx][cy] {
						visited[nx][ny] = true
						q = append(q, Inseong{x: nx, y: ny, f: cf - 1})
					} else {
						continue
					}
				} else {
					visited[nx][ny] = true
					q = append(q, Inseong{x: nx, y: ny, f: cf - 1})
				}
			}
		}
	}

	if goal {
		fmt.Fprintln(writer, "잘했어!!")
	} else {
		fmt.Fprintln(writer, "인성 문제있어??")
	}
}

func valid(x, y int) bool {
	return x >= 1 && x <= H && y >= 1 && y <= W
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
