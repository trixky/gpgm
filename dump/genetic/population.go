package genetic

import "github.com/trixky/krpsim/algo/parser"

type PopulationInfo struct {
	Size               int
	GenesPerIndividual int
	Genes              []parser.Process
}

type Population []Individual

func RandomPopulation(info PopulationInfo) Population {
	p := make([]Individual, info.Size)
	for i := 0; i < info.Size; i++ {
		p[i] = RandomIndividual(info.GenesPerIndividual, info.Genes)
	}
	return p
}

func (p Population) Size() int {
	return len(p)
}