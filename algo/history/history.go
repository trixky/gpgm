package history

import "strings"

type History struct {
	ProcessIds string
}

// reverse reverses a string
func reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

// Reset resets its process ids
func (h *History) Reset() {
	h.ProcessIds = ""
}

// PushProcessId push a process id at the end
func (h *History) PushProcessId(process_id int) {
	h.ProcessIds = string(rune(process_id+48)) + h.ProcessIds
}

// PushProcessId push a process id at the beginning
func (h *History) InvertedPushProcessId(process_id int) {
	h.ProcessIds = h.ProcessIds + string(rune(process_id+48))
}

// GetLastProcessIds gets the n last process ids
func (h *History) GetLastProcessIds(n int) string {
	if n == 0 || len(h.ProcessIds) < n {
		return h.ProcessIds
	}

	return h.ProcessIds[:n]
}

// Clone returns a copy
func (h *History) Clone() *History {
	return &History{
		ProcessIds: strings.Clone(h.ProcessIds),
	}
}
