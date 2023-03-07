package simulation

func Fitnesss(simulation Simulation) int {
	score := 0
	factor := 1

	for name, forTime := range simulation.InitialContext.Optimize {
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
