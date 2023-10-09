use std::io::{stdin, Write};

// 169395번: 파스칼의 삼각형
// https://www.acmicpc.net/problem/16395
// 난이도: Silver 5
// 메모리: 13160KB
// 시간: 4ms
// 분류: 수학, 다이나믹 프로그래밍, 조합론
fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let (n, k) = (line[0], line[1]);

    let mut dp = vec![vec![0; 30]; 30]; // dp[i][j]: iCj

    dp[0][0] = 1; // 0C0 = 1
    dp[1][0] = 1; // 1C0 = 1
    dp[1][1] = 1; // 1C1 = 1

    // 파스칼의 삼각형을 만든다.
    for i in 2..=29 {
        for j in 0..=i {
            if j == 0 || j == i {
                dp[i][j] = 1;
            } else {
                dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
            }
        }
    }

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    writeln!(writer, "{}", dp[n as usize - 1][k as usize - 1]).unwrap(); // n번째 줄의 k번째 수 출력
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
