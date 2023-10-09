use std::io::{stdin, Write};

// 7795번: 먹을 것인가 먹힐 것인가
// https://www.acmicpc.net/problem/7795
// 난이도: Silver 3
// 메모리: 13784KB
// 시간: 20ms
// 분류: 정렬, 두 포인터
fn main() {
    let mut lines = stdin().lines();

    let t = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    for i in 0..t {
        let (n, m) = {
            let line = lines.next().unwrap().unwrap();
            let mut iter = split_line_to_numbers(&line).into_iter();
            (iter.next().unwrap(), iter.next().unwrap())
        };

        let mut a = split_line_to_numbers(&lines.next().unwrap().unwrap()); // a.len() == n
        let mut b = split_line_to_numbers(&lines.next().unwrap().unwrap()); // b.len() == m

		// 정렬, 순서가 중요하지 않으므로 unstable
        a.sort_unstable();
        b.sort_unstable();

        let (mut l, mut r) = (0, 0); // 두 포인터 l, r 초기화

        let mut ans = 0;

        while l < n && r < m {
			// a[l] <= b[r] 이면, b에 있는 a[l]보다 작은 원소의 개수는 r개
            if a[l as usize] <= b[r as usize] {
                ans += r;
                l += 1;
            } else {
                r += 1;
            }
        }

        ans += (n - l) * r; // a에 남은 원소들은 b의 r개의 원소보다 모두 크므로, r개씩 더해준다.

        writeln!(writer, "{}", ans).unwrap();
    }
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
