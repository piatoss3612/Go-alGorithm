use std::io::{stdin, stdout, BufWriter, Write};

// 6219번: 소수의 자격
// https://www.acmicpc.net/problem/6219
// 난이도: 실버 3
// 메모리: 17128 KB
// 시간: 40ms
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let (a, b, d) = {
        let nums = split_line_to_numbers(&lines.next().unwrap().unwrap());
        (nums[0], nums[1], nums[2])
    };

    let mut isPrime = vec![true; b as usize + 1];

    isPrime[0] = false;
    isPrime[1] = false;

    let mut i = 2;
    while i * i <= b {
        if isPrime[i as usize] {
            let mut j = i * i;
            while j <= b {
                isPrime[j as usize] = false;
                j += i;
            }
        }
        i += 1;
    }

    let mut ans = 0;

    for i in a..=b {
        if isPrime[i as usize] {
            if i.to_string().contains(&d.to_string()) {
                ans += 1;
            }
        }
    }

    writeln!(writer, "{}", ans).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
