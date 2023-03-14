package solver

import (
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

// * Run a single generation for the given solver.RunningSolver
func (solver *RunningSolver) RunGeneration() population.ScoredPopulation {
	scored := solver.Population.RunAllSimulations(solver.Context, &solver.Options)
	crossover_population := scored.Crossover(&solver.Context, &solver.Options)
	mutated_population := crossover_population.Mutate(solver.Context, &solver.Options, solver.Options.MutationChance)
	solver.Population = population.Population{}
	solver.Population = *mutated_population
	solver.Generation += 1

	return scored
}
