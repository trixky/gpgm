package simulation

import (
	"math"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
	"github.com/trixky/krpsim/algo/interpretor"
)

type ProcessToBeExecuted struct {
	Process int
	Amount  int
}

type Simulation struct {
	InitialContext core.InitialContext
	Instance       instance.Instance
	Stock          core.Stock
	ExpectedStock  []ExpectedStock
	Cycle          int
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
		if s.canExecuteAnyProcess() {
			actions1 := interpretor.Interpret(s.Instance, s.InitialContext, s.Stock)
			var actions []ProcessToBeExecuted
			for _, process := range actions1 {
				var processId int
				for i, maybeProcess := range s.InitialContext.Processes {
					if maybeProcess.Name == process.Name {
						processId = i
						break
					}
				}
				actions = append(actions, ProcessToBeExecuted{
					Process: processId,
					Amount:  1,
				})
			}

			// * Calculate stock
			for _, action := range actions {
				process := s.InitialContext.Processes[action.Process]
				for name, quantity := range process.Inputs {
					s.Stock.Remove(name, quantity)
				}
				for name, quantity := range process.Outputs {
					s.ExpectedStock = append(s.ExpectedStock, ExpectedStock{
						name:            name,
						quantity:        quantity * action.Amount,
						remainingCycles: process.Delay,
					})
				}
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
			s.Cycle += closer.remainingCycles
			for _, e := range s.ExpectedStock {
				e.remainingCycles -= closer.remainingCycles
			}
		}
	}
}
