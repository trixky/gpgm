package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type EntryGene struct {
	Process_ids []int
}

func (eg *EntryGene) DeepCopy() *EntryGene {
	return &EntryGene{
		Process_ids: append(make([]int, 0, len(eg.Process_ids)), eg.Process_ids...),
	}
}

// RandomCut cut processes randomly when is possible
func (eg *EntryGene) RandomCut(luck int) {
	// WARNING: luck at one cut all processes except the last one

	if luck > 0 {
		// If the luck is stricly positive
		for i := len(eg.Process_ids); i > 1; i-- {
			// For each process except the last one

			if rand.Intn(luck) == 0 {
				// Cut the first process id randomly regarding the luck
				eg.Process_ids = eg.Process_ids[1:]
			}
		}
	}
}

// CutN cut n processes
func (eg *EntryGene) CutN(n uint) {
	if n > 0 && n < uint(len(eg.Process_ids)) {
		// If n is stricly positive
		// Cut the process ids
		eg.Process_ids = eg.Process_ids[:n]
	}
}

// CutN cut a random n processes
func (eg *EntryGene) CutRandomN() {
	if length := len(eg.Process_ids); length > 0 {
		// If at least one process id is present
		// Get the random n
		n := rand.Intn(length) + 1
		// Cut the process ids
		eg.Process_ids = eg.Process_ids[:n]
	}
}

// InitProcesses initalizes the processes
func (eg *EntryGene) InitProcesses(processes []core.Process, optimize map[string]bool) {
	for resource := range optimize { // "time" is not used
		// For each resource to optimize
		for process_index, process := range processes {
			// For each process
			for output := range process.Outputs {
				// For each output of the process
				if output == resource {
					// If the output corresponds to the resource
					if rand.Intn(2) == 0 { // Randomization
						// Add it at the beginning
						eg.Process_ids = append([]int{process_index}, eg.Process_ids...)
					} else {
						// Add it at the end
						eg.Process_ids = append(eg.Process_ids, process_index)
					}
				}
			}
		}
	}
}

// Init initalizes the processes
func (eg *EntryGene) Init(processes []core.Process, optimize map[string]bool, options *core.Options) {
	// Initializes processes
	eg.InitProcesses(processes, optimize)

	if options.RandomCut {
		// If random option is active
		// Cut randomly processes
		eg.CutRandomN()
	}

	// Cut processes regarding max
	// Remember is ignored if equal 0
	eg.CutN(uint(options.MaxCut))
}
