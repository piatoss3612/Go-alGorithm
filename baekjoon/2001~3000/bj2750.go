package bj2750

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	"os"
	"sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	sort.Ints(nums)
	for _, v := range nums {
		fmt.Fprintln(writer, v)
	}
}
