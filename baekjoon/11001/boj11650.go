package bj11650

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

type Coord struct {
	x int
	y int
}

type Coords []Coord

func (c Coords) Len() int { return len(c) }
func (c Coords) Less(i, j int) bool {
	if c[i].x == c[j].x {
		return c[i].y < c[j].y
	}
	return c[i].x < c[j].x
}
func (c Coords) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var coords []Coord
	for i := 0; i < n; i++ {
		coords = append(coords, scanCoord())
	}
	sort.Sort(Coords(coords))

	for _, c := range coords {
		fmt.Fprintf(writer, "%d %d\n", c.x, c.y)
	}
}

func scanCoord() Coord {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	return Coord{x, y}
}
