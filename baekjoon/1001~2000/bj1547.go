package bj1547

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	cnt, _ := strconv.Atoi(r.Text())

	// 현재 공을 가지고 있는 컵
	pos := "1"

	for i := 0; i < cnt; i++ {
		r.Scan()
		cups := strings.Split(r.Text(), " ")
		swap(cups[0], cups[1], &pos)
	}
	check, _ := strconv.Atoi(pos)
	if check > 3 || check < 1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(pos)
}

func swap(a, b string, pos *string) {
	if *pos == a {
		*pos = b
	} else if *pos == b {
		*pos = a
	} else {
		return
	}
}
