package simulation

import "github.com/trixky/krpsim/algo/core"

// OptimizeOnlyFitness optimize using only the fitness
func OptimizeOnlyFitness(simulation Simulation) int {
	score := 0
	factor := 1

	for name, forTime := range simulation.InitialContext.Optimize {
		if name == "time" {
			continue
		}

		quantity := simulation.Stock.GetResource(name)
		if forTime {
			score += (quantity / simulation.Cycle) * factor
		} else {
			score += quantity * factor
		}
		factor /= 2
	}

	return score
}

// Fitness a simulation
func Fitness(simulation Simulation, options *core.Options) int {
	return OptimizeOnlyFitness(simulation)
}
