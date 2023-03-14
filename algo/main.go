package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"syscall/js"
	"time"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/parser"
	"github.com/trixky/krpsim/algo/population"
	"github.com/trixky/krpsim/algo/simulation"
	"github.com/trixky/krpsim/algo/solver"
)

type Arguments struct {
	Text           string `json:"text"`
	MaxGeneration  int    `json:"generations"`
	MaxCycle       int    `json:"deep"`
	PopulationSize int    `json:"population"`
}

type WASMGenerationReturn struct {
	ScoredPopulation population.ScoredPopulation `json:"scored_population"`
	RunningSolver    solver.RunningSolver        `json:"running_solver"`
}

// * Initialize a solver.RunningSolver from the given args
func initialize(args Arguments) (solver.RunningSolver, error) {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return solver.RunningSolver{}, err
	}
	options := core.Options{ // TODO Collect Options
		MaxGeneration:    args.MaxGeneration,
		TimeLimitSeconds: 60,
		MaxCycle:         args.MaxCycle,
		MaxDepth:         6,
		// Population
		PopulationSize:        args.PopulationSize,
		ElitismAmount:         1,
		SelectionMethod:       core.TournamentSelection,
		TournamentSize:        25,
		TournamentProbability: 0.77,
		CrossoverNewInstances: 1,
		// Mutation
		MutationChance: 0,
		MutationMethod: core.LogarithmicMutation,
		// Genetic
		NEntry:               1,
		HistoryPartMaxLength: 3,
		HistoryKeyMaxLength:  3,
		RandomCut:            true,
		MaxCut:               0,
	}

	return solver.RunningSolver{
		Population: population.NewRandomPopulation(context, &options),
		Context:    context,
		Options:    options,
		Generation: 1,
		Start:      time.Now(),
	}, nil
}

func printDependencies(running_solver solver.RunningSolver) {
	for i_index, instance := range running_solver.Population.Instances {
		fmt.Println("***********************", i_index)
		for g_index, gene := range instance.Chromosome.PriorityGenes {
			fmt.Println("*****", g_index)

			keys := make([]string, len(gene.HistoryProcessDependencies))
			i := 0
			for key := range gene.HistoryProcessDependencies {
				keys[i] = key
				i++
			}

			sort.Strings(keys)

			for _, sorted_key := range keys {
				fmt.Println(sorted_key)
				for _, dependencie := range gene.HistoryProcessDependencies[sorted_key].InputDependencies {
					fmt.Println(dependencie.Input)
					fmt.Println(dependencie.ProcessDependencies)
				}
			}
		}
	}

}

// runWasm parse arguments, run the simulation and return its result
func initializeWasm() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		arguments := Arguments{}

		// --------- extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &arguments); err != nil {
			fmt.Print(err.Error())

			return nil
		}

		// --------- call
		running_solver, err := initialize(arguments)

		if err != nil {
			fmt.Print(err.Error())

			return nil
		}
		running_solver.Context.FindProcessParents()

		// --------- insert the response
		running_solver_json, err := json.Marshal(running_solver)

		if err != nil {
			fmt.Print(err.Error())

			return nil
		}

		return string(running_solver_json)
	})

	return run
}

// runWasm parse arguments, run the simulation and return its result
func runGenerationWasm() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		rand.Seed(time.Now().UnixNano())
		solver := solver.RunningSolver{}

		// --------- extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &solver); err != nil {
			fmt.Print(err.Error())

			return nil
		}

		// --------- call
		population := solver.RunGeneration()

		// --------- insert the response
		scored_population_running_solver_json, err := json.Marshal(WASMGenerationReturn{
			ScoredPopulation: population,
			RunningSolver:    solver,
		})

		if err != nil {
			fmt.Print(err.Error())

			return nil
		}

		return string(scored_population_running_solver_json)
	})

	return run
}

// runWasm parse arguments, run the simulation and return its result
func generateOutput() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		simulation := simulation.Simulation{}

		// --------- extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &simulation); err != nil {
			fmt.Print(err.Error())
			return nil
		}

		// --------- call
		output := simulation.GenerateOutputFile()

		return output
	})

	return run
}

// Parse the given input file to check if there is any errors
func parseInput() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		input := args[0].String()

		_, err := parser.ParseSimulationFile(input)
		if err != nil {
			return err.Error()
		}

		return nil
	})

	return run
}

func main() {
	// Register the shared function
	js.Global().Set("WASM_initialize", initializeWasm())
	js.Global().Set("WASM_run_generation", runGenerationWasm())
	js.Global().Set("WASM_generate_output", generateOutput())
	js.Global().Set("WASM_parse_input", parseInput())

	fmt.Println("Go Web Assembly Loaded")

	// Keep the program open
	<-make(chan bool)
}
