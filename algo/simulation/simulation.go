package simulation

import "github.com/trixky/krpsim/algo/parser"

type Simulation struct {
	Parameters parser.SimulationParameters
}

type SimulationContext struct {
}

func (s *Simulation) Start()
