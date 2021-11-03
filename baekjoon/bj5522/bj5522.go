package bj5522

import "fmt"

func main() {
	var scores [5]int
	var result int
	for i := 0; i < len(scores); i++ {
		fmt.Scan(&scores[i])
		result += scores[i]
	}
	fmt.Println(result)
}
