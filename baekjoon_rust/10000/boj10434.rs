use std::io::{stdin, stdout, BufWriter, Write};

// 10434번: 행복한 소수
// https://www.acmicpc.net/problem/10434
// 난이도: 실버 2
// 메모리: 13216 KB
// 시간: 8 ms
// 분류: 수학, 구현, 정수론, 시뮬레이션, 소수 판정
fn main() {
    let mut writer = BufWriter::new(stdout());
    let mut lines = stdin().lines();

    let mut is_prime = vec![true; 10001];
    is_prime[0] = false;
    is_prime[1] = false;
    let mut i = 2;
    while i * i <= 10000 {
        if is_prime[i] {
            let mut j = i * i;
            while j <= 10000 {
                is_prime[j] = false;
                j += i;
            }
        }
        i += 1;
    }

    let mut is_happy_prime = vec![true; 10001];
    is_happy_prime[0] = false;
    is_happy_prime[1] = false;

    for i in 2..=10000 {
        if is_prime[i] && is_happy_prime[i] {
            let mut n = i;
            let mut sum = 0;
            while n > 0 {
                sum += (n % 10) * (n % 10);
                n /= 10;
            }
            if sum == 1 {
                continue;
            }
            let mut visited = vec![false; 10001];
            visited[i] = true;
            while sum != 1 {
                if visited[sum as usize] {
                    is_happy_prime[i] = false;
                    break;
                }
                visited[sum as usize] = true;
                n = sum;
                sum = 0;
                while n > 0 {
                    sum += (n % 10) * (n % 10);
                    n /= 10;
                }
            }
        } else {
            is_happy_prime[i] = false;
        }
    }

    let p = line_to_number(&lines.next().unwrap().unwrap());

    for _ in 0..p {
        let (n, m) = {
            let nums = split_line_to_numbers(&lines.next().unwrap().unwrap());
            (nums[0], nums[1])
        };

        writeln!(
            writer,
            "{} {} {}",
            n,
            m,
            if is_happy_prime[m as usize] {
                "YES"
            } else {
                "NO"
            }
        )
        .unwrap();
    }
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}

fn line_to_number(s: &str) -> i32 {
    s.parse().unwrap()
}
