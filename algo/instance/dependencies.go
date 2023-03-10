package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type ProcessDependencies struct {
	Processes []int // [inputs (in order)][processes (in order | can be skipped)]
}
type InputDependencies struct {
	Inputs []ProcessDependencies // [inputs (in order)][processes (in order | can be skipped)]
}

// Cut remove processes randomly when is possible
func (pd *ProcessDependencies) Cut(luck int) {
	for len(pd.Processes) > 1 && rand.Intn(luck) == 0 {
		pd.Processes = pd.Processes[1:]
	}
}

// Init initalizes the processes dependencies for an specific input
func (pd *ProcessDependencies) Init(input string, processes []core.Process) {
	for parent_process_index, parent_process := range processes {
		// For each potential parent process
		for output, output_quantity := range parent_process.Outputs {
			// For each output of the potential parent process
			if output == input {
				// Confirm if the potential parent process have the output corresponding to the process input
				if output_process_input_quantity, ok := parent_process.Inputs[input]; !ok || output_quantity > output_process_input_quantity {
					// If its X output is greater than its input if it as an input

					// Random insertion
					if rand.Intn(2) == 0 {
						// Insert as first
						pd.Processes = append([]int{parent_process_index}, pd.Processes...)
					} else {
						// Insert as last
						pd.Processes = append(pd.Processes, parent_process_index)
					}
				}
			}
		}
	}

	pd.Cut(3) // HARDCODED
}

// Init initalizes the input dependencies for an specific process
func (id *InputDependencies) Init(process core.Process, processes []core.Process) {
	id.Inputs = []ProcessDependencies{}

	for input := range process.Inputs {
		// For each input of the process

		// Initalizes its process dependencies
		process_dependencies := ProcessDependencies{}
		process_dependencies.Init(input, processes)

		// Random insertion
		if rand.Intn(2) == 0 {
			// Insert as first
			id.Inputs = append([]ProcessDependencies{process_dependencies}, id.Inputs...)
		} else {
			// Insert as last
			id.Inputs = append(id.Inputs, process_dependencies)
		}
	}
}
