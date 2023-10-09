package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 메모리: 928KB
// 시간: 4ms
// 그리디 알고리즘, 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	b := scanBytes()
	result := []byte{}

	for len(b) > 0 {
		c := b[0]
		b = b[1:]

		if len(result) == 0 {
			result = append(result, c)
			continue
		}

		// 사전순으로 가장 앞에 오는 문자열을 만들기 위한 그리디 알고리즘

		// 조건:
		// 문자열은 처음부터 i만큼 뒤집어야 한다

		// 입력값에서 꺼내온 문자 c가
		// result의 가장 뒤에 있는 문자보자 작으면서
		// result의 가장 앞에 있는 문자보자 작거나 같으면
		// result를 한 번 뒤집은 뒤에 c를 추가하고 다시 뒤집어준다
		// 이 과정은 result의 앞에 c를 추가해주는 것과 동일하다
		if c < result[len(result)-1] && c <= result[0] {
			result = append([]byte{c}, result...)
		} else {
			result = append(result, c)
		}
	}

	fmt.Fprintln(writer, string(result))
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
