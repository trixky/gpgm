package population

import (
	"fmt"
	"math"
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

func NewPopulation(options *core.Options) Population {
	// TODO generate from InitialContext to have the correct amount of genes and other things ?
	return Population{
		Instances: make([]instance.Instance, options.PopulationSize),
	}
}

func NewRandomPopulation(context core.InitialContext, options *core.Options) Population {
	context.FindProcessParents()

	instances := make([]instance.Instance, options.PopulationSize)
	for i := range instances {
		instances[i].Init(context.Processes, context.Optimize, options)
	}
	return Population{
		Instances: instances,
	}
}

func (p *Population) RunAllSimulations(context core.InitialContext, options *core.Options) ScoredPopulation {
	var scored []ScoredInstance

	// Run a simulation on all instances
	for instance_index, instance := range p.Instances {
		fmt.Println("--------- instance", instance_index)
		simulation := simulation.NewSimulation(context, instance)
		simulation.Run(options)
		scored = append(scored, ScoredInstance{
			Instance:   instance,
			Simulation: simulation,
			Score:      simulation.CalulateFitness(options),
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

// https://en.wikipedia.org/wiki/Tournament_selection
func (s *ScoredPopulation) TournamentSelection(forIndex int, options *core.Options) ScoredInstance {
	// Select k instances from all Instances (tournament population size)
	k := len(s.Instances)
	instanceIndexes := make([]int, k)
	for i := 0; i < k; i++ {
		instanceIndexes[i] = i
	}
	rand.Shuffle(k, func(i, j int) { instanceIndexes[i], instanceIndexes[j] = instanceIndexes[j], instanceIndexes[i] })

	// Select individual from the population
	fmt.Printf("%v\n", instanceIndexes)
	stopAt := int(math.Min(float64(k), float64(options.TournamentSize)))
	for i := 0; i < stopAt; i++ {
		globalIndex := instanceIndexes[i]
		if globalIndex == forIndex {
			continue
		}
		chance := options.TournamentProbability
		if i >= 1 {
			chance = chance * math.Pow(1-options.TournamentProbability, float64(i))
		}
		if rand.Float64() <= chance {
			return s.Instances[globalIndex]
		}
	}

	if instanceIndexes[0] == forIndex {
		return s.Instances[instanceIndexes[1]]
	}
	return s.Instances[instanceIndexes[0]]
}

// Set the chance of an instance it's score percentage from the global score of the population
// Add both chances of the instances and roll a dice to use the current instance
func (s *ScoredPopulation) RandomSelection(forIndex int, options *core.Options) ScoredInstance {
	// Calculate the total score of all instances
	total := 1.
	for _, scoredInstance := range s.Instances {
		total += float64(scoredInstance.Score)
	}

	// Add the chance of both instances and roll a dice
	scoredInstance := s.Instances[forIndex]
	for _, otherScoredInstance := range s.Instances[forIndex+1:] {
		chance := float64(scoredInstance.Score)/total + float64(otherScoredInstance.Score)/total
		if rand.Float64() <= chance {
			return otherScoredInstance
		}
	}
	return s.Instances[0]
}

func (s *ScoredPopulation) Crossover(initialContext *core.InitialContext, options *core.Options) Population {
	population := NewPopulation(options)
	if len(population.Instances) == 1 {
		population.Instances[0] = s.Instances[0].Instance
		return population
	}

	// * Elitism
	i := 0
	if options.ElitismAmount > 0 {
		max := options.ElitismAmount
		if options.ElitismAmount >= options.PopulationSize {
			max = int(float64(options.PopulationSize) * 0.9)
		}
		for j := 0; j < max; j++ {
			population.Instances[j] = s.Instances[j].Instance
		}
		i = max
	}

	// * New instances
	if options.CrossoverNewInstances > 0 {
		max := i + options.CrossoverNewInstances
		if max >= options.PopulationSize {
			max = int(float64(options.PopulationSize-i) * 0.9)
		}
		for j := i; j < max; j++ {
			population.Instances[j].Init(initialContext.Processes, initialContext.Optimize, options)
		}
		i = max
	}

	// * Crossover between Instances
	for ; i < options.PopulationSize; i += 2 {
		scoredInstance := s.Instances[i]

		// * Selection
		crossWith := ScoredInstance{}
		switch options.SelectionMethod {
		default:
		case core.RandomSelection:
			crossWith = s.RandomSelection(i, options)
		case core.TournamentSelection:
			crossWith = s.TournamentSelection(i, options)
		}

		// * Genetic operator
		child1, child2 := scoredInstance.Instance.Cross(&crossWith.Instance)
		population.Instances[i] = child1
		if i+1 < options.PopulationSize {
			population.Instances[i+1] = child2
		}
	}

	return population
}

func (p *Population) Mutate(context core.InitialContext, options *core.Options) *Population {
	mutated_population := Population{}
	mutated_population.Instances = make([]instance.Instance, len(p.Instances))

	for instance_index, instance := range p.Instances {
		mutated_population.Instances[instance_index] = *instance.Mutate(context.Processes, context.Optimize, options)
		// mutated_population.Instances[instance_index] = instance
	}

	return &mutated_population
}
