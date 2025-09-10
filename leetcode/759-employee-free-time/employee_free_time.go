package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

/**
 --------
-1------7------------------

*/

func main() {
	tests := []struct {
		schedule [][]Interval
		want     []Interval
	}{
		{
			// LeetCode Example 1
			schedule: [][]Interval{
				{{1, 2}, {5, 6}},
				{{1, 3}},
				{{4, 10}},
			},
			want: []Interval{{3, 4}},
		},
		{
			// LeetCode Example 2
			schedule: [][]Interval{
				{{1, 3}, {6, 7}},
				{{2, 4}},
				{{2, 5}, {9, 12}},
			},
			want: []Interval{{5, 6}, {7, 9}},
		},
		{
			// Meu exemplo custom (traiçoeiro)
			schedule: [][]Interval{
				{{1, 3}, {6, 7}},
				{{2, 4}, {8, 10}},
				{{1, 2}, {5, 9}},
			},
			want: []Interval{{4, 5}},
		},
		{
			// Exemplo com múltiplas lacunas
			schedule: [][]Interval{
				{{1, 2}, {5, 6}},
				{{2, 3}, {6, 8}},
				{{0, 1}, {3, 4}, {9, 12}},
			},
			want: []Interval{{4, 5}, {8, 9}},
		},
	}

	for i, tt := range tests {
		got := employeeFreeTime(tt.schedule)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("❌ Test %d failed\nGot:  %v\nWant: %v\n\n", i+1, got, tt.want)
		} else {
			fmt.Printf("✅ Test %d passed\n", i+1)
		}
	}
}

func employeeFreeTime(schedule [][]Interval) []Interval {
	eventsMap := make(map[int]int, 0)

	for _, interval := range schedule {
		for _, date := range interval {
			eventsMap[date.Start]++
			eventsMap[date.End]--
		}
	}

	keys := make([]int, 0, len(eventsMap))
	for key := range eventsMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	result := []Interval{}
	active := 0
	prevTime := keys[0]

	for _, key := range keys {
		if active == 0 && key > prevTime {
			result = append(result, Interval{Start: prevTime, End: key})
		}

		active += eventsMap[key]
		prevTime = key
	}

	return result
}
