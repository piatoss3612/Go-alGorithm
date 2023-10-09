use std::io::{stdin, Write};

// 25418번: 정수 a를 k로 만들기
// https://www.acmicpc.net/problem/25418
// 난이도: Silver 3
// 메모리: 17064KB
// 시간: 8ms
// 분류: 다이나믹 프로그래밍
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let (a, k) = {
        let line = lines.next().unwrap().unwrap();
        let mut iter = split_line_to_numbers(&line).into_iter();
        (iter.next().unwrap(), iter.next().unwrap())
    };

    let mut dp = vec![987654321; 1000001];
    dp[a as usize] = 0;

    for i in a + 1..=k {
        let mut min = dp[i as usize - 1] + 1;
        if i % 2 == 0 {
            if min > dp[i as usize / 2] + 1 {
                min = dp[i as usize / 2] + 1;
            }
        }

        dp[i as usize] = min;
    }

    writeln!(writer, "{}", dp[k as usize]).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
