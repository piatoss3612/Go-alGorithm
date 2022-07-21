package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	number  []Number // sort.Slice 메서드를 사용하기 위해 슬라이스로 선언
	N       int
)

// 알파벳으로 표현된 수의 정보
type Number struct {
	weight int  // 해당 알파벳으로 표현되는 총 가중치
	isHead bool // 가장 앞자리에 주어졌는지 여부
}

// 메모리: 908KB
// 시간: 4ms
// 그리디 알고리즘
// 요즘 로직을 찾는 과정은 수월해 졌는데, 문제 풀이에서 자꾸 미끄러지니까 너무 하기 싫다...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	// 알파벳 A~J를 슬라이스의 인덱스 0~9로 맵핑한다
	// 즉, 'A'는 0('A'-'A') 'J'는 9('J'-'A')가 된다
	number = make([]Number, 10)

	N = scanInt()
	for i := 0; i < N; i++ {
		scanNumber()
	}

	// 가중치를 기준으로 오름차순 정렬
	// 여기서부터는 어느 자리가 어떤 알파벳인지는 중요하지 않다
	sort.Slice(number, func(i, j int) bool {
		return number[i].weight < number[j].weight
	})

	/*
		문제 조건:

		가능한 수의 합 중 최댓값을 구하는데 0으로 시작하는 수는 없어야 한다

		A~J 각각의 알파벳은 중복되지 않게 0~9와 맵핑되는데 0으로 시작하는 수가 없어야 하므로
		단순히 가장 가중치가 큰 값부터 9~0의 값을 부여하게 되면 오답이 된다

		최적해 찾기:

		문제의 입력 조건을 보면 '적어도 한 알파벳은 수의 가장 처음에 주어지지 않는다'라고 주어진다

		즉, 수의 가장 처음에 주어지지 않은 알파벳은 반드시 존재하며
		그러한 알파벳이 여러 개 있는 경우, 가중치가 가장 작은 알파벳을 0과 맵핑하면 최적해를 찾을 수 있다

		물론 0을 사용하지 않고도 문제를 풀 수 있는 경우 또는
		0과 맵핑된 알파벳이 수의 가장 처음에 주어지지 않는다면 이 과정을 건너뛰어도 좋다
	*/

	// 0과 맵핑된 알파벳이 있고 그 알파벳이 수의 가장 처음에 주어지는 경우
	if number[0].isHead {
		for i := 1; i < 10; i++ {
			// 수의 가장 처음에 주어지지 않는 알파벳 중 가중치가 가장 작은 i번째 알파벳을 찾는다
			// 그리고 i번째 알파벳 뒤에 있는 알파벳들을 앞으로 1칸씩 시프트하고
			// i번째 알파벳을  0번째 자리로 옮긴다
			if !number[i].isHead {
				temp := number[i]
				for j := i; j >= 1; j-- {
					number[j] = number[j-1]
				}
				number[0] = temp
				break
			}
		}
	}

	ans := 0

	// 가중치가 작은 값부터 0~9를 곱하여 누적합을 구한다
	for i := 0; i < 10; i++ {
		ans += number[i].weight * i
	}

	fmt.Fprintln(writer, ans)
}

// byte 타입의 슬라이스로 문자열을 입력받아 각각의 알파벳의 자릿수에 해당하는 가중치를 누적해준다
// 예를 들어, ABC, BCA 입력받으면 A의 가중치는 101, B의 가중치는 110, C의 가중치는 11이 된다
func scanNumber() {
	scanner.Scan()
	b := scanner.Bytes()
	n := len(b)
	digit := 1 // 1의 자릿수부터 시작

	for i := n - 1; i >= 0; i-- {
		number[b[i]-'A'].weight += digit // 가중치 누적
		digit *= 10                      // 자릿수 증가
	}

	number[b[0]-'A'].isHead = true // 가장 앞자리에 주어진 수 확인
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
