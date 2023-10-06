use std::cmp::{max, min};
use std::io::{stdin, Write};

// 6220번: Making Change
// https://www.acmicpc.net/problem/6220
// 난이도: Silver 1
// 메모리: 13156KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let (c, n) = {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        (line[0], line[1])
    };

    let mut dp = vec![987654321; c as usize + 1];

    dp[0] = 0;

    for _ in 1..=n {
        let coin = line_to_number(&lines.next().unwrap().unwrap());

        for i in coin..=c {
            dp[i as usize] = min(dp[i as usize], dp[(i - coin) as usize] + 1);
        }
    }

    writeln!(writer, "{}", dp[c as usize]).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
