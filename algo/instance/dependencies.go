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

// --------------------------- InputDependencies

// DeepCopy make a deep copy of itself
func (id *InputDependencies) DeepCopy() *InputDependencies {
	new_input_dependencies := InputDependencies{
		Input:               id.Input,
		ProcessDependencies: make([]int, len(id.ProcessDependencies)),
	}

	copy(new_input_dependencies.ProcessDependencies, id.ProcessDependencies)

	return &new_input_dependencies
}

// Cut remove processes randomly when is possible
func (id *InputDependencies) Cut(luck int) {
	for len(id.ProcessDependencies) > 1 && rand.Intn(luck) == 0 {
		id.ProcessDependencies = id.ProcessDependencies[1:]
	}
}

// Init initalizes the processes dependencies for an specific input
func (id *InputDependencies) Init(input string, processes []core.Process) {
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
						id.ProcessDependencies = append([]int{parent_process_index}, id.ProcessDependencies...)
					} else {
						// Insert as last
						id.ProcessDependencies = append(id.ProcessDependencies, parent_process_index)
					}
				}
			}
		}
	}
}

// --------------------------- ProcessDependencies

// DeepCopy make a deep copy of itself
func (pd *ProcessDependencies) DeepCopy() *ProcessDependencies {
	new_processd_ependencies := ProcessDependencies{
		InputDependencies: make([]InputDependencies, len(pd.InputDependencies)),
	}

	for index_input_dependencie, input_dependencie := range pd.InputDependencies {
		new_processd_ependencies.InputDependencies[index_input_dependencie] = *input_dependencie.DeepCopy()
	}

	return &new_processd_ependencies
}

// Init initalizes the input dependencies for an specific process
func (pd *ProcessDependencies) Init(process core.Process, processes []core.Process) {
	pd.InputDependencies = []InputDependencies{}

	for input := range process.Inputs {
		// For each input of the process

		// Initalizes its process dependencies
		process_dependencies := InputDependencies{}
		process_dependencies.Init(input, processes)

		process_dependencies.Input = input

		// Random insertion
		if rand.Intn(2) == 0 {
			// Insert as first
			pd.InputDependencies = append([]InputDependencies{process_dependencies}, pd.InputDependencies...)
		} else {
			// Insert as last
			pd.InputDependencies = append(pd.InputDependencies, process_dependencies)
		}
	}
}
