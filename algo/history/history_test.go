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
		{
			history: History{},
			expected: History{
				ProcessesIds: "",
			},
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			expected: History{},
		},
		{
			history: History{
				ProcessesIds: "BA@?>=<;:9876543210",
			},
			expected: History{},
		},
	}

	for test_index, test := range tests {
		test.history.Reset()
		if test.history.ProcessesIds != test.expected.ProcessesIds {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.expected.ProcessesIds, test.history.ProcessesIds)
		}
	}
}

func TestHistoryPushProcessId(t *testing.T) {
	tests := []struct {
		history       History
		processes_ids []int
		expected      History
	}{
		{
			history:       History{},
			processes_ids: []int{},
			expected: History{
				ProcessesIds: "",
			},
		},
		{
			history:       History{},
			processes_ids: []int{0, 1, 2, 3},
			expected: History{
				ProcessesIds: "3210",
			},
		},
		{
			history:       History{},
			processes_ids: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
			expected: History{
				ProcessesIds: "BA@?>=<;:9876543210",
			},
		},
	}

	for test_index, test := range tests {
		for _, process_id := range test.processes_ids {
			test.history.PushProcessId(process_id)
		}
		if test.history.ProcessesIds != test.expected.ProcessesIds {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.history.ProcessesIds, test.expected.ProcessesIds)
		}
	}
}

func TestHistoryGetLastProcessIds(t *testing.T) {
	tests := []struct {
		history  History
		n        int
		expected string
	}{
		{
			history:  History{},
			n:        0,
			expected: "",
		},
		{
			history:  History{},
			n:        1,
			expected: "",
		},
		{
			history:  History{},
			n:        2,
			expected: "",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        0,
			expected: "3210",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        1,
			expected: "3",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        2,
			expected: "32",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        3,
			expected: "321",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        4,
			expected: "3210",
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			n:        5,
			expected: "3210",
		},
	}

	for test_index, test := range tests {
		last_n_history_part := test.history.GetLastProcessIds(test.n)
		if last_n_history_part != test.expected {
			t.Fatalf(`test %d: expected = %s, got = %s`, test_index, test.history.ProcessesIds, last_n_history_part)
		}
	}
}

func TestHistoryClone(t *testing.T) {
	tests := []struct {
		history       History
		processes_ids []int
		expected      History
	}{
		{
			history: History{
				ProcessesIds: "",
			},
			expected: History{
				ProcessesIds: "",
			},
		},
		{
			history: History{
				ProcessesIds: "3210",
			},
			expected: History{
				ProcessesIds: "3210",
			},
		},
		{
			history: History{
				ProcessesIds: "333333333333333333333",
			},
			expected: History{
				ProcessesIds: "333333333333333333333",
			},
		},
	}

	for test_index, test := range tests {
		clone := test.history.Clone()

		// ------------------------ reset
		test.history.Reset()
		if clone.ProcessesIds != test.expected.ProcessesIds {
			t.Fatalf(`test %d (reset): expected = %s, got = %s`, test_index, test.expected.ProcessesIds, clone.ProcessesIds)
		}

		// ------------------------ push

		test.history.PushProcessId(1)
		test.history.PushProcessId(2)
		test.history.PushProcessId(3)

		if clone.ProcessesIds != test.expected.ProcessesIds {
			t.Fatalf(`test %d (push): expected = %s, got = %s`, test_index, test.expected.ProcessesIds, clone.ProcessesIds)
		}
	}
}