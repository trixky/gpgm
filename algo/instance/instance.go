package instance

import (
	"github.com/trixky/krpsim/algo/core"
)

type Instance struct {
	Chromosome Chromosome `json:"chromosome"`
}

// Init initializes its attributes randomly
func (i *Instance) Init(processes []core.Process, optimize map[string]bool, options *core.Options) {
	chromosome := Chromosome{}
	chromosome.Init(processes, optimize, options)
	i.Chromosome = chromosome
}

// Cross generates two childs by cross overing itself with another one
func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	chromosome_1, chromosome_2 := i.Chromosome.Cross(&ii.Chromosome)

	child_1.Chromosome = chromosome_1
	child_2.Chromosome = chromosome_2

	return
}

// Mutate make a mutated version of itself
func (i *Instance) Mutate(processes []core.Process, optimize map[string]bool, options *core.Options) *Instance {
	mutated_instance := Instance{}
	mutated_instance.Chromosome = *i.Chromosome.Mutate(processes, optimize, options)

	return &mutated_instance
}
