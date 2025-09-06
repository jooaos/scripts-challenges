package main

import "testing"

func TestTotalDays(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "Nenhuma reunião",
			intervals: [][]int{},
			want:      0,
		},
		{
			name:      "Uma única reunião",
			intervals: [][]int{{1, 5}},
			want:      1,
		},
		{
			name:      "Sem sobreposição",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      1,
		},
		{
			name:      "Sobreposição parcial",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      2,
		},
		{
			name:      "Total sobreposição",
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			want:      4,
		},
		{
			name:      "Reuniões intercaladas",
			intervals: [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:      2,
		},
		{
			name:      "Grandes intervalos e pequenos dentro",
			intervals: [][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}},
			want:      2,
		},
		{
			name:      "Reuniões iguais",
			intervals: [][]int{{1, 5}, {1, 5}, {1, 5}},
			want:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalDays(tt.intervals)
			if got != tt.want {
				t.Errorf("TotalDays(%v) = %d, want %d", tt.intervals, got, tt.want)
			}
		})
	}
}
