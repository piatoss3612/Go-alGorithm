use std::io::{stdin, stdout, BufWriter, Write};

// 15996번: 팩토리얼 나누기
// https://www.acmicpc.net/problem/15996
// 난이도: 실버 3
// 메모리: 13216 KB
// 시간: 4 ms
// 분류: 수학, 정수론
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let (n, a) = {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        (line[0], line[1])
    };

    let mut cnt = 0;

    // n을 a로 나눈 몫이 x라고 했을 때,
    // n을 a^2로 나눈 몫은 x / a임을 이용
    let mut x = n / a;
    cnt += x;

    while x > 0 {
        x /= a;
        cnt += x;
    }

    writeln!(writer, "{}", cnt).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
