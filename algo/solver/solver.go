package solver

import (
	"math"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/population"
	"github.com/trixky/krpsim/algo/timer"
)

type RunningSolver struct {
	Population population.Population `json:"population"`
	Context    core.InitialContext   `json:"context"`
	Options    core.Options          `json:"options"`
	Generation int                   `json:"generation"`
	Timer      timer.Timer           `json:"timer"`
}

// InitTimer initializes its timer
func (solver *RunningSolver) InitTimer() {
	solver.Timer.Init(int64(solver.Options.TimeLimitMS))
}

func (solver *RunningSolver) ComputeMutationRate() {
	switch solver.Options.MutationMethod {
	case core.LinearMutation:
		solver.Options.MutationChance = float64(1 - (solver.Generation / solver.Options.MaxGeneration))
	// https://stackoverflow.com/a/58501336
	case core.LogarithmicMutation:
		steps := float64(solver.Options.MaxGeneration)
		step := (math.Log(steps) - math.Log(1)) / (steps - 1)
		i := float64(solver.Options.MaxGeneration - solver.Generation)
		logChance := math.Exp(math.Log(1) + i*step)
		// Convert to range [0-1] (x - min(x)) / (max(x) - min(x))
		solver.Options.MutationChance = (logChance - 1.) / (float64(solver.Options.MaxGeneration) - 1.)
		// case core.ExponentialMutation:
		// 	steps := float64(solver.Options.MaxGeneration)
		// 	step := (math.Exp(steps) - math.Exp(1)) / (steps - 1)
		// 	i := float64(solver.Options.MaxGeneration - solver.Generation)
		// 	expChance := math.Log(math.Exp(1) + i*step)
		// 	// Convert to range [0-1] (x - min(x)) / (max(x) - min(x))
		// 	solver.Options.MutationChance = math.Max((expChance-1.)/(float64(solver.Options.MaxGeneration)-1.), 0)
	}
}

// * Run a single generation
func (solver *RunningSolver) RunGeneration() population.ScoredPopulation {
	// fmt.Println("---------------------------- SOLVER")
	scored := solver.Population.RunAllSimulations(solver.Context, &solver.Options, &solver.Timer)

	if solver.Options.MaxGeneration > 1 {
		// // ---------------- crossover
		// fmt.Println("--------------- AVANT ~~~~~~~~~~~~~~~~~~ avant")
		// fmt.Println("pop len:", len(solver.Population.Instances))
		// fmt.Println("pop 0 len entry:", len(solver.Population.Instances[0].Chromosome.EntryGene.Process_ids))
		// fmt.Println("pop 0 len len:", len(solver.Population.Instances[0].Chromosome.PriorityGenes))

		crossover_population := scored.Crossover(&solver.Context, &solver.Options)
		solver.Population = crossover_population

		// fmt.Println("--------------- APRES ~~~~~~~~~~~~~~~~~~ avant")
		// fmt.Println("pop len:", len(solver.Population.Instances))
		// fmt.Println("pop 0 len entry:", len(solver.Population.Instances[1].Chromosome.EntryGene.Process_ids))
		// fmt.Println("pop 0 len priority:", len(solver.Population.Instances[1].Chromosome.PriorityGenes))
		// fmt.Println("pop 0 len priority 0 name:", solver.Population.Instances[1].Chromosome.PriorityGenes[0].Process.Name)
		// fmt.Println("pop 0 len priority 0 dep:", len(solver.Population.Instances[1].Chromosome.PriorityGenes[0].HistoryProcessDependencies))
		// fmt.Println("pop 0 len priority 00 dep:", len(solver.Population.Instances[1].Chromosome.PriorityGenes[0].HistoryProcessDependencies[""].InputDependencies))
		// fmt.Println("pop 0 len priority 000 dep:", solver.Population.Instances[1].Chromosome.PriorityGenes[0].HistoryProcessDependencies[""].InputDependencies[0])

		// ---------------- mutate
		solver.ComputeMutationRate() // Compute the mutation rate for the current generation
		mutated_population := solver.Population.Mutate(solver.Context, &solver.Options)
		solver.Population = *mutated_population
	}

	solver.Generation += 1

	return scored
}
