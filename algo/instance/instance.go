package instance

import (
	"math"
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type Instance struct {
	Chromosome Chromosome
}

func (i *Instance) Init(initial_context core.SimulationInitialContext) {
	// priority

	for _ = range initial_context.Processes {
		i.Chromosome.Genes = append(i.Chromosome.Genes, Gene{
			Value: uint16(rand.Intn(math.MaxUint16)),
		})
	}
}

func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	genome_1, genome_2 := i.Genome.Cross(&ii.Genome)
	child_1.Genome = genome_1
	child_2.Genome = genome_2
	return
}
