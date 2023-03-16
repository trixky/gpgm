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
	Cycle    int          `json:"cycle"`
	Process  core.Process `json:"process"`
	Quantity int          `json:"quantity"`
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
		Stock:          *info.Stock.DeepCopy(),
		ExpectedStock:  []ExpectedStock{},
		Cycle:          0,
	}
}

func (s *Simulation) CalulateFitness(options *core.Options) int {
	return Fitness(*s, options)
}

func (s *Simulation) canExecuteAnyProcess() bool {
	for _, process := range s.InitialContext.Processes {
		if process.CanBeExecuted(&s.Stock) {
			return true
		}
	}
	return false
}

func (s *Simulation) Run(options *core.Options) {
	s.Cycle = -1
	for s.Cycle < options.MaxCycle {
		// ? Abort early if there is no executable processes and no expected stocks
		if !s.canExecuteAnyProcess() && len(s.ExpectedStock) == 0 {
			break
		}
		s.Cycle++

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
			s.Stock.AddResource(e.Name, e.Quantity)
		}
		s.ExpectedStock = incomplete

		// ? Execute actions from genes
		if s.canExecuteAnyProcess() {
			stock_copy := s.Stock.DeepCopy()

			process_quantities_stack := interpretor.Interpret(s.Instance, s.InitialContext, stock_copy, options)
			// * Calculate stock
			for _, process_quantity := range process_quantities_stack.Stack {
				for name, quantity := range process_quantity.Process.Inputs {
					s.Stock.RemoveResource(name, quantity*process_quantity.Quantity)
				}
				for name, quantity := range process_quantity.Process.Outputs {
					s.ExpectedStock = append(s.ExpectedStock, ExpectedStock{
						Name:            name,
						Quantity:        quantity * process_quantity.Quantity,
						RemainingCycles: process_quantity.Process.Delay,
					})
				}
				s.History = append(s.History, ExecutedProcess{
					Cycle:    s.Cycle,
					Process:  *process_quantity.Process,
					Quantity: process_quantity.Quantity,
				})
			}
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
		s.Stock.AddResource(e.Name, e.Quantity)
	}
	s.ExpectedStock = []ExpectedStock{}
}

type HistoryAction struct {
	Name   string
	Amount int
}

type HistoryEntry struct {
	Cycle   int
	Actions []HistoryAction
}

func (s *Simulation) GenerateOutputFile() string {
	lastCycle := -1
	actions := []HistoryEntry{}
	for _, action := range s.History {
		if lastCycle == action.Cycle {
			lastIndex := len(actions) - 1
			added := false
			for i := 0; i < len(actions[lastIndex].Actions); i++ {
				if actions[lastIndex].Actions[i].Name == action.Process.Name {
					actions[lastIndex].Actions[i].Amount += action.Quantity
					added = true
					break
				}
			}
			if !added {
				actions[lastIndex].Actions = append(actions[lastIndex].Actions, HistoryAction{
					Name:   action.Process.Name,
					Amount: action.Quantity,
				})
			}
		} else {
			actions = append(actions, HistoryEntry{
				Cycle: action.Cycle,
				Actions: []HistoryAction{
					{
						Name:   action.Process.Name,
						Amount: action.Quantity,
					},
				},
			})
		}
		lastCycle = action.Cycle
	}
	lines := make([]string, 0)
	for _, action := range actions {
		parts := []string{}
		for _, action := range action.Actions {
			parts = append(parts, fmt.Sprintf("%s:%d", action.Name, action.Amount))
		}
		lines = append(lines, fmt.Sprintf("%d: %s", action.Cycle, strings.Join(parts, ";")))
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
