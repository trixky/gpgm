package history

import "strings"

type History struct {
	ProcessesIds string
}

func (h *History) Reset() {
	h.ProcessesIds = ""
}

func (h *History) PushProcessId(process_id int) {
	h.ProcessesIds = string(rune(process_id+48)) + h.ProcessesIds
}

func (h *History) GetLastProcessIds(n int) string {
	if n == 0 || len(h.ProcessesIds) < n {
		return h.ProcessesIds
	}

	return h.ProcessesIds[:n]
}

func (h *History) Clone() History {
	return History{
		ProcessesIds: strings.Clone(h.ProcessesIds),
	}
}
