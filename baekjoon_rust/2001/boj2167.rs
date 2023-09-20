use std::io::{stdin, Write};

fn main() {
    let mut lines = stdin().lines();

    let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
    let (n, m) = (line[0], line[1]);

    let mut sum = vec![vec![0; (m + 1) as usize]; (n + 1) as usize]; // 누적합을 저장할 벡터

    for i in 1..=n {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        for j in 1..=m {
            sum[i as usize][j as usize] = line[(j - 1) as usize]
                + sum[i as usize][(j - 1) as usize]
                + sum[(i - 1) as usize][j as usize]
                - sum[(i - 1) as usize][(j - 1) as usize]; // 누적합 계산
        }
    }

    let k = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];

    let mut writer = std::io::BufWriter::new(std::io::stdout()); // 출력을 위한 버퍼

    for _ in 0..k {
        let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
        let (i, j, x, y) = (line[0], line[1], line[2], line[3]);

        writeln!(
            writer,
            "{}",
            sum[x as usize][y as usize]
                - sum[x as usize][(j - 1) as usize]
                - sum[(i - 1) as usize][y as usize]
                + sum[(i - 1) as usize][(j - 1) as usize]
        )
        .unwrap(); // 누적합을 이용해 부분합을 계산하여 출력
    }
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
