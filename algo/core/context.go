package core

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

func (p *Process) Execute(stock *Stock) bool {
	new_stock := Stock{}

	for product, quantity := range p.Inputs {
		stock_quantity := stock.Get(product)

		if stock_quantity < quantity {
			return false
		} else {
			new_stock.Insert(product, stock_quantity-quantity)
		}
	}

	*stock = new_stock

	return true
}

type InitialContext struct {
	Stock     Stock
	Processes []Process
	Optimize  map[string]bool
}

func (sm *InitialContext) IsInOutput(product string) bool {
	for _, process := range sm.Processes {
		for outputProduct := range process.Outputs {
			if product == outputProduct {
				return true
			}
		}
	}
	return false
}
