use std::io::{stdin, Write};

// 2018번: 수들의 합 5
// https://www.acmicpc.net/problem/2018
// 난이도: Silver 5
// 메모리: 13152KB
// 시간: 16ms
// 분류: 수학, 투 포인터
fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let n = line[0];

    // n이 2이하일 경우 연속된 수의 합으로 n을 만들 수 있는 경우는 자기 자신 뿐이다.
    if n <= 2 {
        println!("{}", 1);
        return;
    }

    let (mut l, mut r) = (1, 2); // l: 왼쪽 포인터, r: 오른쪽 포인터

    let mut cnt = 1; // 연속된 수의 합으로 n을 만들 수 있는 경우의 수 (초기값: 자기 자신)
    let mut sum = l + r; // 연속된 수의 합

    // l과 r이 n을 넘어가면 종료
    while r < n {
        if sum < n {
            // sum이 n보다 작으면 r을 증가시켜 sum을 증가시킨다.
            r += 1;
            sum += r;
        } else if sum > n {
            // sum이 n보다 크면 l을 증가시켜 sum을 감소시킨다.
            sum -= l;
            l += 1;
        } else {
            // sum == n
            cnt += 1;
            r += 1;
            sum += r;
        }
    }

    println!("{}", cnt); // 결과 출력
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
