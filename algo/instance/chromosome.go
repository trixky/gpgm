package instance

import "math/rand"

type Chromosome struct {
	Genes []Gene
}

// Cross generates two childs by cross overing itself with another one
func (c *Chromosome) Cross(cc *Chromosome) (child_1 Chromosome, child_2 Chromosome) {
	gene_nb := len(c.Genes)

	cross := rand.Intn(gene_nb)

	// child 1
	// extract the first genes of the first parent
	child_1.Genes = c.Genes[:cross]
	// extract the last genes of the last parent
	child_1.Genes = append(child_1.Genes, cc.Genes[cross:]...)

	// child 2
	// extract the last genes of the first parent
	child_2.Genes = c.Genes[cross:]
	// extract the first genes of the last parent
	child_2.Genes = append(child_2.Genes, cc.Genes[:cross]...)

	return
}

// Mutate generates a child by mutation
func (c *Chromosome) Mutate(max int) (child Chromosome) {
	gene_nb := len(c.Genes)

	child.Genes = make([]Gene, gene_nb)

	for index, gene := range c.Genes {
		// Extract the mutation of all genes from the parent ones
		child.Genes[index] = gene.Mutate(max)
	}

	return
}
