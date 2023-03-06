package simulation

import (
	"math"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
	"github.com/trixky/krpsim/algo/parser"
)

type ProcessToBeExecuted struct {
	Process int
	Amount  int
}

type Simulation struct {
	InitialContext parser.SimulationInitialContext
	Stock          core.Stock
	ExpectedStock  []ExpectedStock
	Time           int
	Instance       instance.Instance
}

func NewSimulation(info parser.SimulationInitialContext, instance instance.Instance) Simulation {
	return Simulation{
		InitialContext: info,
		Stock:          info.Stock.DeepCopy(),
		ExpectedStock:  []ExpectedStock{},
		Time:           0,
		Instance:       instance,
	}
}

func (s *Simulation) canExecuteProcess(process core.Process) bool {
	for product, quantity := range process.Inputs {
		if s.Stock.Get(product) < quantity {
			return false
		}
	}
	return true
}

func (s *Simulation) canExecuteAnyProcess() bool {
	for _, process := range s.InitialContext.Processes {
		if s.canExecuteProcess(process) {
			return true
		}
	}
	return false
}

func (s *Simulation) Run(maxCycle int) {
	for ; s.Time < maxCycle; s.Time++ {
		// ? Abort early if there is no executable processes and no expected stocks
		if !s.canExecuteAnyProcess() && len(s.ExpectedStock) == 0 {
			break
		}

		// * Update stock for the current cycle
		ready := []ExpectedStock{}
		incomplete := []ExpectedStock{}
		for _, e := range s.ExpectedStock {
			if e.remainingCycles == 0 {
				ready = append(ready, e)
			} else {
				incomplete = append(incomplete, e)
			}
		}
		for _, e := range ready {
			s.Stock.Add(e.name, e.quantity)
		}
		s.ExpectedStock = incomplete

		// ? Execute actions from genes
		// actions := s.ApplyGenes()
		var actions []ProcessToBeExecuted

		// * Calculate stock
		for _, action := range actions {
			process := s.InitialContext.Processes[action.Process]
			for name, quantity := range process.Inputs {
				s.Stock.Remove(name, quantity)
			}
			for name, quantity := range process.Outputs {
				s.ExpectedStock = append(s.ExpectedStock, ExpectedStock{
					name:            name,
					quantity:        quantity,
					remainingCycles: process.Delay,
				})
			}
		}

		// * Skip cycles until the next expected stock is ready if no process can be executed
		if !s.canExecuteAnyProcess() && len(s.ExpectedStock) > 0 {
			var closer ExpectedStock
			closer.remainingCycles = math.MaxInt
			for _, e := range s.ExpectedStock {
				if e.remainingCycles < closer.remainingCycles {
					closer = e
				}
			}
			s.Time += closer.remainingCycles
			for _, e := range s.ExpectedStock {
				e.remainingCycles -= closer.remainingCycles
			}
		}
	}
}
