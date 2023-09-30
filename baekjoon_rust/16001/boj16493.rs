use std::io::{stdin, Write};

fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let (n, m) = {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        let n: i32 = line[0];
        let m: i32 = line[1];
        (n, m)
    };

    let mut dp = vec![0; (n + 1) as usize];
    dp[0] = 1;

    for _ in 1..=m {
        let (a, b) = {
            let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
            let a: i32 = line[0];
            let b: i32 = line[1];
            (a, b)
        };

        for j in (0..=n).rev() {
            if j - a >= 0 && dp[(j - a) as usize] != 0 {
                dp[j as usize] = std::cmp::max(dp[j as usize], dp[(j - a) as usize] + b);
            }
        }
    }

    let mut ans = 0;

    for i in 0..=n {
        ans = std::cmp::max(ans, dp[i as usize]);
    }

    writeln!(writer, "{}", ans - 1).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
