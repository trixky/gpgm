package instance

import (
	"math"
	"math/rand"
	"testing"

	"github.com/trixky/krpsim/algo/core"
)

// TestEntryGeneRandomCut tests the RandomCut methode of the EntryGene struct
func TestEntryGeneRandomCut(t *testing.T) {
	tests := []struct {
		process_ids []int
		luck        int
		expected    []int
	}{
		// --------- empty process ids
		{
			process_ids: []int{},
			luck:        0,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			luck:        1,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			luck:        10,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			luck:        math.MaxInt,
			expected:    []int{},
		},
		// --------- no cut
		{
			process_ids: []int{0, 1, 2, 3, 4},
			luck:        0,
			expected:    []int{0, 1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			luck:        3000,
			expected:    []int{0, 1, 2, 3, 4},
		},
		// --------- cut
		{
			process_ids: []int{0, 1, 2, 3, 4},
			luck:        2,
			expected:    []int{1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			luck:        3,
			expected:    []int{0, 1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			luck:        4,
			expected:    []int{2, 3, 4},
		},
		{
			process_ids: []int{4},
			luck:        1,
			expected:    []int{4},
		},
	}

	rand.Seed(42)

	for test_index, test := range tests {
		// For each test
		entry_gene := EntryGene{
			Process_ids: test.process_ids,
		}

		entry_gene.RandomCut(test.luck)

		if expected_length, got_length := len(test.expected), len(entry_gene.Process_ids); expected_length != got_length {
			// If the process ids length is corrupted
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected_length, got_length)
		}

		for entry_process_index, entry_process := range entry_gene.Process_ids {
			// For each entry process
			if expected := test.expected[entry_process_index]; entry_process != expected {
				// If the entry process id is corrupted
				t.Fatalf(`test %d (process: %d): expected = %d, got = %d`, test_index, entry_process_index, expected, entry_process)
			}
		}
	}
}

// TestEntryGeneCutN tests the CutN methode of the EntryGene struct
func TestEntryGeneCutN(t *testing.T) {
	tests := []struct {
		process_ids []int
		n           uint
		expected    []int
	}{
		// --------- empty process ids
		{
			process_ids: []int{},
			n:           0,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			n:           1,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			n:           10,
			expected:    []int{},
		},
		{
			process_ids: []int{},
			n:           math.MaxUint,
			expected:    []int{},
		},
		// --------- no cut
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           0,
			expected:    []int{0, 1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           5,
			expected:    []int{0, 1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           150,
			expected:    []int{0, 1, 2, 3, 4},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           math.MaxUint,
			expected:    []int{0, 1, 2, 3, 4},
		},
		// --------- cut
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           4,
			expected:    []int{0, 1, 2, 3},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           3,
			expected:    []int{0, 1, 2},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           2,
			expected:    []int{0, 1},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			n:           1,
			expected:    []int{0},
		},
	}

	for test_index, test := range tests {
		// For each test
		entry_gene := EntryGene{
			Process_ids: test.process_ids,
		}

		entry_gene.CutN(test.n)

		if expected_length, got_length := len(test.expected), len(entry_gene.Process_ids); expected_length != got_length {
			// If the process ids length is corrupted
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected_length, got_length)
		}

		for entry_process_index, entry_process := range entry_gene.Process_ids {
			// For each entry process
			if expected := test.expected[entry_process_index]; entry_process != expected {
				// If the entry process id is corrupted
				t.Fatalf(`test %d (process: %d): expected = %d, got = %d`, test_index, entry_process_index, expected, entry_process)
			}
		}
	}
}

// TestEntryGeneCutRandomN tests the CutRandomN methode of the EntryGene struct
func TestEntryGeneCutRandomN(t *testing.T) {
	tests := []struct {
		process_ids []int
		expected    []int
	}{
		// --------- empty process ids
		{
			process_ids: []int{},
			expected:    []int{},
		},
		{
			process_ids: []int{},
			expected:    []int{},
		},
		{
			process_ids: []int{},
			expected:    []int{},
		},
		{
			process_ids: []int{},
			expected:    []int{},
		},
		// --------- cut
		{
			process_ids: []int{0},
			expected:    []int{0},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4},
			expected:    []int{0, 1, 2},
		},
		{
			process_ids: []int{0, 1},
			expected:    []int{0, 1},
		},
		{
			process_ids: []int{0, 1, 2},
			expected:    []int{0, 1, 2},
		},
		{
			process_ids: []int{0, 1, 2},
			expected:    []int{0, 1},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			process_ids: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:    []int{0, 1, 2, 3, 4, 5},
		},
	}

	for test_index, test := range tests {
		// For each test
		entry_gene := EntryGene{
			Process_ids: test.process_ids,
		}

		entry_gene.CutRandomN()

		if expected_length, got_length := len(test.expected), len(entry_gene.Process_ids); expected_length != got_length {
			// If the process ids length is corrupted
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected_length, got_length)
		}

		for entry_process_index, entry_process := range entry_gene.Process_ids {
			// For each entry process
			if expected := test.expected[entry_process_index]; entry_process != expected {
				// If the entry process id is corrupted
				t.Fatalf(`test %d (process: %d): expected = %d, got = %d`, test_index, entry_process_index, expected, entry_process)
			}
		}
	}
}

// TestEntryGeneInit tests the Init methode of the EntryGene struct
func TestEntryGeneInit(t *testing.T) {
	tests := []struct {
		processes []core.Process
		max       uint
		random    bool
		context   core.InitialContext
		expected  []int
	}{
		{
			processes: []core.Process{
				{ // ------------------------ 0
					Name: "gold_wood",
					Inputs: map[string]int{
						"gold": 1,
					},
					Outputs: map[string]int{
						"wood": 10,
					},
				},
				{ // ------------------------ 1
					Name: "gold_stone",
					Inputs: map[string]int{
						"gold": 1,
					},
					Outputs: map[string]int{
						"stone": 10,
					},
				},
				{ // ------------------------ 2
					Name: "wood_stone_house_1",
					Inputs: map[string]int{
						"wood":  10,
						"stone": 10,
					},
					Outputs: map[string]int{
						"house": 1,
					},
				},
				{ // ------------------------ 3
					Name: "sale_house",
					Inputs: map[string]int{
						"house": 1,
					},
					Outputs: map[string]int{
						"gold": 3,
					},
				},
				{ // ------------------------ 4
					Name: "sale_house_2",
					Inputs: map[string]int{
						"house": 2,
					},
					Outputs: map[string]int{
						"gold": 7,
					},
				},
			},
			max:    2,
			random: false,
			context: core.InitialContext{
				Stock: core.Stock{
					"gold":  10,
					"wood":  0,
					"stone": 0,
					"house": 0,
				},
				Optimize: map[string]bool{
					"gold": true,
				},
			},
			expected: []int{3},
		},
		{
			processes: []core.Process{
				{ // ------------------------ 0
					Name: "buy_cat_1",
					Inputs: map[string]int{
						"gold": 1,
					},
					Outputs: map[string]int{
						"cat": 1,
					},
				},
				{ // ------------------------ 1
					Name: "buy_cat_2",
					Inputs: map[string]int{
						"gold": 2,
					},
					Outputs: map[string]int{
						"cat": 3,
					},
				},
				{ // ------------------------ 2
					Name: "buy_cat_3",
					Inputs: map[string]int{
						"gold": 3,
					},
					Outputs: map[string]int{
						"cat": 6,
					},
				},
				{ // ------------------------ 3
					Name: "buy_cat_4",
					Inputs: map[string]int{
						"gold": 4,
					},
					Outputs: map[string]int{
						"cat": 11,
					},
				},
			},
			max:    3,
			random: true,
			context: core.InitialContext{
				Stock: core.Stock{
					"gold": 3,
					"cat":  1,
				},
				Optimize: map[string]bool{
					"cat": false,
				},
			},
			expected: []int{0},
		},
	}

	options := core.Options{
		RandomCut: true,
		MaxCut:    0,
	}

	rand.Seed(42)

	for test_index, test := range tests {
		// For each test
		entry_gene := EntryGene{}

		entry_gene.Init(test.processes, test.context.Optimize, &options)

		if expected_length, got_length := len(test.expected), len(entry_gene.Process_ids); expected_length != got_length {
			// If the process ids length is corrupted
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected_length, got_length)
		}

		for entry_process_index, entry_process := range entry_gene.Process_ids {
			// For each entry process
			if expected := test.expected[entry_process_index]; entry_process != expected {
				// If the entry process id is corrupted
				t.Fatalf(`test %d (process: %d): expected = %d, got = %d`, test_index, entry_process_index, expected, entry_process)
			}
		}
	}
}
