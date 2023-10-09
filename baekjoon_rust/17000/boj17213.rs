use std::io::{stdin, Write};

// 17213번: 과일 서리
// https://www.acmicpc.net/problem/17213
// 메모리: 13160KB
// 시간: 0ms
// 분류: 수학, 다이나믹 프로그래밍, 조합론
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];
    let m = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let mut dp = vec![vec![0; m as usize + 1]; n as usize + 1];

    dp[1] = vec![1; m as usize + 1]; // 1개의 과일을 m개의 그릇에 담는 방법은 1가지 뿐

    // 과일의 종류가 i개일 때, i개 미만의 과일을 m개의 그릇에 담는 방법은 없으므로
    // j < i 일 때 dp[i][j] = 0
    for i in 2..=n {
        for j in i..=m {
            // dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
            dp[i as usize][j as usize] =
                dp[i as usize][j as usize - 1] + dp[i as usize - 1][j as usize - 1];
        }
    }

    writeln!(writer, "{}", dp[n as usize][m as usize]).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
