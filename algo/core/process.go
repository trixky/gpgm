package core

import "math"

type Process struct {
	Name    string         `json:"name"`
	Inputs  map[string]int `json:"inputs"`
	Outputs map[string]int `json:"outputs"`
	Delay   int            `json:"delay"`
	Parents []int          `json:"parent"`
}

// CanBeExecuted check if it can be executed regarding the stock
func (p *Process) CanBeExecuted(stock *Stock) bool {
	return p.CanBeExecutedXTimes(stock, 1)
}

// CanBeExecutedXTimes check if it can be executed x times regarding the stock
func (p *Process) CanBeExecutedXTimes(stock *Stock, amount int) bool {
	for product, quantity := range p.Inputs {
		if stock.GetResource(product) < quantity*amount {
			return false
		}
	}
	return true
}

// CanBeExecutedXTimes check the maximum number of times it can be executed regarding the stock
func (p *Process) CanBeExecutedMaxXTimes(stock *Stock) int {
	var max_global int = math.MaxInt

	for product, quantity := range p.Inputs {
		max_production := stock.GetResource(product) / quantity

		if max_production < max_global {
			max_global = max_production
		}

		if max_global == 0 {
			return max_global
		}
	}

	return max_global
}

// TryExecute try to execute one time regarding the stock
func (p *Process) TryExecute(stock *Stock) bool {
	new_stock := Stock{}

	for product, cost := range p.Inputs {
		available := stock.GetResource(product)

		if available < cost {
			return false
		} else {
			new_stock.SetResource(product, available-cost)
		}
	}

	*stock = new_stock

	return true
}

// TryExecuteN try to execute n times regarding the stock
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

// ExecuteN executes n times without regarding the stock
func (p *Process) ExecuteN(stock *Stock, n int) {
	// WARNING: the stock is not verified

	for product, cost := range p.Inputs {
		available := stock.GetResource(product)
		stock.SetResource(product, available-(cost*n))
	}
}

// HaveInput check if it has an resource as input
func (p *Process) HaveInput(resource string) bool {
	for input := range p.Inputs {
		// For each input

		if input == resource {
			// If the input corresponding to the resource
			return true
		}
	}

	return false
}

// HaveOutput check if it has an resource as output
func (p *Process) HaveOutput(resource string) bool {
	for output := range p.Outputs {
		// For each output

		if output == resource {
			// If the output corresponding to the resource
			return true
		}
	}

	return false
}
