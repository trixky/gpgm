package population

import (
	"math/rand"
	"sort"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
	"github.com/trixky/krpsim/algo/simulation"
)

type ScoredInstance struct {
	Instance   instance.Instance     `json:"instance"`
	Simulation simulation.Simulation `json:"simulation"`
	Score      int                   `json:"score"`
	Cycle      int                   `json:"cycle"`
}

type ScoredPopulation struct {
	Instances []ScoredInstance `json:"instances"`
}

type Population struct {
	Instances []instance.Instance `json:"instances"`
}

func NewPopulation(options core.Options) Population {
	// TODO generate from InitialContext to have the correct amount of genes and other things ?
	return Population{
		Instances: make([]instance.Instance, options.PopulationSize),
	}
}

func NewRandomPopulation(context core.InitialContext, options core.Options) Population {
	context.FindProcessParents()

	instances := make([]instance.Instance, options.PopulationSize)
	for i := range instances {
		instances[i].Init(context.Processes)
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
		return scored[i].Score > scored[j].Score
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

	i := 0
	if options.UseElitism {
		max := options.ElitismAmount
		if options.ElitismAmount >= options.PopulationSize {
			max = int(float64(options.PopulationSize) * 0.9)
		}
		for j := 0; j < max; j++ {
			population.Instances[j] = s.Instances[j].Instance
		}
		i = max
	}

	// TODO New Instances: Keep a few open slots for totally new Instances ?
	for ; i < options.PopulationSize; i += 2 {
		scoredInstance := s.Instances[i]

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
		if with == nil {
			with = &s.Instances[0]
		}
		child1, child2 := scoredInstance.Instance.Cross(&with.Instance)
		population.Instances[i] = child1
		if i+1 < options.PopulationSize {
			population.Instances[i+1] = child2
		}
	}

	return population
}

func (p *Population) Mutate(context core.InitialContext, options core.Options) {
	for _, instance := range p.Instances {

		process_max := uint16(len(context.Processes))
		process_shift := 1
		quantity_shift := 1
		activation_chance := 10

		instance.Chromosome.Mutate(process_max, process_shift, quantity_shift, activation_chance) // TODO pass options
	}
}
