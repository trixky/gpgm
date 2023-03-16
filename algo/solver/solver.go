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

// ComputeMutationRate compute the mutation rate of its population
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
		solver.Options.MutationChance = (logChance - 1.) / (float64(solver.Options.MaxGeneration) - 1.)
	}
}

// RunGeneration run a single generation
func (solver *RunningSolver) RunGeneration() population.ScoredPopulation {
	scored := solver.Population.RunAllSimulations(solver.Context, &solver.Options, &solver.Timer)

	if solver.Options.MaxGeneration > 1 {
		// // ---------------- Crossover
		crossover_population := scored.Crossover(&solver.Context, &solver.Options)
		solver.Population = crossover_population

		// ---------------- Mutate
		solver.ComputeMutationRate() // Compute the mutation rate for the current generation
		mutated_population := solver.Population.Mutate(solver.Context, &solver.Options)
		solver.Population = *mutated_population
	}

	solver.Generation += 1

	return scored
}
