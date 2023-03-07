package population

import (
	"math"
	"math/rand"
	"sort"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
	"github.com/trixky/krpsim/algo/simulation"
)

type ScoredInstance struct {
	Instance   instance.Instance
	Simulation simulation.Simulation
	Score      int
	Cycle      int
}

type ScoredPopulation struct {
	Instances []ScoredInstance
}

type Population struct {
	Instances []instance.Instance
}

func NewPopulation(options core.Options) Population {
	// TODO generate from InitialContext to have the correct amount of genes and other things ?
	return Population{
		Instances: make([]instance.Instance, options.PopulationSize),
	}
}

func NewRandomPopulation(context core.InitialContext, options core.Options) Population {
	instances := make([]instance.Instance, options.PopulationSize)
	for i := range instances {
		instances[i].Init(context)
	}
	return Population{
		Instances: instances,
	}
}

func (p *Population) RunAllSimulations(context core.InitialContext, options core.Options) ScoredPopulation {
	var scored []ScoredInstance

	// Run a simulation on all instances
	for _, instance := range p.Instances {
		simulation := simulation.NewSimulation(context, instance)
		simulation.Run(options.MaxCycle)
		scored = append(scored, ScoredInstance{
			Instance:   instance,
			Simulation: simulation,
			Score:      simulation.CalulateFitness(),
			Cycle:      simulation.Cycle,
		})
	}

	// Sort by Score
	// TODO Improve and sort by fitness on each optimize fields instead ?
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score < scored[j].Score
	})

	return ScoredPopulation{
		Instances: scored,
	}
}

func (p *ScoredPopulation) Best() ScoredInstance {
	return p.Instances[0]
}

func (s *ScoredPopulation) Crossover(options core.Options) Population {
	population := NewPopulation(options)

	// Calculate the total score of all instances
	total := 1.
	for _, scoredInstance := range s.Instances {
		total += float64(scoredInstance.Score)
	}

	// TODO Elitism: Select the best and keep them
	// TODO New Instances: Keep a few open slots for totally new Instances

	for i, scoredInstance := range s.Instances {
		// Select the other instance to cross with
		// Add the chance of both instances and roll a dice
		var with *ScoredInstance
		for j, otherScoredInstance := range s.Instances {
			if i != j {
				chance := float64(scoredInstance.Score)/total + float64(otherScoredInstance.Score)/total
				if rand.Float64() <= chance {
					with = &otherScoredInstance
					break
				}
			}
		}
		// TODO Generate a new instance if none was selected (instead of selecting the first one)
		if with == nil {
			with = &s.Instances[0]
		}
		child1, child2 := scoredInstance.Instance.Cross(&with.Instance)
		population.Instances = append(population.Instances, child1, child2)
	}

	return population
}

func (p *Population) Mutate(options core.Options) {
	for _, instance := range p.Instances {
		instance.Chromosome.Mutate(math.MaxUint16 / 2) // TODO pass options
	}
}