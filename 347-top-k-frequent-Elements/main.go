package main

import (
	"sort"
)

func reverseMap(m map[int]int) map[int][]int {
	rm := map[int][]int{}
	for k, v := range m {
		rm[v] = append(rm[v], k)
	}
	return rm
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	for _, num := range nums {
		if _, exists := m[num]; exists {
			m[num] += 1
		} else {
			m[num] = 1
		}
	}

	rm := reverseMap(m)

	sortedKeys := make([]int, len(rm))
	i := 0
	for k, _ := range rm {
		sortedKeys[i] = k
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))

	values := make([]int, k)
	index := 0
	for _, key := range sortedKeys {
		for _, v := range rm[key] {
			values[index] = v
			index++
			if index == k {
				return values
			}
		}
	}
	return values
}
