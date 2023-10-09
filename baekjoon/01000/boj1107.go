package bj1107

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n, m    int
	broken  [10]bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	for i := 1; i <= m; i++ {
		broken[scanInt()] = true
	}

	ans := int(math.Abs(float64(n - 100)))

	for i := 0; i <= 999999; i++ {
		s := strconv.Itoa(i)
		isBroken := false

		for j := 0; j < len(s); j++ {
			if broken[s[j]-'0'] {
				isBroken = true
				break
			}
		}

		if !isBroken {
			tmp := int(math.Abs(float64(n-i))) + len(s)
			ans = getMin(ans, tmp)
		}
	}

	fmt.Fprintln(writer, ans)
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
