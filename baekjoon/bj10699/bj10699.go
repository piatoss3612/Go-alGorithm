package bj10699

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	fmt.Printf("%4d-%02d-%02d", year, int(month), day)
}
