package task1

import (
	"testing"
)

func TestGraphMatrix_BFS(t *testing.T) {
	tests := []struct {
		name        string
		graph       GraphMatrix
		startVertex int
		expected    []int
	}{
		{
			/*
				  0
				 / \
				1   4
				 \   |
				  2  3
			*/
			name: "обычный граф",
			graph: GraphMatrix{
				costMatrix: [][]int{
					{0, 1, 0, 0, 1},
					{1, 0, 1, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 0, 0, 0, 1},
					{0, 0, 0, 1, 0},
				},
			},
			startVertex: 0,
			expected:    []int{0, 1, 4, 2, 3},
		},
		{
			name: "граф с одним узлом",
			graph: GraphMatrix{
				costMatrix: [][]int{
					{0},
				},
			},
			startVertex: 0,
			expected:    []int{0},
		},
		{
			name: "граф без вершин",
			graph: GraphMatrix{
				costMatrix: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			startVertex: 0,
			expected:    []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.graph.bfsCost(tt.startVertex)

			if len(got) != len(tt.expected) {
				t.Errorf("длина массива %v в результате выполнения bfsCost, отличается от ожидаемой длины массива %v", got, tt.expected)
				return
			}
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("массив элементов bfsCost = %v, отличается от ожидаемого %v", got, tt.expected)
					return
				}
			}
		})
	}
}
