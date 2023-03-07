package interpretor

import (
	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
)

func Interpret(i instance.Instance, initial_context core.InitialContext, stock core.Stock) (processes []core.Process) {
	i_genes_cpy := make([]instance.Gene, len(i.Chromosome.Genes))
	processes_order := []core.Process{}
	copy(i_genes_cpy, i.Chromosome.Genes)

	for len(i_genes_cpy) > 0 {
		var best_value uint16 = 0
		var best_index int = 0

		for index, gene := range i_genes_cpy {
			if gene.Value > best_value {
				best_value = gene.Value
				best_index = index
			}
		}

		i_genes_cpy = append(i_genes_cpy[:best_index], i_genes_cpy[best_index+1:]...)
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
