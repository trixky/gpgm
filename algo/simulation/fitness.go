package simulation

import (
	"github.com/trixky/krpsim/algo/core"
)

// OptimizeOnlyFitness optimize using only the fitness
func OptimizeOnlyFitness(simulation Simulation) float64 {
	score := 0.
	factor := 1.
	global_time := false

	for name := range simulation.InitialContext.Optimize {
		if name == "time" {
			global_time = true
			break
		}
	}

	for name, forTime := range simulation.InitialContext.Optimize {
		if name == "time" {
			continue
		}

		quantity := float64(simulation.Stock.GetResource(name))

		if global_time || forTime {
			score += (quantity / float64(simulation.Cycle)) * factor
		} else {
			score += quantity * factor
		}
		factor /= 2
	}

	return score
}

// Fitness a simulation
func Fitness(simulation Simulation, options *core.Options) float64 {
	return OptimizeOnlyFitness(simulation)
}
