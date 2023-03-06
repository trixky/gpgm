package simulation

import "github.com/trixky/krpsim/algo/genetic"

type Fitness struct {
	Individual genetic.Individual
	GoodGenes  int
	Score      int
	TrueScore  int
	Delay      int
}

type Fitnesses []Fitness
