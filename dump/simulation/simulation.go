package simulation

import (
	"sort"

	"github.com/trixky/krpsim/algo/genetic"
	"github.com/trixky/krpsim/algo/parser"
)

type Simulation struct {
	Parameters         parser.SimulationParameters
	MutationRate       float32
	GenesPerIndividual int
	Storage            Storage
	PopulationSize     int
	Population         genetic.Population
	MaxGeneration      int
	KPoints            int
}

type SimulationContext struct {
}

func NewSimulation(info parser.SimulationParameters) {
	var simulation Simulation
	simulation.Storage = NewStorage(info.Stock)
	simulation.PopulationSize = 50
	simulation.GenesPerIndividual = 50
	simulation.Population = genetic.RandomPopulation(genetic.PopulationInfo{
		Size:               simulation.PopulationSize,
		GenesPerIndividual: simulation.GenesPerIndividual,
		Genes:              simulation.Parameters.Processes,
	})

}

func (s *Simulation) Start() genetic.Population {
	for i := 0; i < s.MaxGeneration; i++ {
		// no sort here because many selection types don't use it
		fitnesses := s.getFitnesses()

		// select
		// for now get best 2
		selection := s.getElites(fitnesses)

		// create operator
		operator := genetic.Operator{
			PopulationSize:     s.PopulationSize,
			Selection:          selection,
			MutationRate:       s.MutationRate,
			KPoints:            s.KPoints,
			GenesPerIndividual: s.GenesPerIndividual,
			Genes:              s.Parameters.Processes,
		}
		// make next generation
	}
}

func (s *Simulation) getElites(f Fitnesses) []genetic.Individual {
	sort.Slice(f, func(p, q int) bool {
		return f[p].TrueScore > f[q].TrueScore
	})
	return []genetic.Individual{f[0].Individual, f[1].Individual}
}

func (s *Simulation) getFitness(ind genetic.Individual) Fitness {
	f := Fitness{
		Individual: ind,
	}
	storage := s.Storage.DeepCopy()
	for _, p := range ind {
		if storage.ConsumeIfAvailable(p.Inputs) {
			storage.Store(p.Outputs)

			f.GoodGenes += 1
			f.Delay += p.Delay
			for _, v := range s.Parameters.Optimize {
				f.TrueScore += storage.Get(v)
				// need to do Delay stuff when parsing gives map[string]bool for Parameters.Optimize
			}
		} else {
			break
		}
	}
	return f
}

func (s *Simulation) getFitnesses() Fitnesses {
	f := make(Fitnesses, 0, s.PopulationSize)
	for i := 0; i < s.PopulationSize; i++ {
		f = append(f, s.getFitness(s.Population[i]))
	}
	return f
}
