package interpretor

import (
	"testing"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/instance"
)

func TestInterpretor(t *testing.T) {
	stock := core.Stock{
		"cat": 3,
		"dog": 0,
		"pig": 0,
	}

	initial_context := core.InitialContext{
		Stock: core.Stock{
			"cat": 5,
			"dog": 0,
			"pig": 0,
		},
		Processes: []core.Process{
			{
				Name: "111",
				Inputs: map[string]int{
					"cat": 1,
				},
				Outputs: map[string]int{
					"pig": 1,
				},
				Delay: 100,
			},
			{
				Name: "222",
				Inputs: map[string]int{
					"cat": 1,
				},
				Outputs: map[string]int{
					"pig": 10,
				},
				Delay: 100,
			},
		},
		Optimize: map[string]bool{
			"cat": false,
		},
	}

	instance := instance.Instance{
		Chromosome: instance.Chromosome{
			Genes: []instance.Gene{
				{
					FirstPriorityExon: instance.Exon{
						Value: 1,
					},
					LastPriorityExon: instance.Exon{
						Value: 2,
					},
					RatioExons: []instance.Exon{
						{
							Value: 1,
						},
						{
							Value: 2,
						},
					},
				},
			},
		},
	}

	processes := Interpret(instance, initial_context, stock)

	if len(processes) != 3 {
		t.Fatalf(`expected = %d, got = %d`, 3, len(processes))
	}
}
