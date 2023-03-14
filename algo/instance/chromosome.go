package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type Chromosome struct {
	EntryGene     EntryGene      `json:"entry_gene"`
	PriorityGenes []PriorityGene `json:"genes"`
}

func (c *Chromosome) Init(processes []core.Process, optimize map[string]bool, options *core.Options) {
	// Initializes the entry gene
	c.EntryGene = EntryGene{}

	c.EntryGene.Init(processes, optimize, options)

	// Initializes the priority genes
	c.PriorityGenes = make([]PriorityGene, len(processes))

	for index, process := range processes {
		// For each process
		// Initializes the corresponding priority gene
		priority_gene := PriorityGene{}
		priority_gene.Init(&process, processes, optimize, options)
		c.PriorityGenes[index] = priority_gene
	}
}

// Cross generates two childs by cross overing itself with another one
func (c *Chromosome) Cross(cc *Chromosome) (child_1 Chromosome, child_2 Chromosome) {
	// ------------- entry gene
	// child 1
	child_1.EntryGene = *c.EntryGene.DeepCopy()
	// child 2
	child_2.EntryGene = *cc.EntryGene.DeepCopy()

	// ------------- priority genes
	priority_gene_nb := len(c.PriorityGenes)

	priority_cross := rand.Intn(priority_gene_nb)

	// child 1
	// extract the first genes of the first parent
	child_1.PriorityGenes = c.PriorityGenes[:priority_cross]
	// extract the last genes of the last parent
	child_1.PriorityGenes = append(child_1.PriorityGenes, cc.PriorityGenes[priority_cross:]...)

	// child 2
	// extract the last genes of the first parent
	child_2.PriorityGenes = c.PriorityGenes[priority_cross:]
	// extract the first genes of the last parent
	child_2.PriorityGenes = append(child_2.PriorityGenes, cc.PriorityGenes[:priority_cross]...)

	return
}

// Mutate generates a child by mutation
func (c *Chromosome) Mutate(processes []core.Process, optimize map[string]bool, options *core.Options) *Chromosome {
	mutated_chromosome := Chromosome{}

	new_chromosome := Chromosome{}
	new_chromosome.Init(processes, optimize, options)

	// ----------- entry gene
	mutated_chromosome.EntryGene = *c.EntryGene.DeepCopy()
	mutated_chromosome.EntryGene = *mutated_chromosome.EntryGene.Mutate(&new_chromosome.EntryGene, percentage)

	// ----------- priority gene
	mutated_chromosome.PriorityGenes = new_chromosome.PriorityGenes

	return &mutated_chromosome
}
