package bj1934

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		a, b := getTest()
		result := lcm(a, b, gcd(a, b))
		fmt.Fprintf(writer, "%d\n", result)
	}
}

func getTest() (a, b int) {
	scanner.Scan()
	nums := strings.Split(scanner.Text(), " ")
	a, _ = strconv.Atoi(nums[0])
	b, _ = strconv.Atoi(nums[1])
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b, gcd int) int {
	return a * b / gcd
}
