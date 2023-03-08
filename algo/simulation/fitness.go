package simulation

func Fitness(simulation Simulation) int {
	score := 0
	factor := 1

	for name, forTime := range simulation.InitialContext.Optimize {
		if name == "time" {
			continue
		}

		quantity := simulation.Stock.Get(name)
		if forTime {
			score += (quantity / simulation.Cycle) * factor
		} else {
			score += quantity * factor
		}
		factor /= 2
	}

	return score
}
