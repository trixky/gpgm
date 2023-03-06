package instance

type Genome struct {
	Chromosomes []Chromosome
}

// Cross generates two childs by cross overing itself with another one
func (c *Genome) Cross(cc *Genome) (child_1 Genome, child_2 Genome) {
	chromosome_nb := len(c.Chromosomes)

	child_1.Chromosomes = make([]Chromosome, chromosome_nb)
	child_2.Chromosomes = make([]Chromosome, chromosome_nb)

	for index, chromosome := range c.Chromosomes {
		child_1.Chromosomes[index], child_2.Chromosomes[index] = chromosome.Cross(&cc.Chromosomes[index])
	}

	return
}

// Mutate generates a child by mutation
func (c *Genome) Mutate(max int) (child Genome) {
	for index, chromosome := range c.Chromosomes {
		child.Chromosomes[index] = chromosome.Mutate(max)
	}

	return
}
