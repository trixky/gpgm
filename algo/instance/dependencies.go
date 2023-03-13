package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type InputDependencies struct {
	Input               string
	ProcessDependencies []int
}
type ProcessDependencies struct {
	InputDependencies []InputDependencies
}

// Cut remove processes randomly when is possible
func (pd *InputDependencies) Cut(luck int) {
	for len(pd.ProcessDependencies) > 1 && rand.Intn(luck) == 0 {
		pd.ProcessDependencies = pd.ProcessDependencies[1:]
	}
}

// Init initalizes the processes dependencies for an specific input
func (pd *InputDependencies) Init(input string, processes []core.Process) {
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
						pd.ProcessDependencies = append([]int{parent_process_index}, pd.ProcessDependencies...)
					} else {
						// Insert as last
						pd.ProcessDependencies = append(pd.ProcessDependencies, parent_process_index)
					}
				}
			}
		}
	}
}

// Init initalizes the input dependencies for an specific process
func (id *ProcessDependencies) Init(process core.Process, processes []core.Process) {
	id.InputDependencies = []InputDependencies{}

	for input := range process.Inputs {
		// For each input of the process

		// Initalizes its process dependencies
		process_dependencies := InputDependencies{}
		process_dependencies.Init(input, processes)

		process_dependencies.Input = input

		// Random insertion
		if rand.Intn(2) == 0 {
			// Insert as first
			id.InputDependencies = append([]InputDependencies{process_dependencies}, id.InputDependencies...)
		} else {
			// Insert as last
			id.InputDependencies = append(id.InputDependencies, process_dependencies)
		}
	}
}
