package bj1297

import (
	"fmt"
	"math"
)

func main() {
	var d, h, w float64
	fmt.Scan(&d, &h, &w)
	a := math.Sqrt((d * d) / (h*h + w*w))
	fmt.Printf("%d %d", int(h*a), int(w*a))
}
