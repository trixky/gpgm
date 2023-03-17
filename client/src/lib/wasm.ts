import { get } from "svelte/store";
import { useWorker } from "./stores/useWorker";
import { workerReady } from "./stores/workerReady";
import type { Simulation } from "../types/simulation";
import type { Arguments, RunningSolver, WASMGenerationReturn } from "../types";
import { parse_as } from "./utils/parse";

export class Wasm {
	public async parseInput(file: string): Promise<string | null> {
		let result: string | null;
		if (get(useWorker) && get(workerReady)) {
			const response = await fetch('/sw/parse', {
				method: 'POST',
				body: file
			});
			result = await response.text();
		} else {
			result = WASM_parse_input(file);
		}
		return !result ? null : result;
	}

	public async initialize(args: Arguments): Promise<RunningSolver | null> {
		let result: string;
		if (get(useWorker) && get(workerReady)) {
			const response = await fetch('/sw/initialize', {
				method: 'POST',
				body: JSON.stringify(args)
			});
			result = await response.text();
		} else {
			result = WASM_initialize(JSON.stringify(args));
		}
		try {
			return parse_as<RunningSolver>(result);
		} catch (error) {
			return null
		}
	}

	public async runGeneration(solver: RunningSolver): Promise<WASMGenerationReturn> {
		let result: string;
		if (get(useWorker) && get(workerReady)) {
			const response = await fetch('/sw/generate', {
				method: 'POST',
				body: JSON.stringify(solver)
			});
			result = await response.text();
		} else {
			result = WASM_run_generation(JSON.stringify(solver));
		}
		return parse_as<WASMGenerationReturn>(result);
	}

	public async generateOutput(simulation: Simulation): Promise<string> {
		let result: string;
		if (get(useWorker) && get(workerReady)) {
			const response = await fetch('/sw/output', {
				method: 'POST',
				body: JSON.stringify(simulation)
			});
			result = await response.text();
		} else {
			result = WASM_generate_output(JSON.stringify(simulation));
		}
		return result;
	}
}

export default new Wasm()
