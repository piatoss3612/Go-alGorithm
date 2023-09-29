use std::{
    collections::HashMap,
    io::{stdin, Write},
};

fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let n = lines.next().unwrap().unwrap().parse::<i32>().unwrap();

    let mut fibo: HashMap<i32, i32> = HashMap::new();

    let mut a = 1;
    let mut b = 1;

    for _ in 1..30 {
        let c = a + b;
        fibo.insert(c, b);
        a = b;
        b = c;
    }

    let mut maxDp = vec![0; n as usize + 1];
    let mut minDp = vec![987654321; n as usize + 1];

    minDp[0] = 0;

    for i in 1..=n {
        fibo.iter().for_each(|(&k, &v)| {
            if i >= k {
                maxDp[i as usize] = std::cmp::max(maxDp[i as usize], maxDp[(i - k) as usize] + v);
                minDp[i as usize] = std::cmp::min(minDp[i as usize], minDp[(i - k) as usize] + v);
            }
        });
    }

    writeln!(writer, "{} {}", minDp[n as usize], maxDp[n as usize]).unwrap();
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
