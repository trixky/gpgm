package core

import "testing"

// TestStockReset tests the Reset methode of the Stock structure
func TestStockReset(t *testing.T) {
	tests := []struct {
		stock Stock
	}{
		{ // ----------------------- 0
			stock: Stock{},
		},
		{ // ----------------------- 1
			stock: Stock{
				"dollars": 42,
			},
		},
		{ // ----------------------- 2
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
		},
		{ // ----------------------- 3
			stock: Stock{
				"wood":  1,
				"stone": 2,
				"gold":  3,
			},
		},
	}

	const expected = 0

	for test_index, test := range tests {
		// For each test

		// Reset the stock
		test.stock.Reset()

		if stock_length := len(test.stock); stock_length > expected {
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected, stock_length)
		}
	}
}

// TestStockDeepCopy tests the DeepCopy methode of the Stock structure
func TestStockDeepCopy(t *testing.T) {
	tests := []struct {
		stock    Stock
		expected Stock
	}{
		{ // ----------------------- 0
			stock:    Stock{},
			expected: Stock{},
		},
		{ // ----------------------- 1
			stock: Stock{
				"dollars": 42,
			},
			expected: Stock{
				"dollars": 42,
			},
		},
		{ // ----------------------- 2
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			expected: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
		},
		{ // ----------------------- 3
			stock: Stock{
				"wood":  1,
				"stone": 2,
				"gold":  3,
			},
			expected: Stock{
				"wood":  1,
				"stone": 2,
				"gold":  3,
			},
		},
	}

	for test_index, test := range tests {
		// For each test

		// Make a deep copy of the stock
		stock_copy := test.stock.DeepCopy()
		// Reset the stock
		test.stock.Reset()

		stock_copy_length := len(*stock_copy)
		expected_length := len(test.expected)

		if expected_length != stock_copy_length {
			t.Fatalf(`test %d (length): expected = %d, got = %d`, test_index, expected_length, stock_copy_length)
		}

		for resource, quantity := range *stock_copy {
			// For each resource of the copied stock

			if quantity != test.expected[resource] {
				t.Fatalf(`test %d (resource: %s): expected = %d, got = %d`, test_index, resource, test.expected[resource], quantity)
			}
		}
	}
}

// TestStockRemoveResource tests the RemoveResource methode of the Stock structure
func TestStockRemoveResource(t *testing.T) {
	tests := []struct {
		stock    Stock
		resource string
		quantity int
		expected int
	}{
		{ // ----------------------- 0
			stock: Stock{
				"dollars": 42,
			},
			resource: "dollars",
			quantity: 2,
			expected: 40,
		},
		{ // ----------------------- 1
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			resource: "mouse",
			quantity: 3,
			expected: 0,
		},
		{ // ----------------------- 2
			stock: Stock{
				"wood":  1,
				"stone": 2,
				"gold":  3,
			},
			resource: "stone",
			quantity: 1,
			expected: 1,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Remove a quantity of resource of the stock
		test.stock.RemoveResource(test.resource, test.quantity)

		if quantity := test.stock[test.resource]; quantity != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %d, got = %d`, test_index, test.resource, test.expected, quantity)
		}
	}
}

// TestStockSetResource tests the SetResource methode of the Stock structure
func TestStockSetResource(t *testing.T) {
	tests := []struct {
		stock    Stock
		resource string
		quantity int
		expected int
	}{
		{ // ----------------------- 0
			stock: Stock{
				"dollars": 42,
			},
			resource: "dollars",
			quantity: 2,
			expected: 2,
		},
		{ // ----------------------- 1
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			resource: "mouse",
			quantity: 4,
			expected: 4,
		},
		{ // ----------------------- 2
			stock: Stock{
				"wood":  1,
				"stone": 1,
				"gold":  3,
			},
			resource: "stone",
			quantity: 1,
			expected: 1,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Set a quantity of resource of the stock
		test.stock.SetResource(test.resource, test.quantity)

		if quantity := test.stock[test.resource]; quantity != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %d, got = %d`, test_index, test.resource, test.expected, quantity)
		}
	}
}

// TestStockAddResource tests the AddResource methode of the Stock structure
func TestStockAddResource(t *testing.T) {
	tests := []struct {
		stock    Stock
		resource string
		quantity int
		expected int
	}{
		{ // ----------------------- 0
			stock: Stock{
				"dollars": 42,
			},
			resource: "dollars",
			quantity: 2,
			expected: 44,
		},
		{ // ----------------------- 1
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			resource: "mouse",
			quantity: 4,
			expected: 7,
		},
		{ // ----------------------- 2
			stock: Stock{
				"wood":  1,
				"stone": 1,
				"gold":  3,
			},
			resource: "stone",
			quantity: 1,
			expected: 2,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Add a quantity of resource of the stock
		test.stock.AddResource(test.resource, test.quantity)

		if quantity := test.stock[test.resource]; quantity != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %d, got = %d`, test_index, test.resource, test.expected, quantity)
		}
	}
}

// TestStockHaveResource tests the HaveResource methode of the Stock structure
func TestStockHaveResource(t *testing.T) {
	tests := []struct {
		stock    Stock
		resource string
		expected bool
	}{
		{ // ----------------------- 0
			stock: Stock{
				"dollars": 42,
			},
			resource: "dollars",
			expected: true,
		},
		{ // ----------------------- 1
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			resource: "mouse",
			expected: true,
		},
		{ // ----------------------- 2
			stock: Stock{
				"wood":  1,
				"stone": 1,
				"gold":  3,
			},
			resource: "ruby",
			expected: false,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Check if the stock has a resource
		have_resource := test.stock.HaveResource(test.resource)

		if have_resource != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %t, got = %t`, test_index, test.resource, test.expected, have_resource)
		}
	}
}

// TestStockGetResource tests the GetResource methode of the Stock structure
func TestStockGetResource(t *testing.T) {
	tests := []struct {
		stock    Stock
		resource string
		expected int
	}{
		{ // ----------------------- 0
			stock: Stock{
				"dollars": 42,
			},
			resource: "dollars",
			expected: 42,
		},
		{ // ----------------------- 1
			stock: Stock{
				"cat":   1,
				"dog":   2,
				"mouse": 3,
			},
			resource: "mouse",
			expected: 3,
		},
		{ // ----------------------- 2
			stock: Stock{
				"wood":  1,
				"stone": 1,
				"gold":  3,
			},
			resource: "ruby",
			expected: 0,
		},
	}

	for test_index, test := range tests {
		// For each test

		// Get the of a resource of the stock
		quantity := test.stock.GetResource(test.resource)

		if quantity != test.expected {
			t.Fatalf(`test %d (resource: %s): expected = %d, got = %d`, test_index, test.resource, test.expected, quantity)
		}
	}
}
