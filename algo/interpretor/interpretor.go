package interpretor

import (
	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
)

func InterpretBasicPriority(i instance.Instance, initial_context core.InitialContext, stock core.Stock) (processes []core.Process) {
	processes_order := []core.Process{}

	processes_cpy := make([]core.Process, len(initial_context.Processes))
	genes_cpy := make([]instance.Gene, len(initial_context.Processes))

	// can be optimized
	copy(processes_cpy, initial_context.Processes)
	copy(genes_cpy, i.Chromosome.Genes)

	for len(processes_cpy) > 0 {
		var best_value uint16 = 0
		var best_index int = 0

		for index, gene := range genes_cpy {
			if gene.FirstPriorityExon.Value > best_value {
				best_value = gene.FirstPriorityExon.Value
				best_index = index
			}
		}

		processes_cpy = append(processes_cpy[:best_index], processes_cpy[best_index+1:]...)
		genes_cpy = append(genes_cpy[:best_index], genes_cpy[best_index+1:]...)
		processes_order = append(processes_order, initial_context.Processes[best_index])
	}

	execution := true

	for execution {
		execution = false
		for _, process := range processes_order {
			if ok := process.Execute(&stock); ok {
				execution = true
				processes = append(processes, process)
			}
		}
	}

	return
}
