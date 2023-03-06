package genetic

import "math/rand"

type OperatorInfo struct {
	PopulationSize int
	MutationRate   float32
	KPoints        int
}

type Operator struct {
	PopulationSize     int
	Selection          []Individual
	NextPopulation     Population
	MutationRate       float32
	KPoints            int
	GenesPerIndividual int
	Genes              Genes
}

func (ope Operator) NextGeneration() Population {
	ope.NextPopulation = make(Population, 0, ope.PopulationSize)
	for len(ope.NextPopulation) < ope.PopulationSize {
		ope.NextPopulation = append(ope.NextPopulation, ope.kPointsCrossover()...)
	}
	ope.mutate()
	return ope.NextPopulation
}

func (ope Operator) kPointsCrossover() []Individual {
	// get parents
	i1 := rand.Intn(len(ope.Selection))
	i2 := rand.Intn(len(ope.Selection))
	for i1 == i2 {
		i2 = rand.Intn(len(ope.Selection))
	}
	p1 := ope.Selection[i1]
	p2 := ope.Selection[i2]

	// get k different parent genes
	var alreadySeen map[int]bool
	for i := 0; i < ope.KPoints; i++ {
		for {
			iGene := rand.Intn(ope.GenesPerIndividual)
			_, ok := alreadySeen[iGene]
			if ok {
				continue
			}
			alreadySeen[iGene] = true
		}
	}

	// make two children
	child1 := make(Individual, 0, ope.GenesPerIndividual)
	child2 := make(Individual, 0, ope.GenesPerIndividual)
	for i := 0; i < ope.GenesPerIndividual; i++ {
		_, ok := alreadySeen[i]
		if !ok {
			child1 = append(child1, p1[i])
			child2 = append(child2, p2[i])
			continue
		}
		child1 = append(child1, p2[i])
		child2 = append(child2, p1[i])
	}
	return []Individual{child1, child2}
}

func (ope Operator) mutate() {
	for i := 0; i < len(ope.NextPopulation); i++ {
		if rand.Float32() < ope.MutationRate {
			ope.NextPopulation[i][rand.Intn(ope.GenesPerIndividual)] = RandGene(ope.Genes)
		}
	}
}
