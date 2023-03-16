package core

import "testing"

func TestProcessCanBeExecutedXTimes(t *testing.T) {
	tests := []struct {
		stock   Stock
		process Process
		amounts map[int]bool
	}{
		{ // ----------------------- 0
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
		{ // ----------------------- 1
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
		{ // ----------------------- 2
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
		{ // ----------------------- 3
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

func TestProcessCanBeExecutedMaxXTimes(t *testing.T) {
	tests := []struct {
		stock    Stock
		process  Process
		expected int
	}{
		{ // ----------------------- 0
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
		{ // ----------------------- 1
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
		{ // ----------------------- 2
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
		{ // ----------------------- 3
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
		{ // ----------------------- 4
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
		// For each test

		if result := test.process.CanBeExecutedMaxXTimes(&test.stock); result != test.expected {
			t.Fatalf(`test %d: expected = %d, got = %d`, test_index, test.expected, result)
		}
	}
}

func TestProcessExecuteN(t *testing.T) {
	tests := []struct {
		stock    Stock
		process  Process
		expected map[int]Stock
	}{
		{ // ----------------------- 0
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
		// For each test

		func() {
			for n, expected := range test.expected {
				// For each expected result corresponding to n

				// Make a deep copy of the initial stock
				stock_cpy := test.stock.DeepCopy()

				// Execute the process n time
				test.process.ExecuteN(stock_cpy, n)

				for key, resource := range *stock_cpy {
					// For each resource of the copied stock

					if resource != expected[key] {
						t.Fatalf(`test %d (n: %d, resource: %s): expected = %d, got = %d`, test_index, n, key, expected[key], resource)
					}
				}
			}
		}()
	}
}

func TestProcessHaveInput(t *testing.T) {
	tests := []struct {
		process  Process
		resource string
		expected bool
	}{
		{ // ----------------------- 0
			process: Process{
				Inputs: map[string]int{
					"cat":   3,
					"dog":   3,
					"mouse": 3,
				},
			},
			resource: "cat",
			expected: true,
		},
		{ // ----------------------- 1
			process: Process{
				Inputs: map[string]int{
					"wood":  3,
					"stone": 3,
					"gold":  3,
				},
			},
			resource: "cat",
			expected: false,
		},
	}

	for test_index, test := range tests {
		// For each test

		result := test.process.HaveInput(test.resource)

		if result != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %t, got = %t`, test_index, test.resource, test.expected, result)
		}
	}
}

func TestProcessHaveOutput(t *testing.T) {
	tests := []struct {
		process  Process
		resource string
		expected bool
	}{
		{ // ----------------------- 0
			process: Process{
				Outputs: map[string]int{
					"cat":   3,
					"dog":   3,
					"mouse": 3,
				},
			},
			resource: "cat",
			expected: true,
		},
		{ // ----------------------- 1
			process: Process{
				Outputs: map[string]int{
					"wood":  3,
					"stone": 3,
					"gold":  3,
				},
			},
			resource: "cat",
			expected: false,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Check if the process have a resource output
		result := test.process.HaveOutput(test.resource)

		if result != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %t, got = %t`, test_index, test.resource, test.expected, result)
		}
	}
}
