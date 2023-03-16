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

	for index := range processes {
		// For each process

		// Create its corresponding priority gene
		// and initialize it
		priority_gene := PriorityGene{}
		priority_gene.Init(&processes[index], processes, optimize, options)

		// Add it to the chromosome
		c.PriorityGenes[index] = priority_gene
	}
}

// Cross generates two childs by cross overing itself with another one
func (c *Chromosome) Cross(cc *Chromosome) (child_1 Chromosome, child_2 Chromosome) {
	// ------------- Entry gene
	// --- Child 1
	child_1.EntryGene = *c.EntryGene.DeepCopy()
	// --- Child 2
	child_2.EntryGene = *cc.EntryGene.DeepCopy()

	// ------------- priority genes
	priority_gene_nb := len(c.PriorityGenes)

	// Get the random cross over index
	priority_cross := rand.Intn(priority_gene_nb)

	// --- Child 1
	c_copy_1 := c.DeepCopy()
	cc_copy_1 := cc.DeepCopy()
	// extract the first genes of the first parent
	child_1.PriorityGenes = c_copy_1.PriorityGenes[:priority_cross]
	// extract the last genes of the last parent
	child_1.PriorityGenes = append(child_1.DeepCopy().PriorityGenes, cc_copy_1.PriorityGenes[priority_cross:]...)

	// --- Child 2
	c_copy_2 := c.DeepCopy()
	cc_copy_2 := cc.DeepCopy()
	// extract the last genes of the first parent
	child_2.PriorityGenes = cc_copy_2.PriorityGenes[:priority_cross]
	// extract the first genes of the last parent
	child_2.PriorityGenes = append(child_2.DeepCopy().PriorityGenes, c_copy_2.PriorityGenes[priority_cross:]...)

	return
}

// Mutate generates a child by mutation
func (c *Chromosome) Mutate(processes []core.Process, optimize map[string]bool, options *core.Options) *Chromosome {
	mutated_chromosome := Chromosome{}

	new_chromosome := Chromosome{}
	new_chromosome.Init(processes, optimize, options)

	c_copy := c.DeepCopy()

	// ----------- Entry gene
	// Use the entry gene of the chromosome copy
	mutated_chromosome.EntryGene = c_copy.EntryGene
	// Mutate the entry gene using the new one
	mutated_chromosome.EntryGene = *mutated_chromosome.EntryGene.Mutate(&new_chromosome.EntryGene, options)

	// ----------- priority gene
	// Initializes the mutated priority genes
	mutated_chromosome.PriorityGenes = make([]PriorityGene, len(c_copy.PriorityGenes))

	for priority_gene_index, priority_gene := range c_copy.PriorityGenes {
		// For each priority gene of the chromosome copy
		// Mutate the priority gene using the new one copie's
		mutated_chromosome.PriorityGenes[priority_gene_index] = *priority_gene.Mutate(&new_chromosome.PriorityGenes[priority_gene_index], options)
	}

	return &mutated_chromosome
}

// DeepCopy make a deep copy of itself
func (c *Chromosome) DeepCopy() *Chromosome {
	deep_copy := Chromosome{}

	// ---------- Entry gene
	deep_copy.EntryGene = *c.EntryGene.DeepCopy()

	// ---------- priority genes
	deep_copy.PriorityGenes = make([]PriorityGene, len(c.PriorityGenes))

	for priority_gene_index, priority_gene := range c.PriorityGenes {
		deep_copy.PriorityGenes[priority_gene_index] = *priority_gene.DeepCopy()
	}

	return &deep_copy
}
