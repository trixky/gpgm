package instance

type Instance struct {
	Genome Genome
}

func (i *Instance) Cross(ii *Instance) (child_1 Instance, child_2 Instance) {
	genome_1, genome_2 := i.Genome.Cross(&ii.Genome)
	child_1.Genome = genome_1
	child_2.Genome = genome_2
	return
}
