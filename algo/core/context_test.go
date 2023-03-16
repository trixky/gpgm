package core

import (
	"testing"
)

func TestFindProcessParents(t *testing.T) {
	tests := []struct {
		initial_context InitialContext
		expected        []Process
	}{
		{ // ----------------------- 0
			initial_context: InitialContext{
				Processes: []Process{
					{
						Name: "wood_stone",
						Inputs: map[string]int{
							"wood": 1,
						},
						Outputs: map[string]int{
							"stone": 1,
						},
					},
					{
						Name: "stone_gold",
						Inputs: map[string]int{
							"stone": 1,
						},
						Outputs: map[string]int{
							"gold": 1,
						},
					},
				},
			},
			expected: []Process{
				{
					Name:    "wood_stone",
					Parents: []int{1},
				},
				{
					Name:    "stone_gold",
					Parents: []int{},
				},
			},
		},
		{ // ----------------------- 1
			initial_context: InitialContext{
				Processes: []Process{
					{
						Name: "gold_wood",
						Inputs: map[string]int{
							"gold": 1,
						},
						Outputs: map[string]int{
							"wood": 1,
						},
					},
					{
						Name: "gold_stone",
						Inputs: map[string]int{
							"gold": 1,
						},
						Outputs: map[string]int{
							"stone": 1,
						},
					},
					{
						Name: "wood_stone_house",
						Inputs: map[string]int{
							"wood":  1,
							"stone": 1,
						},
						Outputs: map[string]int{
							"house": 1,
						},
					},
					{
						Name: "house_gold",
						Inputs: map[string]int{
							"house": 1,
						},
						Outputs: map[string]int{
							"gold": 1,
						},
					},
					{
						Name: "house_ruby",
						Inputs: map[string]int{
							"house": 1,
						},
						Outputs: map[string]int{
							"ruby": 1,
						},
					},
				},
			},
			expected: []Process{
				{
					Name:    "gold_wood",
					Parents: []int{2},
				},
				{
					Name:    "gold_stone",
					Parents: []int{2},
				},
				{
					Name:    "wood_stone_house",
					Parents: []int{3, 4},
				},
				{
					Name:    "house_gold",
					Parents: []int{0, 1},
				},
				{
					Name:    "house_ruby",
					Parents: []int{},
				},
			},
		},
	}

	for test_index, test := range tests {
		// For each test

		// Find process parents
		test.initial_context.FindProcessParents()

		for process_index, process := range test.initial_context.Processes {
			// For each process

			if len(process.Parents) != len(test.expected[process_index].Parents) {
				t.Fatalf(`test %d (process: %d): expected = %d, got = %d`, test_index, process_index, len(test.expected[process_index].Parents), len(process.Parents))
			}
		}
	}
}
