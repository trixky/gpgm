package history

import (
	"testing"
)

func TestHistoryReset(t *testing.T) {
	tests := []struct {
		history       History
		processes_ids []int
		expected      History
	}{
		{ // ----------------------- 0
			history: History{},
			expected: History{
				ProcessIds: "",
			},
		},
		{ // ----------------------- 1
			history: History{
				ProcessIds: "3210",
			},
			expected: History{},
		},
		{ // ----------------------- 2
			history: History{
				ProcessIds: "BA@?>=<;:9876543210",
			},
			expected: History{},
		},
	}

	for test_index, test := range tests {
		// For each test

		// Reset the history
		test.history.Reset()

		if test.history.ProcessIds != test.expected.ProcessIds {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.expected.ProcessIds, test.history.ProcessIds)
		}
	}
}

func TestHistoryPushProcessId(t *testing.T) {
	tests := []struct {
		history       History
		processes_ids []int
		expected      History
	}{
		{ // ----------------------- 0
			history:       History{},
			processes_ids: []int{},
			expected: History{
				ProcessIds: "",
			},
		},
		{ // ----------------------- 1
			history:       History{},
			processes_ids: []int{0, 1, 2, 3},
			expected: History{
				ProcessIds: "3210",
			},
		},
		{ // ----------------------- 2
			history:       History{},
			processes_ids: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
			expected: History{
				ProcessIds: "BA@?>=<;:9876543210",
			},
		},
	}

	for test_index, test := range tests {
		// For each test

		for _, process_id := range test.processes_ids {
			// For each process id to push in the history
			// Push it
			test.history.PushProcessId(process_id)
		}

		if test.history.ProcessIds != test.expected.ProcessIds {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.history.ProcessIds, test.expected.ProcessIds)
		}
	}
}

func TestHistoryGetLastProcessIds(t *testing.T) {
	tests := []struct {
		history  History
		n        int
		expected string
	}{
		{ // ----------------------- 0
			history:  History{},
			n:        0,
			expected: "",
		},
		{ // ----------------------- 1
			history:  History{},
			n:        1,
			expected: "",
		},
		{ // ----------------------- 2
			history:  History{},
			n:        2,
			expected: "",
		},
		{ // ----------------------- 3
			history: History{
				ProcessIds: "3210",
			},
			n:        0,
			expected: "3210",
		},
		{ // ----------------------- 4
			history: History{
				ProcessIds: "3210",
			},
			n:        1,
			expected: "3",
		},
		{ // ----------------------- 5
			history: History{
				ProcessIds: "3210",
			},
			n:        2,
			expected: "32",
		},
		{ // ----------------------- 6
			history: History{
				ProcessIds: "3210",
			},
			n:        3,
			expected: "321",
		},
		{ // ----------------------- 7
			history: History{
				ProcessIds: "3210",
			},
			n:        4,
			expected: "3210",
		},
		{ // ----------------------- 8
			history: History{
				ProcessIds: "3210",
			},
			n:        5,
			expected: "3210",
		},
	}

	for test_index, test := range tests {
		// For each test

		// Get the last process ids
		last_n_history_part := test.history.GetLastProcessIds(test.n)

		if last_n_history_part != test.expected {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.history.ProcessIds, last_n_history_part)
		}
	}
}

func TestHistoryClone(t *testing.T) {
	tests := []struct {
		history       History
		processes_ids []int
		expected      History
	}{
		{ // ----------------------- 0
			history: History{
				ProcessIds: "",
			},
			expected: History{
				ProcessIds: "",
			},
		},
		{ // ----------------------- 1
			history: History{
				ProcessIds: "3210",
			},
			expected: History{
				ProcessIds: "3210",
			},
		},
		{ // ----------------------- 2
			history: History{
				ProcessIds: "333333333333333333333",
			},
			expected: History{
				ProcessIds: "333333333333333333333",
			},
		},
	}

	for test_index, test := range tests {
		// For each test

		// Clone the history
		clone := test.history.Clone()

		// ------------------------ reset
		test.history.Reset()
		if clone.ProcessIds != test.expected.ProcessIds {
			t.Fatalf(`test %d (reset): expected = %s, got = %s`, test_index, test.expected.ProcessIds, clone.ProcessIds)
		}

		// ------------------------ push
		test.history.PushProcessId(1)
		test.history.PushProcessId(2)
		test.history.PushProcessId(3)

		if clone.ProcessIds != test.expected.ProcessIds {
			t.Fatalf(`test %d (push): expected = %s, got = %s`, test_index, test.expected.ProcessIds, clone.ProcessIds)
		}
	}
}
