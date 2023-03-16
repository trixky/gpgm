package interpretor

import (
	"math/bits"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/history"
	"github.com/trixky/krpsim/algo/instance"
)

// TryExecuteMProcess try to execute n time a process and generate a process quantities stack including its processes dependencies
func TryExecuteMProcess(history *history.History, process_id int, i *instance.Instance, stock *core.Stock, processes []core.Process, n int, depth int, options *core.Options) (process_quantities_stack *ProcessQuantities, x int, complete bool) {
	// Update the depth
	depth++

	process_quantities_stack = &ProcessQuantities{}

	// Get the current process
	process := processes[process_id]

	// Try to execute n time the process
	if xx := process.TryExecuteN(stock, n); xx > 0 {
		n -= xx // Decrement n
		x += xx // Increment x

		// Add them to the process stack
		process_quantities_stack.Push(&ProcessQuantity{
			Process:  &process,
			Quantity: xx,
		})
	}
	if n > 0 {
		// If process executions are missing
		if depth < options.MaxDepth {
			// If the depth is not exceeded
			complete = true

			// Get the last process ids key from the history
			last_history_part := history.GetLastProcessIds(options.HistoryPartMaxLength)

			// Get the input dependencies of the process
			input_dependencies := i.Chromosome.PriorityGenes[process_id].HistoryProcessDependencies[last_history_part].InputDependencies

			// Update the history for process dependencies
			history.InvertedPushProcessId(process_id)
			for _, input_dependencie := range input_dependencies {
				// For each input dependencie of the process

				// Get the name of the input dependencie
				input_name := input_dependencie.Input

				process_dependencie_ids := input_dependencie.ProcessDependencies

				if len(process_dependencie_ids) > 0 {
					for _, process_dependencie_id := range process_dependencie_ids {
						// For each process dependencie of the input dependencie of the process

						input_available := stock.Get(input_name)
						input_needed := process.Inputs[input_name] * n

						if input_needed > input_available {
							// If input dependencies are needed
							// Get the process dependencie
							process_dependencie := processes[process_dependencie_id]
							// Compute the wanted input
							input_wanted := input_needed - input_available

							// Compute the number of process dependencie execution needed
							nn, nn_rest := bits.Div(0, uint(input_wanted), uint(process_dependencie.Outputs[input_name]))
							if nn_rest > 0 {
								nn++
							}

							// Clone the history for the process dependencies
							history_clone := history.Clone()

							// Execute the process dependencies recursively
							dependencie_process_quantities_stack, xx, _ := TryExecuteMProcess(&history_clone, process_dependencie_id, i, stock, processes, int(nn), depth, options)

							if xx != int(nn) {
								// If the process dependencie executions are not complete
								complete = false
							}

							// Add executed processes to the process quantities stack
							process_quantities_stack.Concatenate(*dependencie_process_quantities_stack)
						}
					}
				} else {
					complete = false
				}
			}
		}
	} else {
		// Else
		// All the executions are completed
		// fmt.Println("^^^ 11")
		complete = true
	}
	// fmt.Println("^^^ 12")

	return
}

// Interpret generate a process quantities stack by interpreting the chromosome of an instance
func Interpret(i instance.Instance, initial_context core.InitialContext, stock *core.Stock, options *core.Options) (process_quantities_stack *ProcessQuantities) {
	// Initialize the process stack
	process_quantities_stack = &ProcessQuantities{}

	for _, entry_process_id := range i.Chromosome.EntryGene.Process_ids {
		// For each entry process id

		// Initializes the loop vairables
		complete := true
		var executed_processes *ProcessQuantities

		for complete {
			// While entry processes are completely executed
			complete = false

			// Execute the entry process
			executed_processes, _, complete = TryExecuteMProcess(&history.History{}, entry_process_id, &i, stock, initial_context.Processes, options.NEntry, 0, options)

			// Add executed processes to the process quantities stack
			process_quantities_stack.Concatenate(*executed_processes)
		}
	}

	return
}
