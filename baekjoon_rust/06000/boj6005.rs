use std::cmp::max;
use std::io::{stdin, Write};

// 6005번: Cow Pinball
// https://www.acmicpc.net/problem/6005
// 난이도: Silver 1
// 메모리: 13160KB
// 시간: 0ms
// 분류: 다이나믹 프로그래밍
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let mut dp = vec![Vec::<i32>::new(); n as usize];

    for i in 0..n {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        dp[i as usize] = line;
    }

    for i in 1..n {
        for j in 0..i + 1 {
            if j == 0 {
                dp[i as usize][j as usize] += dp[i as usize - 1][j as usize];
            } else if i == j {
                dp[i as usize][j as usize] += dp[i as usize - 1][j as usize - 1];
            } else {
                dp[i as usize][j as usize] += max(
                    dp[i as usize - 1][j as usize],
                    dp[i as usize - 1][j as usize - 1],
                );
            }
        }
    }

    let mut ans = 0;

    for i in 0..n {
        ans = max(ans, dp[n as usize - 1][i as usize]);
    }

    writeln!(writer, "{}", ans).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
