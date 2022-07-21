package bj17286

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	coords  [][2]float64
	counts  []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	coords = make([][2]float64, 4)
	for i := 0; i <= 3; i++ {
		coords[i] = [2]float64{scanFloat64(), scanFloat64()}
	}
	checked := make([]bool, 4)
	checked[0] = true
	getMinDistance(0, 0, checked, float64(0))

	sort.Ints(counts)
	fmt.Fprintln(writer, counts[0])
}

func getDistance(a, b, c, d float64) float64 {
	return math.Sqrt((c-a)*(c-a) + (d-b)*(d-b))
}

func getMinDistance(curIdx, count int, checked []bool, distance float64) {
	if count == 3 {
		counts = append(counts, int(distance))
		return
	}

	for i := 1; i <= 3; i++ {
		if checked[i] {
			continue
		}
		checked[i] = true
		tmp := getDistance(coords[curIdx][0], coords[curIdx][1], coords[i][0], coords[i][1])
		getMinDistance(i, count+1, checked, distance+tmp)
		checked[i] = false
	}
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
}
