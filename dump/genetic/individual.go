package genetic

import (
	"github.com/trixky/krpsim/algo/parser"
)

type Individual Genes

func RandomIndividual(size int, p []parser.Process) Individual {
	g := make(Individual, 0, size)
	for i := 0; i < size; i++ {
		g = append(g, RandGene(p))
	}
	return g
}
