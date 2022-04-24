package bj1152

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)

	// s가 공백으로 이루어진 문자열일 경우, Split()을 실행하면 0이 아닌 자기 자신이 반환되어 1이 반환되는 문제
	// 따라서 Split을 실행하기 전에 문자열 좌우의 공백을 모두 TrimSpace()로 제거하고
	// 문자열이 비어있는 경우 먼저 조건검사를 실행
	if s == "" {
		fmt.Println(0)
	} else {
		slice := strings.Split(s, " ")
		fmt.Println(len(slice))
	}
}
