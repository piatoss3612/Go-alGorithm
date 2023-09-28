use std::io::{stdin, Write};

// 1535번: 안녕
// https://www.acmicpc.net/problem/1535
// 난이도: Silver 2
// 메모리: 13156KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let n = lines.next().unwrap().unwrap().parse::<i32>().unwrap();

    let mut loss = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let mut joy = split_line_to_numbers(&lines.next().unwrap().unwrap());

    let mut dp = vec![-1; 100];

    dp[0] = 0;

    for i in 0..n {
        // 역순으로 진행해야 중복을 피할 수 있다.
        for j in (loss[i as usize]..100).rev() {
            if dp[j as usize - loss[i as usize] as usize] != -1 {
                dp[j as usize] = std::cmp::max(
                    dp[j as usize],
                    dp[j as usize - loss[i as usize] as usize] + joy[i as usize],
                );
            }
        }
    }

    let mut max = 0;
    for i in 0..100 {
        if dp[i] > max {
            max = dp[i];
        }
    }

    writeln!(writer, "{}", max).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
