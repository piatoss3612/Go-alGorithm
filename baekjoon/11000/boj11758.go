package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Point struct {
	x, y int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	p := make([]Point, 4)
	for i := 1; i <= 3; i++ {
		p[i] = Point{scanInt(), scanInt()}
	}

	/*
		ccw: counter clockwise - 세 점의 방향 관계를 구하는 알고리즘

		x1 x2 x3 x1		x1 x2 x3 x1
		  \  \  \	 -	  /  /  /
		y1 y2 y3 y1		y1 y2 y3 y1

		ccw = x1*y2 + x2*y3 + x3*y1 - (y1*x2 + y2*x3 + y3*x1) = (x2 -x1)(y3 - y1) - (x3 - x1)(y2 - y1)

		ccw가 0보다 큰 경우: 시계 방향
		ccw가 0인 경우: 일직선
		ccw가 0보다 작은 경우: 반시계 방향
	*/

	ccw := (p[2].x-p[1].x)*(p[3].y-p[1].y) - (p[3].x-p[1].x)*(p[2].y-p[1].y)

	if ccw > 0 {
		fmt.Fprintln(writer, 1)
	} else if ccw == 0 {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
