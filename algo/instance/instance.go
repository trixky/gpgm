package instance

import (
	"github.com/trixky/krpsim/algo/core"
)

type Instance struct {
	Chromosome Chromosome `json:"chromosome"`
}

func (i *Instance) Init(processes []core.Process, optimize map[string]bool, options *core.Options) {
	chromosome := Chromosome{}
	chromosome.Init(processes, optimize, options)
	i.Chromosome = chromosome
}

// // Init initializes an instance with random genes/exons
// func (i *Instance) Init(initial_context core.InitialContext) {
// 	const chromosome_length_multiplicator = 50

// 	i.Chromosome.Genes = make([]Gene, 0)

// 	for j := 0; j < chromosome_length_multiplicator; j++ {
// 		for range initial_context.Processes {
// 			// quantity_1 := uint16(rand.Intn(math.MaxUint16))
// 			// quantity_2 := uint16(rand.Intn(math.MaxUint16))
// 			quantity_1 := uint16(rand.Intn(1) + 1)
// 			quantity_2 := uint16(rand.Intn(2) + 1)

// 			gene := Gene{
// 				ProcessId:         uint16(rand.Intn(len(initial_context.Processes))),
// 				MinQuantityActive: rand.Intn(2) == 0,
// 				MaxQuantityActive: rand.Intn(2) == 0,
// 			}

// 			if quantity_1 < quantity_2 {
// 				gene.MinQuantity = quantity_1
// 				gene.MaxQuantity = quantity_2
// 			} else {
// 				gene.MinQuantity = quantity_2
// 				gene.MaxQuantity = quantity_1
// 			}

// 			i.Chromosome.Genes = append(i.Chromosome.Genes, gene)
// 		}
// 	}
// }

// Cross generates two childs by cross overing itself with another one
func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	chromosome_1, chromosome_2 := i.Chromosome.Cross(&ii.Chromosome)

	child_1.Chromosome = chromosome_1
	child_2.Chromosome = chromosome_2

	return
}
