package core

type Product struct {
	Name     string
	Quantity int
}

type Process struct {
	Name    string
	Inputs  map[string]int
	Outputs map[string]int
	Delay   int
}

func (p *Process) CanBeExecuted(stock Stock) bool {
	for product, quantity := range p.Inputs {
		if stock.Get(product) < quantity {
			println("got", stock.Get(product), product, "needed", quantity)
			return false
		}
	}
	return true
}

type SimulationInitialContext struct {
	Stock     Stock
	Processes []Process
	Optimize  map[string]bool
}

func (sm *SimulationInitialContext) IsInOutput(product string) bool {
	for _, process := range sm.Processes {
		for outputProduct := range process.Outputs {
			if product == outputProduct {
				return true
			}
		}
	}
	return false
}