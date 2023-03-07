package main

import (
	"encoding/json"
	"fmt"
	"syscall/js" // Skip the error vvv
	"time"

	// syscall/js package is supposed to be compiled on wasm
	// architecture with js as the OS but the editor is not aware of this

	"github.com/trixky/krpsim/algo/core"
	"github.com/trixky/krpsim/algo/parser"
	"github.com/trixky/krpsim/algo/population"
)

type Arguments struct {
	Text  string
	Delay int
	// ...
}

type RunningSolver struct {
	population population.Population
	context    core.InitialContext
	options    core.Options
	generation int
	start      time.Time
}

// * Run a single generation for the given RunningSolver
func runGeneration(solver RunningSolver) (population.ScoredPopulation, RunningSolver) {
	scored := solver.population.RunAllSimulations(solver.context, solver.options)
	solver.population = scored.Crossover(solver.options)
	solver.population.Mutate(solver.options)
	solver.generation += 1

	return scored, solver
}

// * Initialize a RunningSolver from the given args
func initialize(args Arguments) (RunningSolver, error) {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return RunningSolver{}, err
	}
	options := core.Options{ // TODO Collect Options
		PopulationSize:   50,
		MaxGeneration:    100,
		MaxCycle:         100,
		TimeLimitSeconds: 60,
	}

	return RunningSolver{
		population: population.NewRandomPopulation(context, options),
		context:    context,
		options:    options,
		generation: 1,
		start:      time.Now(),
	}, nil
}

// runSimulation run the main simulation
func runSimulation(args Arguments) string {
	context, err := parser.ParseSimulationFile(args.Text)
	if err != nil {
		return fmt.Sprintf("unexpected error: %v", err)
	}
	options := core.Options{
		PopulationSize:   50,
		MaxGeneration:    100,
		MaxCycle:         100,
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
		population.Mutate(options)
		// fmt.Printf("%v\n", population)
	}
	best := scored.Best()
	// s, _ = json.MarshalIndent(best, "", "\t")
	s, _ := json.MarshalIndent(best, "", "\t")

	return string(s)
}

// runWasm parse arguments, run the simulation and return its result
func runWasm() js.Func {
	run := js.FuncOf(func(this js.Value, args []js.Value) any {
		parsed_args := Arguments{
			Text:  args[0].String(),
			Delay: args[1].Int(),
		}

		result := runSimulation(parsed_args)

		return result
	})

	return run
}

func main() {
	// Register the shared function
	js.Global().Set("Run", runWasm())

	fmt.Println("Go Web Assembly Loaded")

	// Keep the program open
	<-make(chan bool)
}
