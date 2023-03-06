package genetic

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/parser"
)

type Genes []parser.Process

func RandGene(p []parser.Process) parser.Process {
	return p[rand.Int()%len(p)]
}

func (genes Genes) Push(g parser.Process) {
	genes = append(genes, g)
}