package simulation

import (
	"fmt"
	"math"
	"strings"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
	"github.com/trixky/krpsim/algo/interpretor"
)

type ExecutedProcess struct {
	Cycle   int          `json:"cycle"`
	Process core.Process `json:"process"`
	Amount  int          `json:"amount"`
}

type Simulation struct {
	InitialContext core.InitialContext `json:"initial_context"`
	Instance       instance.Instance   `json:"instance"`
	Stock          core.Stock          `json:"stock"`
	ExpectedStock  []ExpectedStock     `json:"expected_stock"`
	History        []ExecutedProcess   `json:"history"`
	Cycle          int                 `json:"cycle"`
}

func NewSimulation(info core.InitialContext, instance instance.Instance) Simulation {
	return Simulation{
		InitialContext: info,
		Instance:       instance,
		Stock:          info.Stock.DeepCopy(),
		ExpectedStock:  []ExpectedStock{},
		Cycle:          0,
	}
}

func (s *Simulation) CalulateFitness() int {
	return Fitness(*s)
}

func (s *Simulation) canExecuteAnyProcess() bool {
	for _, process := range s.InitialContext.Processes {
		if process.CanBeExecuted(&s.Stock) {
			return true
		}
	}
	return false
}

func (s *Simulation) Run(maxCycle int) {
	for ; s.Cycle < maxCycle; s.Cycle++ {
		// fmt.Println("simulation ========================================== cycle", s.Cycle)
		// ? Abort early if there is no executable processes and no expected stocks
		if !s.canExecuteAnyProcess() && len(s.ExpectedStock) == 0 {
			break
		}

		// * Update stock for the current cycle
		ready := []ExpectedStock{}
		incomplete := []ExpectedStock{}
		for _, e := range s.ExpectedStock {
			if e.RemainingCycles == 0 {
				ready = append(ready, e)
			} else {
				incomplete = append(incomplete, e)
			}
		}
		for _, e := range ready {
			s.Stock.Add(e.Name, e.Quantity)
		}
		s.ExpectedStock = incomplete

		// ? Execute actions from genes
		if s.canExecuteAnyProcess() {
			// fmt.Println("simulation =========================== stock avant")
			// fmt.Println(s.Stock)
			process_quantities_stack := interpretor.Interpret(s.Instance, s.InitialContext, s.Stock.DeepCopy())

			// * Calculate stock
			for _, process_quantity := range process_quantities_stack.Stack {
				for name, quantity := range process_quantity.Process.Inputs {
					s.Stock.Remove(name, quantity*process_quantity.Amount)
				}
				for name, quantity := range process_quantity.Process.Outputs {
					s.ExpectedStock = append(s.ExpectedStock, ExpectedStock{
						Name:            name,
						Quantity:        quantity * process_quantity.Amount,
						RemainingCycles: process_quantity.Process.Delay,
					})
				}
				s.History = append(s.History, ExecutedProcess{
					Cycle:   s.Cycle,
					Process: *process_quantity.Process,
					Amount:  process_quantity.Amount,
				})
			}
			// fmt.Println("simulation =========================== stock apres")
			// fmt.Println(s.Stock)
		}

		// * Skip cycles until the next expected stock is ready if no process can be executed
		if !s.canExecuteAnyProcess() && len(s.ExpectedStock) > 0 {
			var closer ExpectedStock
			closer.RemainingCycles = math.MaxInt
			for i := range s.ExpectedStock {
				if s.ExpectedStock[i].RemainingCycles < closer.RemainingCycles {
					closer = s.ExpectedStock[i]
				}
			}
			s.Cycle += closer.RemainingCycles - 1
			for i := range s.ExpectedStock {
				s.ExpectedStock[i].RemainingCycles -= closer.RemainingCycles
			}
		} else {
			for i := range s.ExpectedStock {
				s.ExpectedStock[i].RemainingCycles -= 1
			}
		}
	}
	for _, e := range s.ExpectedStock {
		s.Stock.Add(e.Name, e.Quantity)
	}
	s.ExpectedStock = []ExpectedStock{}
}

func (s *Simulation) GenerateOutputFile() string {
	lines := make([]string, 0)
	for _, action := range s.History {
		lines = append(lines, fmt.Sprintf("%d: %s:%d", action.Cycle, action.Process.Name, action.Amount))
	}
	stock := ""
	for product, quantity := range s.Stock {
		if stock == "" {
			stock = fmt.Sprintf("%s:%d", product, quantity)
		} else {
			stock = fmt.Sprintf("%s;%s:%d", stock, product, quantity)
		}
	}
	lines = append(lines, fmt.Sprintf("stock: %s", stock))
	return strings.Join(lines, "\n")
}
