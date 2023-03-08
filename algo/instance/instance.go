package instance

import (
	"math"
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type Instance struct {
	Chromosome Chromosome `json:"chromosome"`
}

func (i *Instance) Init(initial_context core.InitialContext) {
	// priority

	i.Chromosome.Genes = make([]Gene, 0)
	for range initial_context.Processes {
		i.Chromosome.Genes = append(i.Chromosome.Genes, Gene{
			Value: uint16(rand.Intn(math.MaxUint16)),
		})
	}
}

func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	chromosome_1, chromosome_2 := i.Chromosome.Cross(&ii.Chromosome)
	child_1.Chromosome = chromosome_1
	child_2.Chromosome = chromosome_2
	return
}
