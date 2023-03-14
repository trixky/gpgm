package solver

import (
	"math"
	"time"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/population"
)

type RunningSolver struct {
	Population population.Population `json:"population"`
	Context    core.InitialContext   `json:"context"`
	Options    core.Options          `json:"options"`
	Generation int                   `json:"generation"`
	Start      time.Time             `json:"start"`
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
	scored := solver.Population.RunAllSimulations(solver.Context, &solver.Options)
	crossover_population := scored.Crossover(&solver.Context, &solver.Options)
	solver.ComputeMutationRate() // Compute the mutation rate for the current generation
	mutated_population := crossover_population.Mutate(solver.Context, &solver.Options)
	solver.Population = population.Population{}
	solver.Population = *mutated_population
	solver.Generation += 1

	return scored
}
