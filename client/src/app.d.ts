// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}

	// WASM functions
	function WASM_initialize(arguments: string): string;
	function WASM_run_generation(solver: string): string;
	function WASM_generate_output(simulation: string): string;
	function WASM_parse_input(inputFile: string): string | null;
}

export { };
