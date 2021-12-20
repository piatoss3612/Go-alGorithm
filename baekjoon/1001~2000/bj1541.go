package bj1541

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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	slice := strings.Split(scanner.Text(), "")
	slice = append(slice, "")

	flag := false // - 부호가 나왔는지 여부
	result := 0
	tmp := ""
	// '+' 연산을 괄호로 묶어 미리 계산하고 나온 값을 모두 빼주면
	// 최솟값, 즉 최적해를 찾을 수 있다
	// 따라서 최초로 발견한 '-' 연산부호 뒤에 나오는 모든 값을 빼주면 된다
	for i := 0; i < len(slice); i++ {
		if slice[i] == "+" || slice[i] == "-" || i == len(slice)-1 {
			n, _ := strconv.Atoi(tmp)
			if flag == true {
				result -= n
			} else {
				result += n
			}
			tmp = ""
		} else {
			tmp += slice[i]
		}
		if slice[i] == "-" {
			flag = true
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
