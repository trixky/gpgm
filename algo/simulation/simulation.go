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

type SimulationList []Simulation

func (l *SimulationList) pop() Simulation {
	length := len(*l)
	s := (*l)[length-1]
	*l = (*l)[:length-1]
	return s
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

func (s *Simulation) DeepCopy() Simulation {
	expectedStock := make([]ExpectedStock, len(s.ExpectedStock))
	copy(expectedStock, s.ExpectedStock)
	history := make([]ExecutedProcess, len(s.History))
	copy(history, s.History)
	return Simulation{
		InitialContext: s.InitialContext,
		Instance:       s.Instance,
		Stock:          s.Stock.DeepCopy(),
		ExpectedStock:  expectedStock,
		Cycle:          s.Cycle,
		History:        history,
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

func (s *Simulation) ExecutableProcesses() []core.Process {
	var processes = make([]core.Process, 0)
	for _, process := range s.InitialContext.Processes {
		if process.CanBeExecuted(&s.Stock) {
			processes = append(processes, process)
		}
	}
	return processes
}

func (s *Simulation) Execute(process core.Process, amount int) {
	for name, quantity := range process.Inputs {
		s.Stock.Remove(name, quantity*amount)
	}
	for name, quantity := range process.Outputs {
		s.ExpectedStock = append(s.ExpectedStock, ExpectedStock{
			Name:            name,
			Quantity:        quantity * amount,
			RemainingCycles: process.Delay,
		})
	}
	s.History = append(s.History, ExecutedProcess{
		Cycle:   s.Cycle,
		Process: process,
		Amount:  amount,
	})
}

func (s *Simulation) UpdateExpectedStocks() {
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
}

func (s *Simulation) SkipToClosestStock() {
	closer := math.MaxInt
	for i := range s.ExpectedStock {
		if s.ExpectedStock[i].RemainingCycles < closer {
			closer = s.ExpectedStock[i].RemainingCycles
		}
	}
	s.Cycle += closer
	for i := range s.ExpectedStock {
		s.ExpectedStock[i].RemainingCycles -= closer
	}
}

func (s *Simulation) CopyAndExecute(process core.Process, amount int) Simulation {
	sm := s.DeepCopy()
	sm.Execute(process, amount)
	return sm
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
			process_quantities := interpretor.Interpret(s.Instance, s.InitialContext, s.Stock.DeepCopy())

			// * Calculate stock
			for _, process_quantity := range process_quantities {
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

func (s *Simulation) RunBrute(options core.Options) []Simulation {
	explored := 0
	nodes := make(SimulationList, 0)
	nodes = append(nodes, *s)
	best := Simulation{}
	bestScore := -1

	for len(nodes) > 0 {
		explored += 1
		node := nodes.pop()
		// Early exit
		if node.Cycle > options.MaxCycle {
			for len(node.ExpectedStock) > 0 {
				node.SkipToClosestStock()
				node.UpdateExpectedStocks()
			}
			score := node.CalulateFitness()
			if bestScore < score {
				bestScore = score
				best = node
			}
			continue
		}

		processes := node.ExecutableProcesses()
		// Skip to closest stock until all used or a process is executable
		for len(processes) == 0 && len(node.ExpectedStock) > 0 {
			node.SkipToClosestStock()
			node.UpdateExpectedStocks()
			processes = node.ExecutableProcesses()
		}
		if len(processes) == 0 {
			score := node.CalulateFitness()
			if bestScore < score {
				bestScore = score
				best = node
			}
			continue
		}
		for _, process := range processes {
			maxExecution := process.CanBeExecutedMaxXTimes(&node.Stock)
			for i := 1; i <= int(maxExecution); i++ {
				simulated := node.CopyAndExecute(process, i)
				nodes = append(nodes, simulated)
			}
		}
	}

	fmt.Printf("Explored %d nodes\n", explored)
	return []Simulation{best}
}

func (s *Simulation) GenerateOutputFile() string {
	lines := make([]string, 0)
	lastCycle := -1
	for cycle, action := range s.History {
		if lastCycle == cycle {
			lastLine := lines[len(lines)-1]
			lines[len(lines)-1] = fmt.Sprintf("%s;%s:%d", lastLine, action.Process.Name, action.Amount)
		} else {
			lines = append(lines, fmt.Sprintf("%d: %s:%d", action.Cycle, action.Process.Name, action.Amount))
		}
		lastCycle = cycle
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
