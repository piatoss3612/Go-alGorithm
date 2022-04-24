package bj1267

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n)
	for i := 0; i < len(t); i++ {
		fmt.Scanf("%d", &t[i])
	}

	y := youngsik(t)
	m := minsik(t)

	if y > m {
		fmt.Printf("M %d\n", m)
	} else if y < m {
		fmt.Printf("Y %d\n", y)
	} else {
		fmt.Printf("Y M %d\n", y)
	}
}

func youngsik(n []int) int {
	bill := 0
	for i := 0; i < len(n); i++ {
		bill += ((n[i] / 30) + 1) * 10
	}
	return bill
}

func minsik(n []int) int {
	bill := 0
	for i := 0; i < len(n); i++ {
		bill += ((n[i] / 60) + 1) * 15
	}
	return bill
}
