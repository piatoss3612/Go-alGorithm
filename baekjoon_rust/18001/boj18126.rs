use std::io::{stdin, Write};

#[derive(Clone)]
struct Edge {
    to: i32,
    cost: i32,
}

// 18126번: 너구리 구구
// https://www.acmicpc.net/problem/18126
// 난이도: Silver 2
// 메모리: 13912KB
// 시간: 4ms
// 분류: 그래프 이론, 트리, 깊이 우선 탐색, 너비 우선 탐색
fn main() {
    let mut writer = std::io::BufWriter::new(std::io::stdout());
    let mut lines = stdin().lines();

    let n = split_line_to_numbers(&lines.next().unwrap().unwrap())[0];
    let mut graph = vec![Vec::<Edge>::new(); n as usize + 1];

    for _ in 1..n {
        let (a, b, c) = {
            let line = split_line_to_numbers(&lines.next().unwrap().unwrap());
            (line[0], line[1], line[2])
        };

        // 양방향 그래프
        graph[a as usize].push(Edge { to: b, cost: c });
        graph[b as usize].push(Edge { to: a, cost: c });
    }

    let mut dist: Vec<i64> = vec![0; n as usize + 1]; // 최대 거리를 저장할 배열
    let mut visited: Vec<bool> = vec![false; n as usize + 1]; // 방문 여부를 저장할 배열

    dfs(&graph, &mut dist, &mut visited, 1); // 1번 노드에서 시작하여 최대 거리를 구함

    let mut max_dist = 0; // 최대 거리를 가지는 노드를 찾음

    for i in 1..n + 1 {
        if dist[i as usize] > dist[max_dist as usize] {
            max_dist = i;
        }
    }

    writeln!(writer, "{}", dist[max_dist as usize]).unwrap(); // 최대 거리를 출력
}

fn dfs(graph: &Vec<Vec<Edge>>, dist: &mut Vec<i64>, visited: &mut Vec<bool>, start: i32) {
    visited[start as usize] = true; // 방문 표시

    for edge in &graph[start as usize] {
        if !visited[edge.to as usize] {
            // 방문하지 않은 노드라면
            dist[edge.to as usize] = std::cmp::max(
                dist[edge.to as usize],
                dist[start as usize] + edge.cost as i64,
            ); // 최대 거리를 저장
            dfs(graph, dist, visited, edge.to); // 재귀 호출
        }
    }
}

fn split_line_to_numbers(s: &str) -> Vec<i32> {
    s.split_whitespace().map(|s| s.parse().unwrap()).collect()
}
