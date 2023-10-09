package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, H    int
	rooms   []Room
)

type Room struct {
	t, a, h int
}

// 난이도: Gold 4
// 메모리: 5752KB
// 시간: 136ms
// 분류: 이분 탐색, 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, H = scanInt(), scanInt()
	rooms = make([]Room, N)
	for i := 0; i < N; i++ {
		rooms[i] = Room{scanInt(), scanInt(), scanInt()}
	}
}

func Solve() {
	l, r := 1, 1000000*1000000*123456
	for l <= r {
		m := (l + r) / 2
		if ClearDungeon(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	fmt.Fprintln(writer, l) // 문제 조건을 만족하는 lower bound 출력
}

func ClearDungeon(maxHP int) bool {
	curHP, curATK := maxHP, H

	for _, room := range rooms {
		switch room.t {
		case 1: // 전투
			// 현재 용사의 공격력으로 몬스터를 잡기 위해 필요한 턴 수
			turn := room.h / curATK
			if room.h%curATK != 0 {
				turn += 1
			}
			// 용사가 몬스터를 잡고 살아남는지 확인
			// 용사의 우선 공격권 고려, 몬스터의 턴 수는 1만큼 차감
			if room.a*(turn-1) >= curHP {
				return false
			}

			curHP -= room.a * (turn - 1)
		case 2: // 회복
			curATK += room.a
			curHP += room.h
			if curHP > maxHP {
				curHP = maxHP
			}
		}
	}

	return true // 던전 클리어
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
