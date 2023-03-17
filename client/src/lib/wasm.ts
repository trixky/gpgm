import { dev } from "$app/environment";
import type { Arguments, RunningSolver, WASMGenerationReturn } from "../types";
import type { Simulation } from "../types/simulation";
import { parse_as } from "./utils/parse";

export class Wasm {
	public async parseInput(file: string): Promise<string | null> {
		let result: string | null;
		if (dev) {
			result = WASM_parse_input(file);
		} else {
			// Use a fetch request to send a message to a ServiceWorker than can run a generation
			const response = await fetch('/sw/parse', {
				method: 'POST',
				body: file
			});
			result = await response.text();
		}
		return !result ? null : result;
	}

	public async initialize(args: Arguments): Promise<RunningSolver | null> {
		let result: string;
		if (dev) {
			result = WASM_initialize(JSON.stringify(args));
		} else {
			// Use a fetch request to send a message to a ServiceWorker than can run a generation
			const response = await fetch('/sw/initialize', {
				method: 'POST',
				body: JSON.stringify(args)
			});
			result = await response.text();
		}
		try {
			return parse_as<RunningSolver>(result);
		} catch (error) {
			return null
		}
	}

	public async runGeneration(solver: RunningSolver): Promise<WASMGenerationReturn> {
		let result: string;
		if (dev) {
			result = WASM_run_generation(JSON.stringify(solver));
		} else {
			// Use a fetch request to send a message to a ServiceWorker than can run a generation
			const response = await fetch('/sw/generate', {
				method: 'POST',
				body: JSON.stringify(solver)
			});
			result = await response.text();
		}
		return parse_as<WASMGenerationReturn>(result);
	}

	public async generateOutput(simulation: Simulation): Promise<string> {
		let result: string;
		if (dev) {
			result = WASM_generate_output(JSON.stringify(simulation));
		} else {
			// Use a fetch request to send a message to a ServiceWorker than can run a generation
			const response = await fetch('/sw/output', {
				method: 'POST',
				body: JSON.stringify(simulation)
			});
			result = await response.text();
		}
		return result;
	}
}

export default new Wasm()
