package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"syscall/js" // Skip the error vvv
	"time"

	// syscall/js package is supposed to be compiled on wasm
	// architecture with js as the OS but the editor is not aware of this

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/parser"
	"github.com/trixky/krpsim/algo/population"
	"github.com/trixky/krpsim/algo/simulation"
)

type Arguments struct {
	Text           string `json:"text"`
	MaxGeneration  int    `json:"generations"`
	MaxCycle       int    `json:"deep"`
	PopulationSize int    `json:"population"`
}

type RunningSolver struct {
	Population population.Population `json:"population"`
	Context    core.InitialContext   `json:"context"`
	Options    core.Options          `json:"options"`
	Generation int                   `json:"generation"`
	Start      time.Time             `json:"start"`
}

type WASMGenerationReturn struct {
	ScoredPopulation population.ScoredPopulation `json:"scored_population"`
	RunningSolver    RunningSolver               `json:"running_solver"`
}

// * Run a single generation for the given RunningSolver
func runGeneration(solver RunningSolver) (population.ScoredPopulation, RunningSolver) {
	scored := solver.Population.RunAllSimulations(solver.Context, solver.Options)
	solver.Population = scored.Crossover(solver.Options)
	solver.Population.Mutate(solver.Context, solver.Options)
	solver.Generation += 1

	return scored, solver
}

// * Initialize a RunningSolver from the given args
func initialize(args Arguments) (RunningSolver, error) {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return RunningSolver{}, err
	}
	options := core.Options{ // TODO Collect Options
		PopulationSize:   args.PopulationSize,
		MaxGeneration:    args.MaxGeneration,
		MaxCycle:         args.MaxCycle,
		TimeLimitSeconds: 60,
		UseElitism:       true,
		ElitismAmount:    1,
	}

	return RunningSolver{
		Population: population.NewRandomPopulation(context, options),
		Context:    context,
		Options:    options,
		Generation: 1,
		Start:      time.Now(),
	}, nil
}

// runSimulation run the main simulation
func runSimulation(args Arguments) string {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return fmt.Sprintf("unexpected error: %v", err)
	}
	options := core.Options{
		PopulationSize: 50,
		MaxGeneration:  100,
		// MaxCycle:         100,
		MaxCycle:         1000,
		TimeLimitSeconds: 60,
	}

	var scored population.ScoredPopulation
	population := population.NewRandomPopulation(context, options)
	// s, _ := json.MarshalIndent(population, "", "\t")
	// fmt.Printf("%s\n", string(s))
	generation := 1
	start := time.Now()
	for ; ; generation += 1 {
		fmt.Printf("generation %d since %fs\n", generation, time.Since(start).Seconds())
		scored = population.RunAllSimulations(context, options)
		if generation >= options.MaxGeneration || time.Since(start).Seconds() > float64(options.TimeLimitSeconds) {
			break
		}
		// fmt.Printf("%v\n", population)
		population := scored.Crossover(options)
		population.Mutate(context, options)
		// fmt.Printf("%v\n", population)
	}
	best := scored.Best()
	// s, _ = json.MarshalIndent(best, "", "\t")
	s, _ := json.MarshalIndent(best, "", "\t")

	return string(s)
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

		for i_index, instance := range running_solver.Population.Instances {
			fmt.Println("***********************", i_index)
			for g_index, gene := range instance.Chromosome.Genes {
				fmt.Println("*****", g_index)

				keys := make([]string, len(gene.History))
				i := 0
				for key := range gene.History {
					keys[i] = key
					i++
				}

				sort.Strings(keys)

				for _, sorted_key := range keys {
					fmt.Println(sorted_key)
				}

			}
		}

		if err != nil {
			fmt.Print(err.Error())

			return nil
		}

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
		solver := RunningSolver{}

		// --------- extract the response
		if err := json.Unmarshal([]byte(args[0].String()), &solver); err != nil {
			fmt.Print(err.Error())

			return nil
		}

		// --------- call
		population, new_solver := runGeneration(solver)

		// --------- insert the response
		scored_population_running_solver_json, err := json.Marshal(WASMGenerationReturn{
			ScoredPopulation: population,
			RunningSolver:    new_solver,
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

func main() {
	// Register the shared function
	js.Global().Set("WASM_initialize", initializeWasm())
	js.Global().Set("WASM_run_generation", runGenerationWasm())
	js.Global().Set("WASM_generate_output", generateOutput())

	fmt.Println("Go Web Assembly Loaded")

	// Keep the program open
	<-make(chan bool)
}
