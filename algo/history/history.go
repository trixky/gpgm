package history

import "strings"

type History struct {
	ProcessesIds string
}

func reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func (h *History) Reset() {
	h.ProcessesIds = ""
}

func (h *History) PushProcessId(process_id int) {
	h.ProcessesIds = string(rune(process_id+48)) + h.ProcessesIds
}

func (h *History) InvertedPushProcessId(process_id int) {
	h.ProcessesIds = h.ProcessesIds + string(rune(process_id+48))
}

func (h *History) GetLastProcessIds(n int) string {
	if n == 0 || len(h.ProcessesIds) < n {
		return h.ProcessesIds
		// return reverse(h.ProcessesIds)
	}

	return h.ProcessesIds[:n]
	// return reverse(h.ProcessesIds[:n])
}

func (h *History) Clone() History {
	return History{
		ProcessesIds: strings.Clone(h.ProcessesIds),
	}
}
