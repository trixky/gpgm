package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/parser"
	"github.com/trixky/krpsim/algo/population"
	"github.com/trixky/krpsim/algo/simulation"
	"github.com/trixky/krpsim/algo/solver"
)

type Arguments struct {
	Text                  string               `json:"text"`
	MaxGeneration         int                  `json:"max_generations"`
	MaxCycle              int                  `json:"max_cycle"`
	MaxDepth              int                  `json:"max_depth"`
	MaxCut                int                  `json:"max_cut"`
	TimeLimitMilliseconds int                  `json:"time_limit"`
	PopulationSize        int                  `json:"population_size"`
	ElitismAmount         int                  `json:"elitism_amount"`
	TournamentSize        int                  `json:"tournament_size"`
	TournamentProbability float64              `json:"tournament_probability"`
	CrossoverNewInstances int                  `json:"crossover_new_instances"`
	SelectionMethod       core.SelectionMethod `json:"selection_method"`
	MutationMethod        core.MutationMethod  `json:"mutation_method"`
}

type WASMGenerationReturn struct {
	ScoredPopulation population.ScoredPopulation `json:"scored_population"`
	RunningSolver    solver.RunningSolver        `json:"running_solver"`
}

// initialize initialize a solver from the given args
func initialize(args Arguments) (solver.RunningSolver, error) {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return solver.RunningSolver{}, err
	}

	history_global_max_length := args.MaxDepth

	if history_global_max_length > 3 {
		history_global_max_length = 3
	}

	options := core.Options{
		MaxGeneration: args.MaxGeneration,
		TimeLimitMS:   args.TimeLimitMilliseconds,
		MaxCycle:      args.MaxCycle,
		MaxDepth:      args.MaxDepth,
		// Population
		PopulationSize:        args.PopulationSize,
		ElitismAmount:         args.ElitismAmount,
		SelectionMethod:       args.SelectionMethod,
		TournamentSize:        args.TournamentSize,
		TournamentProbability: args.TournamentProbability,
		CrossoverNewInstances: args.CrossoverNewInstances,
		// Mutation
		MutationChance: 0,
		MutationMethod: args.MutationMethod,
		// Genetic
		NEntry:               1, // HARDCODED
		HistoryPartMaxLength: history_global_max_length,
		HistoryKeyMaxLength:  history_global_max_length,
		RandomCut:            true,
		MaxCut:               args.MaxCut,
	}

	return solver.RunningSolver{
		Population: population.NewRandomPopulation(context, &options),
		Context:    context,
		Options:    options,
		Generation: 1,
	}, nil
}

// runWasm parse arguments, run the simulation and return its result
func initializeWasm() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		arguments := Arguments{}

		// --------- Extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &arguments); err != nil {
			fmt.Print(err.Error())

			return nil
		}

		// --------- Call
		running_solver, err := initialize(arguments)

		if err != nil {
			fmt.Print(err.Error())

			return nil
		}
		running_solver.Context.FindProcessParents()

		// --------- Insert the response
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

		// --------- Extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &solver); err != nil {
			fmt.Print(err.Error())

			return nil
		}

		// --------- Init the timer
		solver.InitTimer()

		// --------- Call
		population := solver.RunGeneration()

		// --------- Insert the response
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

		// --------- Extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &simulation); err != nil {
			fmt.Print(err.Error())
			return nil
		}

		// --------- Call
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
