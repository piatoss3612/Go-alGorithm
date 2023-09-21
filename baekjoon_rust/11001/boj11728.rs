use std::io::{stdin, Write};

// 11728번: 배열 합치기
// https://www.acmicpc.net/problem/11728
// 메모리: 40440KB
// 시간: 448ms
// 분류: 정렬, 두 포인터
fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let (n, m) = (line[0], line[1]);

    // a와 b는 정렬되어 있음
    let a = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let b = split_line_to_numbers(&lines.next().unwrap().unwrap());

    let (mut l, mut r) = (0, 0); // a를 가리키는 포인터, b를 가리키는 포인터

    let mut sorted: Vec<i32> = Vec::with_capacity((n + m) as usize); // 정렬된 배열

    // 두 배열을 합치면서 정렬
    while l < n && r < m {
        if a[l as usize] < b[r as usize] {
            sorted.push(a[l as usize]);
            l += 1;
        } else {
            sorted.push(b[r as usize]);
            r += 1;
        }
    }

    // 남은 배열을 합치면서 정렬
    while l < n {
        sorted.push(a[l as usize]);
        l += 1;
    }

    while r < m {
        sorted.push(b[r as usize]);
        r += 1;
    }

    let mut writer = std::io::BufWriter::new(std::io::stdout());

    // 출력
    for i in sorted {
        write!(writer, "{} ", i).unwrap();
    }
    write!(writer, "\n").unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
