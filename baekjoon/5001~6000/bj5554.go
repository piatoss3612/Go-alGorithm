package bj5554

import "fmt"

func main() {
	var HtoS, StoP, PtoA, AtoH int
	fmt.Scan(&HtoS)
	fmt.Scan(&StoP)
	fmt.Scan(&PtoA)
	fmt.Scan(&AtoH)

	sum := HtoS + StoP + PtoA + AtoH
	fmt.Println(sum / 60)
	fmt.Println(sum % 60)
}
