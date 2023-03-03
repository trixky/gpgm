package main

import (
	"fmt"
	"syscall/js" // Skip the error vvv
	// syscall/js package is supposed to be compiled on wasm
	// architecture with js as the OS but the editor is not aware of this
)

type Arguments struct {
	Text  string
	Delay int
}

// runSimulation run the main simulation
func runSimulation(args Arguments) string {
	// put the parsing / gens / simulations etc... here

	return "#result#"
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
