package instance

import (
	"sort"
	"strings"
	"testing"

	"github.com/trixky/krpsim/algo/core"
)

func TestInstanceInit(t *testing.T) {
	initial_context := core.InitialContext{
		Stock: core.Stock{
			"wood":   10,
			"stone":  10,
			"door":   10,
			"window": 10,
			"house":  1,
			"gold":   30,
		},
		Processes: []core.Process{
			{ // ----------------------- 0
				Name: "gold_wood",
				Inputs: map[string]int{
					"gold": 1,
				},
				Outputs: map[string]int{
					"wood": 50,
				},
				Delay: 10,
			},
			{ // ----------------------- 1
				Name: "gold_stone",
				Inputs: map[string]int{
					"gold": 1,
				},
				Outputs: map[string]int{
					"stone": 50,
				},
				Delay: 10,
			},
			{ // ----------------------- 2
				Name: "wood_door",
				Inputs: map[string]int{
					"wood": 10,
				},
				Outputs: map[string]int{
					"door": 1,
				},
				Delay: 10,
			},
			{ // ----------------------- 3
				Name: "wood_window",
				Inputs: map[string]int{
					"wood": 3,
				},
				Outputs: map[string]int{
					"window": 1,
				},
				Delay: 10,
			},
			{ // ----------------------- 4
				Name: "build_house",
				Inputs: map[string]int{
					"gold":   10,
					"wood":   50,
					"stone":  50,
					"door":   1,
					"window": 6,
				},
				Outputs: map[string]int{
					"house": 1,
				},
				Delay: 10,
			},
			{ // ----------------------- 5
				Name: "sale_house",
				Inputs: map[string]int{
					"house": 6,
				},
				Outputs: map[string]int{
					"gold": 200,
				},
				Delay: 10,
			},
		},
		Optimize: make(map[string]bool),
	}

	expected := []string{
		"054_154_2_3_4_42_43_454_54_542_543_@",
		"054_154_4_454_54_@",
		"054_154_4_454_54_@",
		"054_154_4_454_54_@",
		"05_15_205_305_405_415_45_5_545_@",
		"0_054_1_154_20_30_4_40_41_420_430_454_54_540_541_@",
	}

	initial_context.FindProcessParents()

	instance := Instance{}

	options := core.Options{
		RandomCut:            true,
		MaxCut:               0,
		HistoryPartMaxLength: 3,
		HistoryKeyMaxLength:  4,
	}

	instance.Init(initial_context.Processes, initial_context.Optimize, &options)

	for g_index, gene := range instance.Chromosome.PriorityGenes {
		keys := make([]string, len(gene.HistoryProcessDependencies))
		i := 0
		for key := range gene.HistoryProcessDependencies {
			if key == "" {
				key = "@"
			}
			keys[i] = key
			i++
		}

		sort.Strings(keys)

		concatened_keys := strings.Join(keys, "_")

		if expected[g_index] != concatened_keys {
			t.Fatalf(`test 0 (gene: %d): expected = %s, got = %s`, g_index, expected[g_index], concatened_keys)
		}
	}
}
