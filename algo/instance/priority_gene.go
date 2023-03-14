package instance

import (
	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/history"
)

type PriorityGene struct {
	HistoryProcessDependencies map[string]ProcessDependencies
	Process                    *core.Process
}

// InitHistory initalizes recursively the history
func (pg *PriorityGene) InitHistory(h *history.History, depth int, process *core.Process, processes []core.Process) {
	depth--

	if h == nil {
		h = &history.History{}
	}

	key := h.GetLastProcessIds(0)
	dependences := ProcessDependencies{}
	dependences.Init(*pg.Process, processes)
	pg.HistoryProcessDependencies[key] = dependences

	if depth > 0 {
		for _, process_parent := range process.Parents {
			// For each process parents
			h_clone := h.Clone()
			h_clone.PushProcessId(process_parent)

			pg.InitHistory(&h_clone, depth, &processes[process_parent], processes)
		}
	}
}

// Init initalizes the gene attributes
func (pg *PriorityGene) Init(process *core.Process, processes []core.Process, optimize map[string]bool, options *core.Options) {
	pg.HistoryProcessDependencies = map[string]ProcessDependencies{}
	pg.Process = process

	pg.InitHistory(nil, options.HistoryKeyMaxLength, process, processes)
}
