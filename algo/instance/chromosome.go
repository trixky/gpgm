package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type Chromosome struct {
	EntryGene     EntryGene      `json:"entry_gene"`
	PriorityGenes []PriorityGene `json:"genes"`
}

func (c *Chromosome) Init(processes []core.Process, optimize map[string]bool) {
	const random = true // HARDCODED
	const max_entry = 0 // HARDCODED

	// Initializes the entry gene
	c.EntryGene = EntryGene{}

	c.EntryGene.Init(processes, optimize, max_entry, random)

	// Initializes the priority genes
	c.PriorityGenes = make([]PriorityGene, len(processes))

	for index, process := range processes {
		// For each process
		// Initializes the corresponding priority gene
		priority_gene := PriorityGene{}
		priority_gene.Init(&process, processes, optimize)
		c.PriorityGenes[index] = priority_gene
	}
}

// Cross generates two childs by cross overing itself with another one
func (c *Chromosome) Cross(cc *Chromosome) (child_1 Chromosome, child_2 Chromosome) {
	gene_nb := len(c.PriorityGenes)

	cross := rand.Intn(gene_nb)

	// child 1
	// extract the first genes of the first parent
	child_1.PriorityGenes = c.PriorityGenes[:cross]
	// extract the last genes of the last parent
	child_1.PriorityGenes = append(child_1.PriorityGenes, cc.PriorityGenes[cross:]...)

	// child 2
	// extract the last genes of the first parent
	child_2.PriorityGenes = c.PriorityGenes[cross:]
	// extract the first genes of the last parent
	child_2.PriorityGenes = append(child_2.PriorityGenes, cc.PriorityGenes[:cross]...)

	return
}

// Mutate generates a child by mutation
func (c *Chromosome) Mutate(process_max uint16, process_shift int, quantity_shift int, activation_chance int, processes []core.Process, optimize map[string]bool) (child Chromosome) {
	gene_nb := len(c.PriorityGenes)

	child.PriorityGenes = make([]PriorityGene, gene_nb)

	for index := range c.PriorityGenes {
		// Extract the mutation of all genes from the parent ones
		child.PriorityGenes[index].Mutate(process_max, process_shift, quantity_shift, activation_chance, processes, optimize)
	}

	return
}
