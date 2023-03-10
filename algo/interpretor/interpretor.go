package interpretor

import (
	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
)

type ProcessQuantities struct {
	Process *core.Process `json:"process"`
	Amount  int           `json:"amount"`
}

// func GetResources(process *core.Process, i *instance.Instance, stock core.Stock) {
// 	for resource_index, resource := range process.Inputs {

// 	}
// }

func Interpret(i instance.Instance, initial_context core.InitialContext, stock core.Stock) (process_quantities []ProcessQuantities) {
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
