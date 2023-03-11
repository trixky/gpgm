package interpretor

import (
	"math/bits"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/history"
	"github.com/trixky/krpsim/algo/instance"
)

type ProcessQuantities struct {
	Process *core.Process `json:"process"`
	Amount  int           `json:"amount"`
}

type ProcessQuantitiesStack struct {
	Stack []ProcessQuantities
}

func (pqs *ProcessQuantitiesStack) DeepCopy() *ProcessQuantitiesStack {
	copy := ProcessQuantitiesStack{
		Stack: make([]ProcessQuantities, len(pqs.Stack)),
	}

	for process_quantities_index, process_quantities := range pqs.Stack {
		copy.Stack[process_quantities_index] = ProcessQuantities{
			Process: process_quantities.Process,
			Amount:  process_quantities.Amount,
		}
	}
	return &copy
}

func (pqs *ProcessQuantitiesStack) Concatenate(second_pqs ProcessQuantitiesStack) {
	pqs.Stack = append(pqs.Stack, second_pqs.Stack...) // deep copy ?
}

func (pqs *ProcessQuantitiesStack) Push(process_quantities *ProcessQuantities) {
	pqs.Stack = append(pqs.Stack, *process_quantities)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func ExecuteProcess(history *history.History, process_id int, i *instance.Instance, stock *core.Stock, processes []core.Process, n int, deep int) (process_quantities_stack *ProcessQuantitiesStack, xx int, full bool) {
	process_quantities_stack = &ProcessQuantitiesStack{}
	n_remaining := n
	// Update history
	last_history_part := history.GetLastProcessIds(3) // HARDCODED

	indentation := ""

	for i := 0; i < deep; i++ {
		indentation += "\t"
	}

	// Get the current process
	process := processes[process_id]
	// fmt.Println(indentation, "execute ################ process:", process.Name, "deep:", deep, "n:", n)

	deep++

	// ------------------
	// Try to execute n time the process
	// fmt.Println("A n_remaining", n_remaining)
	// fmt.Println("------- stocke befoare after")
	// fmt.Println(indentation, stock)
	x := process.TryExecuteN(stock, n_remaining)

	if x > 0 {
		xx++
		process_quantities_stack.Push(&ProcessQuantities{
			Process: &process,
			Amount:  x,
		})
		// if process.Name == "sale" {
		// 	fmt.Println(indentation, "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! EXECUTE", x, "("+process.Name+")")
		// }
	}
	// fmt.Println(indentation, stock)

	n_remaining -= x

	if n_remaining == 0 {
		// If the all the executions are completed
		// Finish
		full = true
		return
	}
	// ------------------

	// -------- deep limit

	if deep > 5 { // HARDCODED
		// fmt.Println("ouii")
		// fmt.Println(indentation, "super fully")
		return
	}

	// fmt.Println("???", process_id, last_history_part)
	// for key := range i.Chromosome.Genes[process_id].HistoryProcessDependencies {
	// fmt.Println(key)
	// }

	last_history_part_reversed := Reverse(last_history_part)
	input_dependencies := i.Chromosome.Genes[process_id].HistoryProcessDependencies[last_history_part_reversed].InputDependencies

	// fmt.Println("@@@@@@@@@@@@@@@ last_history_part:", last_history_part_reversed)
	// for key := range i.Chromosome.Genes[process_id].HistoryProcessDependencies {
	// 	if key == last_history_part_reversed {
	// 		fmt.Println("____________________________________________________ oui", key)
	// 	} else {
	// 		fmt.Println("@@@@ non", key)
	// 	}
	// }

	// fmt.Println("ddd pro:", process.Name, "inputlen:", len(input_dependencies))
	// for _, input_dependencie := range input_dependencies {
	// 	fmt.Println(input_dependencie.Input)
	// }
	full = true

	for _, input_dependencie := range input_dependencies {
		// fmt.Println(indentation, "execute ####### input D", input_dependencie.Input)
		// For each input dependencie of the process
		input_name := input_dependencie.Input

		for _, process_dependencie_id := range input_dependencie.ProcessDependencies {
			// For each process dependencie of the input dependencie of the process

			// Get the process dependencie
			process_dependencie := processes[process_dependencie_id]
			// fmt.Println(indentation, "execute ## process D", process_dependencie.Name)

			// Compute the n of the process dependencies
			input_available := stock.Get(input_name)
			input_needed := process.Inputs[input_name] * n_remaining
			input_wanted := input_needed - input_available
			// fmt.Println("on veut input wanted:", input_wanted)
			// fmt.Println("A:", 0)
			// fmt.Println("B:", uint64(input_wanted))
			// fmt.Println("C:", uint64(process_dependencie.Outputs[input_name]))
			nn, nn_rest := bits.Div64(0, uint64(input_wanted), uint64(process_dependencie.Outputs[input_name]))
			// fmt.Println("nn:", nn, "n:", n)
			if nn_rest > 0 {
				nn++
			}

			// Execute the process dependencies recursively
			history_clone := history.Clone()
			history_clone.PushProcessId(process_id)

			dependencie_process_quantities_stack, nnn, _ := ExecuteProcess(&history_clone, process_dependencie_id, i, stock, processes, int(nn), deep)

			if nnn != int(nn) {
				full = false
				// fmt.Println(indentation, "not fully")
			} // else {
			// 	fmt.Println(indentation, "fully")
			// }

			process_quantities_stack.Concatenate(*dependencie_process_quantities_stack)

			// ------------------
			// Try to execute n time the process
			// fmt.Println("B n_remaining", n_remaining)

			// x := process.TryExecuteN(stock, n_remaining)
			// // fmt.Println(indentation, "on essaye", x)

			// if x > 0 {
			// 	xx++
			// 	process_quantities_stack.Push(&ProcessQuantities{
			// 		Process: &process,
			// 		Amount:  x,
			// 	})
			// 	// fmt.Println(indentation, "B execute", x, "("+process.Name+")", "remaining:", n_remaining)
			// }

			// n_remaining -= x

			// if n_remaining == 0 {
			// 	// If the all the executions are completed
			// 	// Finish
			// 	full = true

			// 	return
			// }
			// ------------------

		}
	}

	return

	// process_dependencies := i.Chromosome.Genes[process_id].History[history]
	// input depdencies
	// input depdencies process
	// new history
	// how many time ?
	// ExecuteProcess x time the process
}

// func SatisfyProcessInputsN(history string, process_id int, i *instance.Instance, stock *core.Stock, processes []core.Process, n int) {
// 	process := processes[process_id]

// 	for input, required := range process.Inputs {
// 		required *= n

// 		in_stock := stock.Get(input)

// 		if missing := required - in_stock; required < in_stock {
// 			for _, input_dependencies := range i.Chromosome.Genes[process_id].History[history].InputDependencies {
// 				if input_dependencies.Input == input {
// 					for _, parent_process_id := range input_dependencies.ProcessDependencies {
// 						parent_process := processes[parent_process_id]

// 						executions, executions_rest := bits.Div(0, uint(missing), uint(parent_process.Outputs[input]))

// 						executions += executions_rest

// 						if complete := parent_process.TryExecuteN(stock, int(executions)); complete {
// 							goto super_continue
// 						}
// 					}
// 				}
// 			}
// 		}
// 	super_continue:
// 	}
// }

func Interpret(i instance.Instance, initial_context core.InitialContext, stock core.Stock) (process_quantities_stack *ProcessQuantitiesStack) {
	stock_copy := stock
	process_quantities_stack = &ProcessQuantitiesStack{}

	ok := true

	idd := 0

	for ok {
		// fmt.Println("---------------------------- interpret loop")
		ok = false
		tutu, _, full := ExecuteProcess(&history.History{}, 0, &i, &stock_copy, initial_context.Processes, 1, 0) // HARDCODED x2
		process_quantities_stack.Concatenate(*tutu)
		if full {
			ok = true
		}

		idd++
		if idd > 100 {
			// fmt.Println("______________________________________________________________AAAOAUUUCH")
			return
			// 	os.Exit(0)
		}
	}
	return

	// tutu, _, _ := ExecuteProcess(&history.History{}, 0, &i, &stock_copy, initial_context.Processes, 1, 0) // HARDCODED x2
	// return tutu

	// history := ""

	// start := i.Chromosome.Genes[0].Process

	// i.Chromosome.Genes[0].Process.ExecuteN(&stock, int(max_execution_time))

	// for _, gene := range i.Chromosome.Genes {
	// 	max_execution_time := initial_context.Processes[gene.ProcessId].CanBeExecutedMaxXTimes(&stock)

	// 	if max_execution_time > 0 {
	// 		initial_context.Processes[gene.ProcessId].ExecuteN(&stock, 1)
	// 		process_quantities = append(process_quantities, ProcessQuantities{
	// 			Process: &initial_context.Processes[gene.ProcessId],
	// 			Amount:  1,
	// 		})
	// 	}
	// }

	// for _, gene := range i.Chromosome.Genes {
	// 	max_execution_time := initial_context.Processes[gene.ProcessId].CanBeExecutedMaxXTimes(&stock)

	// 	if max_execution_time > 0 {
	// 		if !gene.MinQuantityActive || max_execution_time >= gene.MinQuantity {
	// 			if !gene.MaxQuantityActive || max_execution_time <= gene.MaxQuantity {
	// 				initial_context.Processes[gene.ProcessId].ExecuteN(&stock, int(max_execution_time))

	// 				process_quantities = append(process_quantities, ProcessQuantities{
	// 					Process: &initial_context.Processes[gene.ProcessId],
	// 					Amount:  int(max_execution_time),
	// 				})
	// 			}
	// 		}
	// 	}
	// }

	return
}
