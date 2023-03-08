package core

import "math"

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Process struct {
	Name    string         `json:"name"`
	Inputs  map[string]int `json:"inputs"`
	Outputs map[string]int `json:"outputs"`
	Delay   int            `json:"delay"`
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

func (p *Process) CanBeExecutedMaxXTimes(stock *Stock) uint16 {
	var max_global uint16 = math.MaxUint16

	for product, quantity := range p.Inputs {
		max_production := uint16(stock.Get(product) / quantity)

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

func (p *Process) IsInOutput(product string) bool {
	for outputProduct := range p.Outputs {
		if product == outputProduct {
			return true
		}
	}

	return false
}

func (p *Process) ExecuteN(stock *Stock, n int) {
	for product, cost := range p.Inputs {
		available := stock.Get(product)
		stock.Insert(product, available-(cost*n))
	}
}

type InitialContext struct {
	Stock      Stock           `json:"stock"`
	Processes  []Process       `json:"processes"`
	Optimize   map[string]bool `json:"optimize"`
	ScoreRatio map[string]int  `json:"score_ratio"`
}

func (sm *InitialContext) IsInOutput(product string) bool {
	for _, process := range sm.Processes {
		if process.IsInOutput(product) {
			return true
		}
	}
	return false
}
