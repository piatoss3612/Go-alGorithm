package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/42583
// 분류: 큐, 시뮬레이션
func solution(bridge_length int, weight int, truck_weights []int) int {
	bridge := make([]int, bridge_length)
	time := 0
	totalWeight := 0

	for i := 0; i < len(truck_weights); {
		totalWeight -= bridge[0]
		bridge = bridge[1:]
		bridge = append(bridge, 0)

		if totalWeight+truck_weights[i] <= weight {
			bridge[bridge_length-1] = truck_weights[i]
			totalWeight += truck_weights[i]
			i++
		}

		time++
	}

	time += bridge_length // 마지막 트럭이 다리를 지나는 시간

	return time
}
