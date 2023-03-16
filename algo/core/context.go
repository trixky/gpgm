package core

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type InitialContext struct {
	Stock      Stock           `json:"stock"`
	Processes  []Process       `json:"processes"`
	Optimize   map[string]bool `json:"optimize"`
	ScoreRatio map[string]int  `json:"score_ratio"`
}

// HaveOutput check if it has an resource as output
func (sm *InitialContext) HaveOutput(resource string) bool {
	for _, process := range sm.Processes {
		// For each process
		if process.HaveOutput(resource) {
			// If the process have the output
			return true
		}
	}

	return false
}

// FindProcessParents find process parents the initial context
func (sm *InitialContext) FindProcessParents() {
	for child_index, child := range sm.Processes {
		// For each child process
		for parent_index, parent := range sm.Processes {
			// For each parent process
			// NOTE: The parent can be the child
			for resource_name := range parent.Inputs {
				// For each input resource of the parent
				if output, ok := child.Outputs[resource_name]; ok {
					// If the child has the X input resource of the parent as output
					if input, ok := child.Inputs[resource_name]; !ok || output > input {
						// If its X output is greater than its input if it as an input
						sm.Processes[child_index].Parents = append(sm.Processes[child_index].Parents, parent_index)
					}
				}
			}
		}
	}
}
