package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/history"
)

type PriorityGene struct {
	HistoryProcessDependencies map[string]ProcessDependencies
	Process                    *core.Process
}

// DeepCopy
func (pg *PriorityGene) DeepCopy() *PriorityGene {
	new_priority_gene := PriorityGene{
		Process:                    pg.Process,
		HistoryProcessDependencies: make(map[string]ProcessDependencies),
	}

	for history_process_dependencie_index, history_process_dependencie := range pg.HistoryProcessDependencies {
		new_priority_gene.HistoryProcessDependencies[history_process_dependencie_index] = *history_process_dependencie.DeepCopy()
	}

	return &new_priority_gene
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

	if depth >= 0 {
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

// Mutate mutates according to a pourcentage
func (pg *PriorityGene) Mutate(pgpg *PriorityGene, options *core.Options) *PriorityGene {
	new_priority_gene := PriorityGene{
		Process:                    pg.Process,
		HistoryProcessDependencies: make(map[string]ProcessDependencies),
	}

	for pgpg_history_process_dependencie_key, pgpg_history_process_dependencie := range pgpg.HistoryProcessDependencies {
		// For each history process dependencie
		if rand.Intn(1000) < int(options.MutationChance*1000) {
			// Take the value of the new priority gene
			new_priority_gene.HistoryProcessDependencies[pgpg_history_process_dependencie_key] = pgpg_history_process_dependencie
		} else {
			// Keep the value of the initial priority gene
			new_priority_gene.HistoryProcessDependencies[pgpg_history_process_dependencie_key] = pg.HistoryProcessDependencies[pgpg_history_process_dependencie_key]
		}
	}

	return &new_priority_gene
}
