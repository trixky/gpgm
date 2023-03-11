package core

import (
	"math"
)

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

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

// FindProcessParents find process parents the initial context
func (sm *InitialContext) FindProcessParents() {
	for child_index, child := range sm.Processes {
		// For each child process
		for parent_index, parent := range sm.Processes {
			// For each parent process
			// Note that parent can be the child
			for resource_name := range parent.Inputs {
				// For each input resource of the parent
				if output, ok := child.Outputs[resource_name]; ok {
					// If the child has the X input resource of the parent as output
					if input, ok := child.Inputs[resource_name]; !ok || output > input {
						// If its X output is greater than its input if it as an input
						sm.Processes[child_index].Parents = append(sm.Processes[child_index].Parents, parent_index)
					}
				}
			}
		}
	}

	// for _, process := range sm.Processes {
	// 	fmt.Println("*****************", process.Name)
	// 	for _, process_parent := range process.Parents {
	// 		fmt.Println(sm.Processes[process_parent].Name)
	// 	}

	// }

	// os.Exit(1)
	// fmt.Println(sm.Processes)
}
