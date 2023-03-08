package instance

import (
	"math"
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type Instance struct {
	Chromosome Chromosome `json:"chromosome"`
}

// Init initializes an instance with random genes/exons
func (i *Instance) Init(initial_context core.InitialContext) {
	i.Chromosome.Genes = make([]PriorityGene, 0)
	for range initial_context.Processes {
		gene := PriorityGene{
			// first priority exon
			FirstPriorityExon: Exon{
				Value: uint16(rand.Intn(math.MaxUint16)),
			},
			// last priority exon
			LastPriorityExon: Exon{
				Value: uint16(rand.Intn(math.MaxUint16)),
			},
			// ratio exon
			RatioExons: make([]Exon, len(initial_context.Processes)),
		}

		// ratio exon
		for index := range gene.RatioExons {
			gene.RatioExons[index].Value = uint16(rand.Intn(math.MaxUint16))
		}

		i.Chromosome.Genes = append(i.Chromosome.Genes, gene)
	}
}

// Cross generates two childs by cross overing itself with another one
func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	chromosome_1, chromosome_2 := i.Chromosome.Cross(&ii.Chromosome)

	child_1.Chromosome = chromosome_1
	child_2.Chromosome = chromosome_2

	return
}
