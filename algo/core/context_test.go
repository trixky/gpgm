package core

import (
	"testing"
)

func TestCanBeExecutedXTimes(t *testing.T) {
	tests := []struct {
		stock   Stock
		process Process
		amounts map[int]bool
	}{
		{
			stock: Stock{
				"wood":  3,
				"stone": 2,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood": 1,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			amounts: map[int]bool{ // [amount]expected
				1: true,
				2: true,
				3: true,
				4: false,
				5: false,
				6: false,
			},
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 2,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood":  1,
					"stone": 2,
				},
				Outputs: map[string]int{
					"gold": 4,
				},
				Delay: 100,
			},
			amounts: map[int]bool{ // [amount]expected
				1: true,
				2: false,
				3: false,
				4: false,
			},
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 10,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"stone": 2,
					"gold":  1,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			amounts: map[int]bool{ // [amount]expected
				1: true,
				2: false,
				3: false,
				4: false,
			},
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood":  1,
					"stone": 2,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			amounts: map[int]bool{ // [amount]expected
				1: false,
				2: false,
				3: false,
				4: false,
			},
		},
	}

	for test_index, test := range tests {
		for amount, expected := range test.amounts {
			if result := test.process.CanBeExecutedXTimes(&test.stock, amount); result != expected {
				t.Fatalf(`test %d (amount: %d): expected = %t, got = %t`, test_index, amount, expected, result)
			}
		}
	}
}

func TestCanBeExecutedMaxXTimes(t *testing.T) {
	tests := []struct {
		stock    Stock
		process  Process
		expected int
	}{
		{
			stock: Stock{
				"wood":  3,
				"stone": 2,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood": 1,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			expected: 3,
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 2,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood":  1,
					"stone": 2,
				},
				Outputs: map[string]int{
					"gold": 4,
				},
				Delay: 100,
			},
			expected: 1,
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 10,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"stone": 2,
					"gold":  1,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			expected: 1,
		},
		{
			stock: Stock{
				"wood":  3,
				"stone": 1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood":  1,
					"stone": 2,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			expected: 0,
		},
		{
			stock: Stock{
				"wood":  30,
				"stone": 10,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood":  3,
					"stone": 1,
				},
				Outputs: map[string]int{
					"wood":  60,
					"stone": 20,
				},
				Delay: 100,
			},
			expected: 10,
		},
	}

	for test_index, test := range tests {
		if result := test.process.CanBeExecutedMaxXTimes(&test.stock); result != uint16(test.expected) {
			t.Fatalf(`test %d: expected = %d, got = %d`, test_index, test.expected, result)
		}
	}
}

func TestExecuteN(t *testing.T) {
	tests := []struct {
		stock    Stock
		process  Process
		expected map[int]Stock
	}{
		{
			stock: Stock{
				"wood":  3,
				"stone": 2,
				"gold":  1,
			},
			process: Process{
				Name: "test_process",
				Inputs: map[string]int{
					"wood": 1,
				},
				Outputs: map[string]int{
					"gold": 2,
				},
				Delay: 100,
			},
			expected: map[int]Stock{
				1: {
					"wood":  2,
					"stone": 2,
					"gold":  1,
				},
				2: {
					"wood":  1,
					"stone": 2,
					"gold":  1,
				},
				3: {
					"wood":  0,
					"stone": 2,
					"gold":  1,
				},
			},
		},
	}

	for test_index, test := range tests {
		func() {
			for n, expected := range test.expected {

				stock_cpy := test.stock.DeepCopy()

				test.process.ExecuteN(&stock_cpy, n)

				for key, resource := range stock_cpy {
					if resource != expected[key] {
						t.Fatalf(`test %d (n: %d, resource: %s): expected = %d, got = %d`, test_index, n, key, expected[key], resource)
					}
				}
			}
		}()
	}
}
