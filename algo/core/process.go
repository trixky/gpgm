package core

import "math"

type Process struct {
	Name    string         `json:"name"`
	Inputs  map[string]int `json:"inputs"`
	Outputs map[string]int `json:"outputs"`
	Delay   int            `json:"delay"`
	Parents []int          `json:"parent"`
}

func (p *Process) CanBeExecuted(stock *Stock) bool {
	return p.CanBeExecutedXTimes(stock, 1)
}

func (p *Process) CanBeExecutedXTimes(stock *Stock, amount int) bool {
	for product, quantity := range p.Inputs {
		if stock.Get(product) < quantity*amount {
			return false
		}
	}
	return true
}

func (p *Process) CanBeExecutedMaxXTimes(stock *Stock) int {
	var max_global int = math.MaxInt

	for product, quantity := range p.Inputs {
		max_production := stock.Get(product) / quantity

		if max_production < max_global {
			max_global = max_production
		}

		if max_global == 0 {
			return max_global
		}
	}

	return max_global
}

func (p *Process) TryExecute(stock *Stock) bool {
	new_stock := Stock{}

	for product, cost := range p.Inputs {
		available := stock.Get(product)

		if available < cost {
			return false
		} else {
			new_stock.Insert(product, available-cost)
		}
	}

	*stock = new_stock

	return true
}

// TryExecuteN try to execute n time itself and returns number of execution
func (p *Process) TryExecuteN(stock *Stock, n int) int {
	max := p.CanBeExecutedMaxXTimes(stock)

	if max > 0 {
		if max > n {
			max = n
		}
		p.ExecuteN(stock, max)
	}

	return max
}

func (p *Process) ExecuteN(stock *Stock, n int) {
	for product, cost := range p.Inputs {
		available := stock.Get(product)
		stock.Insert(product, available-(cost*n))
	}
}

// HaveInput check if it has an resource as input
func (p *Process) HaveInput(resource string) bool {
	for input := range p.Inputs {
		if input == resource {
			return true
		}
	}

	return false
}

// HaveOutput check if it has an resource as output
func (p *Process) HaveOutput(resource string) bool {
	for output := range p.Outputs {
		if output == resource {
			return true
		}
	}

	return false
}
