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
	}

	expected := []string{
		".2_.2.4_.2.4.5_.3_.3.4_.3.4.5_.4_.4.5_.4.5.0_.4.5.1_.4.5.4",
		".4_.4.5_.4.5.0_.4.5.1_.4.5.4",
		".4_.4.5_.4.5.0_.4.5.1_.4.5.4",
		".4_.4.5_.4.5.0_.4.5.1_.4.5.4",
		".5_.5.0_.5.0.2_.5.0.3_.5.0.4_.5.1_.5.1.4_.5.4_.5.4.5",
		".0_.0.2_.0.2.4_.0.3_.0.3.4_.0.4_.0.4.5_.1_.1.4_.1.4.5_.4_.4.5_.4.5.0_.4.5.1_.4.5.4",
	}

	initial_context.FindProcessParents()

	instance := Instance{}

	instance.Init(initial_context.Processes)

	for g_index, gene := range instance.Chromosome.Genes {
		keys := make([]string, len(gene.History))
		i := 0
		for key := range gene.History {
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
